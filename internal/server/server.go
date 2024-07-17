package server

import (
	"fmt"
	"time"

	"github.com/codegangsta/negroni"

	"net/http"
)

// Сервер
type Server struct {
	config *Config
	router Router
}

// Конфиг сервера
type Config struct {
	// Максимальная продолжительность чтения всего запроса в миллисекундах
	ReadTimeout time.Duration
	// Максимальная продолжительность до истечения времени ожидания записи ответа в миллисекундах
	WriteTimeout time.Duration
	// Максимальное время ожидания следующего запроса при включенном режиме keep-alive в миллисекундах
	IdleTimeout time.Duration
	// Количество времени, отведенное на чтение заголовков запроса в миллисекундах
	ReadHeaderTimeout time.Duration

	// IP-адрес
	Host string
	// Порт
	Port int
}

// Адрес сервера
func (config *Config) Address() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}

// Роутер
type Router interface {
	Bind() http.Handler
}

// Конструктор сервера
func NewServer(config *Config, router Router) *Server {
	return &Server{
		config,
		router,
	}
}

// Запуск сервера
func (server *Server) Run() error {
	ngRouter := server.router.Bind()

	ngClassic := negroni.Classic()

	ngClassic.UseHandler(ngRouter)

	srv := &http.Server{
		ReadTimeout:       server.config.ReadTimeout * time.Millisecond,
		WriteTimeout:      server.config.WriteTimeout * time.Millisecond,
		IdleTimeout:       server.config.IdleTimeout * time.Millisecond,
		ReadHeaderTimeout: server.config.ReadHeaderTimeout * time.Millisecond,

		Addr:    server.config.Address(),
		Handler: ngClassic,
	}

	return srv.ListenAndServe()
}
