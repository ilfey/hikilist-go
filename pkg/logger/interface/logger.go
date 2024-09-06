package loggerInterface

import (
	"context"
	"io"
)

type Logger interface {
	Log(err error)
	LogPropagate(err error) error
	LogData(data any)

	Info(strOrErr any)
	InfoPropagate(strOrErr any) error

	Debug(strOrErr any)
	DebugPropagate(strOrErr any) error

	Warn(strOrErr any)
	WarnPropagate(strOrErr any) error

	Error(strOrErr any)
	ErrorPropagate(strOrErr any) error

	Critical(strOrErr any)
	CriticalPropagate(strOrErr any) error

	GetOutput() io.Writer
	SetOutput(writer io.Writer)

	GetContext() context.Context
	SetContext(context context.Context)

	Close() func()
}
