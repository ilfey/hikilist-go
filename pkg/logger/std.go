package logger

import (
	"context"
	"os"
)

type Std struct {
	*abstract
}

func NewStdErr(ctx context.Context, errBuff int, reqBuff int) (logger *Std, closeFunc func()) {
	abstractLogger, closeFunc := newAbstractLogger(ctx, os.Stderr, errBuff, reqBuff)

	return &Std{
		abstract: abstractLogger,
	}, closeFunc
}
