package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golibrary/config"
	"golibrary/docs"
	"golibrary/internal/db"
	"golibrary/internal/logger"

	author "golibrary/internal/author/repository"
	book "golibrary/internal/book/repository"
	user "golibrary/internal/user/repository"

	aservice "golibrary/internal/author/service"
	bservice "golibrary/internal/book/service"
	uservice "golibrary/internal/user/service"	
	
	acontroller "golibrary/internal/author/controller"
	bcontroller "golibrary/internal/book/controller"
	ucontroller "golibrary/internal/user/controller"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repos struct {
	Author author.Authorer
	Book   book.Booker
	User   user.Userer
}

func NewRepositories(db *sqlx.DB) *Repos {
	return &Repos{
		Author: author.NewAuthorRepository(db),
		Book:   book.NewBookRepository(db),
		User:   user.NewUserRepository(db),
	}
}

type Services struct {
	Author aservice.Authorer
	Book   bservice.Booker
	User   uservice.Userer
}

func NewServices(repos *Repos, logger *zap.Logger) *Services {
	return &Services{
		Author: aservice.NewAuthorService(repos.Author, logger),
		Book:   bservice.NewBookService(repos.Book, logger),
		User:   uservice.NewUserService(repos.User, logger),
	}
}

type Controllers struct {
	Author acontroller.Authorer
	Book   bcontroller.Booker
	User   ucontroller.Userer
}

func NewControllers(services *Services) *Controllers{
	authorController := acontroller.NewAuthor(services.Author)
	bookController := bcontroller.NewBook(services.Book)
	userController := ucontroller.NewUser(services.User)

	return &Controllers{
		Author: authorController,
		Book:   bookController,
		User:   userController,
	}
}

func main() {
	// Создание конфигурации приложения
	conf := config.NewAppConf()
	// Создание логгера
	logger := logger.NewLogger(conf, os.Stdout)
	// Инициализация конфигурации приложения с логгером
	conf.Init(logger)

	// Инициализация БД
	dbx, err := db.Init(conf.DB, logger)
	if err != nil {
		logger.Fatal("error init db", zap.Error(err))
	}

	// Создание репозиториев
	repos := NewRepositories(dbx)

	// Создание сервисов
	services := NewServices(repos, logger)

	// Создание контроллеров
	controllers := NewControllers(services)

	r := chi.NewRouter()

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
	})

	// swagger документация
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		docs.SwaggerUI(w, r)
	})
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/swagger.json")
	})

	// маршруты library сервиса
	r.Group(func(r chi.Router) {
		bookController := controllers.Book
		r.Put("/library/take/{bookId}/{userId}", bookController.Take)
		r.Put("/library/return/{bookId}/{userId}", bookController.Return)
		r.Get("/library/books", bookController.List)
		r.Post("/library/book", bookController.Add)

		authorController := controllers.Author
		r.Get("/library/popular-authors", authorController.Top)
		r.Get("/library/authors", authorController.List)
		r.Post("/library/author", authorController.Add)

		userController := controllers.User
		r.Get("/library/users", userController.List)
		r.Post("/library/user", userController.Add)
	})

	server := &http.Server{
		Addr:         ":" + conf.Server.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Создание канала для получения сигналов остановки
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине
	go func() {
		logger.Info("server started", zap.String("port", conf.Server.Port))
		listener, err := net.Listen("tcp", server.Addr)
		if err != nil {
			fmt.Println("net.Listen error:", err.Error())
			return
		}
		server.Serve(listener)
	}()

	// Ожидание сигнала остановки - обработка сигналов SIGINT и SIGTERM
	// Отправить сигнал в контейнер -> docker stop --help
	sig := <-quit
	fmt.Printf("Received signal: %s\n", sig)

	// Создание контекста с таймаутом для graceful shutdown
	ctx, shutdown := context.WithTimeout(context.Background(), conf.Server.ShutdownTimeout)
	defer shutdown()

	// Остановка сервера с использованием graceful shutdown
	server.Shutdown(ctx)
	log.Println("Server stopped gracefully")
}
