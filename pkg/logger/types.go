package logger

import (
	"time"
)

const (
	InfoLevel = iota
	DebugLevel
	WarnLevel
	ErrorLevel
	CriticalLevel

	InfoLevelReadable     = "INFO"
	DebugLevelReadable    = "DEBUG"
	WarnLevelReadable     = "WARN"
	ErrorLevelReadable    = "ERROR"
	CriticalLevelReadable = "CRITICAL"

	DebugLogType = "debug"
	InfoLogType  = "info"
	ErrorLogType = "error"
)

/* ===== LoggableError ===== */

type LoggableError interface {
	Error() string
	Level() int
}

/* ===== RequestIdAware ===== */

type RequestIdAware interface {
	RequestID() string
	SetRequestID(id string)
}

/* ===== introspectionError ===== */

type introspectedError interface {
	Date() time.Time
	Error() string
	File() string
	Func() string
	Line() int
	Level() int
	Type() string
	RequestId() string
	SetRequestId(id string)
}

/* ===== introspectionError ===== */

type introspectionError struct {
	Dt time.Time `json:"date"`
	Rq string    `json:"requestID,omitempty"`
	Tp string    `json:"type"`
	Mg string    `json:"message"`
	Fl string    `json:"file"`
	Fn string    `json:"function"`
	Ln int       `json:"line"`
}

/* ===== introspectionError.<Getter> ===== */

func (e *introspectionError) Date() time.Time {
	return e.Dt
}

func (e *introspectionError) Error() string {
	return e.Mg
}

func (e *introspectionError) File() string {
	return e.Fl
}

func (e *introspectionError) Func() string {
	return e.Fn
}

func (e *introspectionError) Line() int {
	return e.Ln
}

func (e *introspectionError) Level() int {
	return ErrorLevel
}

func (e *introspectionError) Type() string {
	return ErrorLogType
}

func (e *introspectionError) RequestId() string {
	return e.Rq
}

func (e *introspectionError) SetRequestId(id string) {
	e.Rq = id
}

/* ===== errorLevelError ===== */

type infoLevelError struct{ introspectionError }

/* ===== errorLevelError.<Getter> ===== */

func (e *infoLevelError) Level() int {
	return InfoLevel
}

func (e *infoLevelError) Type() string {
	return InfoLogType
}

/* ===== debugLevelError ===== */

type debugLevelError struct{ introspectionError }

/* ===== debugLevelError.<Getter> ===== */

func (e *debugLevelError) Level() int {
	return DebugLevel
}

func (e *debugLevelError) Type() string {
	return DebugLogType
}

/* ===== warnLevelError ===== */

type warnLevelError struct{ introspectionError }

/* ===== warnLevelError.<Getter> ===== */

func (e *warnLevelError) Level() int {
	return WarnLevel
}

func (e *warnLevelError) Type() string {
	return ErrorLogType
}

/* ===== errorLevelError ===== */

type errorLevelError struct{ introspectionError }

/* ===== errorLevelError.<Getter> ===== */

func (e *errorLevelError) Level() int {
	return ErrorLevel
}

func (e *errorLevelError) Type() string {
	return ErrorLogType
}

/* ===== criticalLevelError ===== */

type criticalLevelError struct{ introspectionError }

/* ===== criticalLevelError.<Getter> ===== */

func (e *criticalLevelError) Level() int {
	return CriticalLevel
}

func (e *criticalLevelError) Type() string {
	return ErrorLogType
}
