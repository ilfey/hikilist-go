package logger

import (
	"context"
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

func (l Test) Propagate(err error) error {
	return err
}

func (l Test) Object(_ any) {}

func (l Test) Info(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Info() error: unknown type")
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

func (l Test) Warn(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Warn() error: unknown type")
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

func (l Test) Critical(strOrErr any) {
	if _, ok := strOrErr.(error); ok {
		return
	}

	if _, ok := strOrErr.(string); ok {
		return
	}

	l.t.Fatal("logger.Critical() error: unknown type")
}

func (l Test) GetContext() context.Context { return context.Background() }

func (l Test) SetContext(_ context.Context) {}

func (l Test) Close() func() { return func() {} }
