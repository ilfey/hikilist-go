package logger

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"testing"
)

type Test struct {
	t *testing.T
}

func NewTest(t *testing.T) *Test {
	return &Test{
		t: t,
	}
}

func (l Test) Log(_ error) {}

func (l Test) LogPropagate(err error) error {
	return err
}

func (l Test) LogData(_ any) {}

func (l Test) Info(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Info() error: unknown type")
}

func (l Test) InfoPropagate(strOrErr any) error {
	if err, ok := strOrErr.(error); ok {
		return err
	}

	if str, ok := strOrErr.(string); ok {
		return errors.New(str)
	}

	l.t.Fatal("logger.InfoPropagate() error: unknown type")

	return nil
}

func (l Test) Debug(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Debug() error: unknown type")
}

func (l Test) DebugPropagate(strOrErr any) error {
	if err, ok := strOrErr.(error); ok {
		return err
	}

	if str, ok := strOrErr.(string); ok {
		return errors.New(str)
	}

	l.t.Fatal("logger.DebugPropagate() error: unknown type")

	return nil
}

func (l Test) Warn(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Warn() error: unknown type")
}

func (l Test) WarnPropagate(strOrErr any) error {
	if err, ok := strOrErr.(error); ok {
		return err
	}

	if str, ok := strOrErr.(string); ok {
		return errors.New(str)
	}

	l.t.Fatal("logger.WarnPropagate() error: unknown type")

	return nil
}

func (l Test) Error(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Error() error: unknown type")
}

func (l Test) ErrorPropagate(strOrErr any) error {
	if err, ok := strOrErr.(error); ok {
		return err
	}

	if str, ok := strOrErr.(string); ok {
		return errors.New(str)
	}

	l.t.Fatal("logger.ErrorPropagate() error: unknown type")

	return nil
}

func (l Test) Critical(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Critical() error: unknown type")
}

func (l Test) CriticalPropagate(strOrErr any) error {
	if err, ok := strOrErr.(error); ok {
		return err
	}

	if str, ok := strOrErr.(string); ok {
		return errors.New(str)
	}

	l.t.Fatal("logger.CriticalPropagate() error: unknown type")

	return nil
}

func (l Test) GetOutput() io.Writer {
	return io.Discard
}

func (l Test) SetOutput(_ io.Writer) {}

func (l Test) GetContext() context.Context { return context.Background() }

func (l Test) SetContext(_ context.Context) {}

func (l Test) Close() func() { return func() {} }
