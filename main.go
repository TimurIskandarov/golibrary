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
	"golibrary/internal/db"
	"golibrary/internal/logger"
	"golibrary/internal/responder"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func main() {
	// Создание конфигурации приложения
	conf := config.NewAppConf()
	// Создание логгера
	logger := logger.NewLogger(conf, os.Stdout)
	// Инициализация конфигурации приложения с логгером
	conf.Init(logger)


	// Создание респондера
	_ := responder.NewResponder(logger)

	// Инициализация БД
	_, err := db.Init(conf.DB, logger)
	if err != nil {
		logger.Fatal("error init db", zap.Error(err))
	}

	// Создание сервисов
	// ...

	// Создание контроллеров
	// ...

	r := chi.NewRouter()

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
	})


	// маршруты library сервиса
	r.Group(func(r chi.Router) {
		libController := controllers.Library
		r.Get("/library/{bookId}/{userId}", libController.BookTake)
		r.Post("/library/{bookId}/{userId}", libController.BookReturn)
		r.Get("/library/popular-authors", libController.AuthorsTop)
		r.Get("/library/authors", libController.AuthorsList)
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
