package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel int

const (
	ERROR LogLevel = iota
	WARN
	INFO
	DEBUG
	TRACE
)

var levelNames = map[LogLevel]string{
	ERROR: "ERROR",
	WARN:  "WARN ",
	INFO:  "INFO ",
	DEBUG: "DEBUG",
	TRACE: "TRACE",
}

type Logger struct {
	level  LogLevel
	logger *log.Logger
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(os.Stderr, "", 0),
	}
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) log(level LogLevel, format string, args ...any) {
	if level <= l.level {
		timestamp := time.Now().Format("15:04:05")
		prefix := fmt.Sprintf("[%s] %s ", timestamp, levelNames[level])
		message := fmt.Sprintf(format, args...)
		l.logger.Printf("%s%s", prefix, message)
	}
}

func (l *Logger) Error(format string, args ...any) {
	l.log(ERROR, format, args...)
}

func (l *Logger) Warn(format string, args ...any) {
	l.log(WARN, format, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.log(INFO, format, args...)
}

func (l *Logger) Debug(format string, args ...any) {
	l.log(DEBUG, format, args...)
}

func (l *Logger) Trace(format string, args ...any) {
	l.log(TRACE, format, args...)
}

var Log *Logger

type Verbosity int

func (v *Verbosity) String() string {
	return fmt.Sprintf("%d", *v)
}

func (v *Verbosity) Set(s string) error {
	*v++
	return nil
}

func (v *Verbosity) IsBoolFlag() bool {
	return true
}

var VerbosityLevel Verbosity

func Init() {
	var level LogLevel
	switch VerbosityLevel {
	case 0:
		level = ERROR
	case 1:
		level = WARN
	case 2:
		level = INFO
	case 3:
		level = DEBUG
	default:
		level = TRACE
	}

	Log = NewLogger(level)
}
