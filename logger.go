package wlog

import (
	"time"
)

func init() {
	newIlog()
}

var timeNow = time.Now

type Record struct {
	Time   time.Time
	Module string
	Level  Level
	Args   string
}

type Logger struct {
	Module string
	ExtraCalldepth int
}

func (l *Logger) IsEnabledFor(level Level) bool {
	return defaultLog.isEnabled(level, l.Module)
}

func (l *Logger) log(lvl Level, args string) {
	if !l.IsEnabledFor(lvl) {
		return
	}

	record := &Record{
		Time:      timeNow(),
		Module:    l.Module,
		Level:     lvl,
		Args:      args,
	}

	defaultLog.log(lvl, 3+l.ExtraCalldepth, record)
}

func (l *Logger) SetLevel(level string) (err error) {
	var lvl Level

	lvl, err = logLevel(level)
	if err != nil {
		return err
	}

	defaultLog.SetLevel(lvl, l.Module)
	return
}

func GetLevel(module string) Level {
	return defaultLog.GetLevel(module)
}

func (l *Logger) Debug(args string) {
	l.log(DEBUG, args)
}

func (l *Logger) Info(args string) {
	l.log(INFO, args)
}

func (l *Logger) Error(args string) {
	l.log(ERROR, args)
}

func (l *Logger) Notice(args string) {
	l.log(NOTICE, args)
}

func (l *Logger) Warning(args string) {
	l.log(WARNING, args)
}

func (l *Logger) Critical(args string) {
	l.log(CRITICAL, args)
}

func NewModuleLogger(module string) *Logger {
	return &Logger{Module: module}
}
