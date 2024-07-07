package server

import (
	"fmt"

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

	return http.ListenAndServe(server.config.Address(), ngClassic)
}
