package config

import (
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type AppConf struct {
	AppName string `yaml:"app_name"`
	Server Server `yaml:"server"`
	Logger Logger `yaml:"logger"`
	DB     DB     `yaml:"db"`
}

type Server struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type DB struct {
	Net      string `yaml:"net"`
	Driver   string `yaml:"driver"`
	Name     string `yaml:"name"`
	User     string `json:"-" yaml:"user"`
	Password string `json:"-" yaml:"password"`
	Host     string `yaml:"host"`
	MaxConn  int    `yaml:"max_conn"`
	Port     string `yaml:"port"`
	Timeout  int    `yaml:"timeout"`
}

func NewAppConf() AppConf {
	port := os.Getenv("SERVER_PORT")

	return AppConf{
		AppName: os.Getenv("APP_NAME"),
		Server: Server{
			Port: port,
		},
		Logger: Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: DB{
			Net:      os.Getenv("DB_NET"),
			Driver:   os.Getenv("DB_DRIVER"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}

func (a *AppConf) Init(logger *zap.Logger) {
	shutDownTimeOut, err := strconv.Atoi(os.Getenv("SHUTDOWN_TIMEOUT"))
	if err != nil {
		logger.Fatal("config: parse server shutdown timeout error")
	}
	shutDownTimeout := time.Duration(shutDownTimeOut) * time.Second
	a.Server.ShutdownTimeout = shutDownTimeout

	dbTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		logger.Fatal("config: parse db timeout err", zap.Error(err))
	}
	a.DB.Timeout = dbTimeout
}
