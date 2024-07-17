package logger

import (
	"fmt"
	"strings"
)

func LevelFormatter(level Level, format string) string {
	var color string
	lvl := level

	switch level {
	case LevelTrace:
		color = "\x1b[37m"
	case LevelDebug:
		color = "\x1b[36m"
	case LevelInfo:
		color = "\x1b[34m"
		lvl = "info "
	case LevelWarn:
		color = "\x1b[33m"
		lvl = "warn "
	case LevelError:
		color = "\x1b[31m"
	case LevelFatal:
		color = "\x1b[35m"
	case LevelPanic:
		color = "\x1b[35m"
	}

	return fmt.Sprintf("%s%s ▶ %s%s", color, strings.ToUpper(lvl), "\x1b[0m", format)
}

func TextFormatter(level Level, format string, args ...interface{}) string {
	return fmt.Sprintf(
		LevelFormatter(level, format)+"\n",
		args...,
	)
}