package server

import (
	"net"
	"time"
)

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
	Port string
}

func (config *Config) Address() string {
	return net.JoinHostPort(config.Host, config.Port)
}
