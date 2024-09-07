package loggerInterface

import (
	"context"
)

type Logger interface {
	Object(obj any)

	Propagate(err error) error

	Debug(strOrErr any)
	Info(strOrErr any)
	Warn(strOrErr any)
	Error(strOrErr any)
	Critical(strOrErr any)

	GetContext() context.Context
	SetContext(context context.Context)

	Close() func()
}
