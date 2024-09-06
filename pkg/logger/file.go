package logger

import (
	"context"
	"os"
	"time"
)

type File struct {
	*abstract
}

func NewFile(ctx context.Context, errBuff int, reqBuff int) (*File, func(), error) {
	dir := "logs/"

	// Check if folder exists.
	_, err := os.Open(dir)
	if err != nil {
		// Handle folder not exist.
		if os.IsNotExist(err) {
			// Create folder.
			err := os.MkdirAll(dir, os.ModeDir)
			if err != nil {
				return nil, nil, err
			}
		} else {
			// Handle other errors...
			return nil, nil, err
		}
	}

	logfile, err := os.Create(dir + time.Now().Format("2006_01_02") + ".log")
	if err != nil {
		return nil, nil, err
	}

	abstractLogger, closeFunc := newAbstractLogger(ctx, logfile, errBuff, reqBuff)

	logger := &File{
		abstract: abstractLogger,
	}

	return logger, closeFunc, nil
}
