package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ilfey/hikilist-go/internal/domain/enum"
	"io"
	"log"
	"runtime"
	"sync"
)

type abstract struct {
	mu     *sync.Mutex
	ctx    context.Context
	writer io.Writer
	errCh  chan LoggableError
	reqCh  chan any
}

func newAbstractLogger(ctx context.Context, w io.Writer, errBuff int, reqBuff int) (logger *abstract, closeFunc func()) {
	l := &abstract{
		mu:     new(sync.Mutex),
		ctx:    ctx,
		writer: w,
		errCh:  make(chan LoggableError, errBuff),
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

func (l *abstract) SetContext(ctx context.Context) {
	defer l.mu.Unlock()
	l.mu.Lock()
	l.ctx = ctx
}

func (l *abstract) GetContext() context.Context {
	return l.ctx
}

func (l *abstract) Object(data any) {
	l.reqCh <- data
}

func (l *abstract) Propagate(err error) error {
	lErr := l.toLoggableErrorOrAddTrace(err, l.trace(), DebugLevel)

	return lErr
}

func (l *abstract) Debug(srtOrErr any) {
	lErr := l.toLoggableErrorOrAddTrace(srtOrErr, l.trace(), DebugLevel)

	l.log(DebugLevel, lErr)
}

func (l *abstract) Info(strOrErr any) {
	lErr := l.toLoggableErrorOrAddTrace(strOrErr, l.trace(), InfoLevel)

	l.log(WarnLevel, lErr)
}

func (l *abstract) Warn(strOrErr any) {
	lErr := l.toLoggableErrorOrAddTrace(strOrErr, l.trace(), WarnLevel)

	l.log(WarnLevel, lErr)
}

func (l *abstract) Error(strOrErr any) {
	lErr := l.toLoggableErrorOrAddTrace(strOrErr, l.trace(), ErrorLevel)

	l.log(ErrorLevel, lErr)
}

func (l *abstract) Critical(strOrErr any) {
	lErr := l.toLoggableErrorOrAddTrace(strOrErr, l.trace(), CriticalLevel)

	l.log(CriticalLevel, lErr)
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

func (l *abstract) log(_ Level, lErr LoggableError) {
	l.errCh <- lErr
}

func (l *abstract) trace() *Trace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return &Trace{
		File: frame.File,
		Func: frame.Func.Name(),
		Line: frame.Line,
	}
}

func (l *abstract) toLoggableErrorOrAddTrace(a any, trace *Trace, lvl Level) LoggableError {
	err, isLoggableErr := a.(LoggableError)
	if !isLoggableErr {
		return newLoggableError(a, trace, lvl)
	}

	err.AddTrace(trace)

	return err
}
