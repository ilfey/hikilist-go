package logger

import (
	"github.com/pkg/errors"
	"time"
)

type Level uint

const (
	InfoLevel Level = iota + 1
	DebugLevel
	WarnLevel
	ErrorLevel
	CriticalLevel

	DebugLogType = "debug"
	InfoLogType  = "info"
	ErrorLogType = "error"
)

/* ===== Trace ===== */

type Trace struct {
	File string `json:"file"`
	Func string `json:"function"`
	Line int    `json:"line"`
}

/* ===== RequestIdAware ===== */

type RequestIdAware interface {
	RequestID() string
	SetRequestID(id string)
}

/* ===== LoggableError ===== */

type LoggableError interface {
	Date() time.Time

	Error() string
	Unwrap() error

	Level() Level
	Type() string

	AddTrace(trace *Trace)

	RequestId() string
	SetRequestId(id string)
}

/* ===== loggableErrorImpl ===== */

type loggableErrorImpl struct {
	Er error     `json:"-"`
	Dt time.Time `json:"date"`
	Rq string    `json:"requestID,omitempty"`
	Tp string    `json:"type"`
	Mg string    `json:"message"`
	Sk []*Trace  `json:"stack,omitempty"`
}

func newLoggableError(errOrString any, trace *Trace, lvl Level) LoggableError {
	var (
		err error
		msg string
	)

	if e, ok := errOrString.(error); ok {
		err = e
		msg = err.Error()
	} else {
		if errString, ok := errOrString.(string); ok {
			msg = errString
			err = errors.New(errString)
		} else {
			panic("newError(): unknown type")
		}
	}

	lErr := loggableErrorImpl{
		Er: err,
		Dt: time.Now(),
		Mg: msg,
		Sk: []*Trace{trace},
	}

	switch lvl {
	case DebugLevel:
		return &debugLevelError{loggableErrorImpl: lErr}
	case InfoLevel:
		return &infoLevelError{loggableErrorImpl: lErr}
	case WarnLevel:
		return &warnLevelError{loggableErrorImpl: lErr}
	case ErrorLevel:
		return &errorLevelError{loggableErrorImpl: lErr}
	case CriticalLevel:
		return &criticalLevelError{loggableErrorImpl: lErr}
	}

	panic("newError(): unknown level")
}

/* ===== loggableErrorImpl.<Getter> ===== */

func (e *loggableErrorImpl) Date() time.Time {
	return e.Dt
}

func (e *loggableErrorImpl) Error() string {
	return e.Mg
}

func (e *loggableErrorImpl) Level() Level {
	return ErrorLevel
}

func (e *loggableErrorImpl) Type() string {
	return ErrorLogType
}

func (e *loggableErrorImpl) RequestId() string {
	return e.Rq
}

func (e *loggableErrorImpl) SetRequestId(id string) {
	e.Rq = id
}

func (e *loggableErrorImpl) Unwrap() error {
	return e.Er
}

func (e *loggableErrorImpl) AddTrace(trace *Trace) {
	e.Sk = append(e.Sk, trace)
}

/* ===== infoLevelError ===== */

type infoLevelError struct{ loggableErrorImpl }

/* ===== errorLevelError.<Getter> ===== */

func (e *infoLevelError) Level() Level {
	return InfoLevel
}

func (e *infoLevelError) Type() string {
	return InfoLogType
}

/* ===== debugLevelError ===== */

type debugLevelError struct{ loggableErrorImpl }

/* ===== debugLevelError.<Getter> ===== */

func (e *debugLevelError) Level() Level {
	return DebugLevel
}

func (e *debugLevelError) Type() string {
	return DebugLogType
}

/* ===== warnLevelError ===== */

type warnLevelError struct{ loggableErrorImpl }

/* ===== warnLevelError.<Getter> ===== */

func (e *warnLevelError) Level() Level {
	return WarnLevel
}

func (e *warnLevelError) Type() string {
	return ErrorLogType
}

/* ===== errorLevelError ===== */

type errorLevelError struct{ loggableErrorImpl }

/* ===== errorLevelError.<Getter> ===== */

func (e *errorLevelError) Level() Level {
	return ErrorLevel
}

func (e *errorLevelError) Type() string {
	return ErrorLogType
}

/* ===== criticalLevelError ===== */

type criticalLevelError struct{ loggableErrorImpl }

/* ===== criticalLevelError.<Getter> ===== */

func (e *criticalLevelError) Level() Level {
	return CriticalLevel
}

func (e *criticalLevelError) Type() string {
	return ErrorLogType
}
