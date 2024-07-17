package logger

import (
	"io"
)

var (
	std = New()
)

func Log(level Level, args ...interface{}) {
	std.Log(level, args...)
}

func Logf(level Level, format string, args ...interface{}) {
	std.Logf(level, format, args...)
}

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func GetOut() io.Writer {
	return std.GetOut()
}

func SetOut(out io.Writer) {
	std.SetOut(out)
}

func SetTextFormatter(f func(level Level, format string, args ...interface{}) string) {
	std.SetTextFormatter(f)
}

func GetLevel() Level {
	return std.GetLevel()
}

func SetLevel(level Level) {
	std.SetLevel(level)
}