package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"github.com/pkg/errors"
	"io"
	"log"
	"runtime"
	"sync"
	"time"
)

type abstract struct {
	mu     *sync.Mutex
	ctx    context.Context
	writer io.Writer
	errCh  chan introspectedError
	reqCh  chan any
}

func newAbstractLogger(ctx context.Context, w io.Writer, errBuff int, reqBuff int) (logger *abstract, closeFunc func()) {
	l := &abstract{
		mu:     new(sync.Mutex),
		ctx:    ctx,
		writer: w,
		errCh:  make(chan introspectedError, errBuff),
		reqCh:  make(chan any, reqBuff),
	}
	l.handle()
	return l, l.Close()
}

func (l *abstract) Close() (closeFunc func()) {
	return func() {
		close(l.errCh)
		close(l.reqCh)
	}
}

func (l *abstract) SetOutput(w io.Writer) {
	defer l.mu.Unlock()
	l.mu.Lock()
	l.writer = w
}

func (l *abstract) GetOutput() io.Writer {
	return l.writer
}

func (l *abstract) SetContext(ctx context.Context) {
	defer l.mu.Unlock()
	l.mu.Lock()
	l.ctx = ctx
}

func (l *abstract) GetContext() context.Context {
	return l.ctx
}

func (l *abstract) LogData(data any) {
	l.reqCh <- data
}

func (l *abstract) Log(err error) {
	file, function, line := l.trace()
	l.log(err, file, function, line)
}

func (l *abstract) LogPropagate(err error) error {
	file, function, line := l.trace()
	l.log(err, file, function, line)
	return err
}

func (l *abstract) Info(strOrErr any) {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &infoLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: InfoLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}
}

func (l *abstract) InfoPropagate(strOrErr any) error {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &infoLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: InfoLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}

	return err
}

func (l *abstract) Debug(strOrErr any) {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &debugLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: DebugLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}
}

func (l *abstract) DebugPropagate(strOrErr any) error {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &debugLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: DebugLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}

	return err
}

func (l *abstract) Warn(strOrErr any) {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &warnLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: ErrorLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}
}

func (l *abstract) WarnPropagate(strOrErr any) error {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &warnLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: ErrorLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}

	return err
}

func (l *abstract) Error(strOrErr any) {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &errorLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: ErrorLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}
}

func (l *abstract) ErrorPropagate(strOrErr any) error {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &errorLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: ErrorLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}

	return err
}

func (l *abstract) Critical(strOrErr any) {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &criticalLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: ErrorLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}
}

func (l *abstract) CriticalPropagate(strOrErr any) error {
	file, function, line := l.trace()

	err := l.error(strOrErr)

	l.errCh <- &criticalLevelError{
		introspectionError{
			Dt: time.Now(),
			Mg: err.Error(),
			Tp: ErrorLogType,
			Fl: file,
			Fn: function,
			Ln: line,
		},
	}

	return err
}

func (l *abstract) handle() {
	go func() {
		for err := range l.errCh {
			l.mu.Lock()
			if uniqReqID := l.ctx.Value(enum.RequestIDContextKey); uniqReqID != nil {
				if strUniqReqID, ok := uniqReqID.(string); ok {

					err.SetRequestId(strUniqReqID)
				}
			}
			l.mu.Unlock()

			j, e := json.MarshalIndent(err, "", "  ")
			if e != nil {
				_, fmtErr := fmt.Fprintln(l.writer, e)
				if fmtErr != nil {
					log.Println(err)
					panic(fmtErr)
				}
			} else {
				_, fmtErr := fmt.Fprintln(l.writer, string(j))
				if fmtErr != nil {
					log.Println(err)
					panic(fmtErr)
				}
			}
		}
	}()

	go func() {
		for info := range l.reqCh {
			l.mu.Lock()
			if uniqReqID := l.ctx.Value(enum.RequestIDContextKey); uniqReqID != nil {
				if strUniqReqID, ok := uniqReqID.(string); ok {
					if infoObj, iok := info.(RequestIdAware); iok {
						infoObj.SetRequestID(strUniqReqID)
					}
				}
			}
			l.mu.Unlock()

			j, e := json.MarshalIndent(info, "", "  ")
			if e != nil {
				_, fmtErr := fmt.Fprintln(l.writer, e)
				if fmtErr != nil {
					log.Println(info)
					panic(fmtErr)
				}
			} else {
				_, fmtErr := fmt.Fprintln(l.writer, string(j))
				if fmtErr != nil {
					log.Println(info)
					panic(fmtErr)
				}
			}
		}
	}()
}

func (l *abstract) log(e error, file string, function string, line int) {
	err, isLoggableErr := e.(LoggableError)
	if !isLoggableErr {
		l.errCh <- &errorLevelError{
			introspectionError{
				Dt: time.Now(),
				Mg: e.Error(),
				Tp: ErrorLogType,
				Fl: file,
				Fn: function,
				Ln: line,
			},
		}
		return
	}

	switch err.Level() {
	case InfoLevel:
		l.errCh <- &infoLevelError{
			introspectionError{
				Dt: time.Now(),
				Mg: err.Error(),
				Tp: InfoLogType,
				Fl: file,
				Fn: function,
				Ln: line,
			},
		}
		return
	case DebugLevel:
		l.errCh <- &debugLevelError{
			introspectionError{
				Dt: time.Now(),
				Mg: err.Error(),
				Tp: DebugLogType,
				Fl: file,
				Fn: function,
				Ln: line,
			},
		}
		return
	case WarnLevel:
		l.errCh <- &warnLevelError{
			introspectionError{
				Dt: time.Now(),
				Mg: err.Error(),
				Tp: ErrorLogType,
				Fl: file,
				Fn: function,
				Ln: line,
			},
		}
		return
	case ErrorLevel:
		l.errCh <- &errorLevelError{
			introspectionError{
				Dt: time.Now(),
				Mg: err.Error(),
				Tp: ErrorLogType,
				Fl: file,
				Fn: function,
				Ln: line,
			},
		}
		return
	case CriticalLevel:
		l.errCh <- &criticalLevelError{
			introspectionError{
				Dt: time.Now(),
				Mg: err.Error(),
				Tp: ErrorLogType,
				Fl: file,
				Fn: function,
				Ln: line,
			},
		}
		return
	}

	panic("logger.log(): undefined error level received")
}

func (l *abstract) trace() (file string, function string, line int) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return frame.File, frame.Func.Name(), frame.Line
}

func (l *abstract) error(strOrErr any) error {
	err, isErr := strOrErr.(error)
	if isErr {
		return err
	} else {
		str, isStr := strOrErr.(string)
		if isStr {
			return errors.New(str)
		}
	}
	panic("logger.error(): logging data is not a string or error type")
}

func ToReadableLevel(err introspectedError) string {
	switch err.Level() {
	case InfoLevel:
		return InfoLevelReadable
	case DebugLevel:
		return DebugLevelReadable
	case WarnLevel:
		return WarnLevelReadable
	case ErrorLevel:
		return ErrorLevelReadable
	case CriticalLevel:
		return CriticalLevelReadable
	}
	panic("logger.ToReadableLevel(): received undefined error level")
}

func ToLevel(readableLevel string) int {
	switch readableLevel {
	case InfoLevelReadable:
		return InfoLevel
	case DebugLevelReadable:
		return DebugLevel
	case WarnLevelReadable:
		return WarnLevel
	case ErrorLevelReadable:
		return ErrorLevel
	case CriticalLevelReadable:
		return CriticalLevel
	}
	panic("logger.ToLevel(): received undefined readable level")
}
