package wlog

import (
	"fmt"
	"strings"
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
	Module         string
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
		Time:   timeNow(),
		Module: l.Module,
		Level:  lvl,
		Args:   args,
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

func (l *Logger) Debug(args ...interface{}) {
	l.log(DEBUG, formatArgs(args))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(DEBUG, templateArgs(format, args))
}

func (l *Logger) Info(args ...interface{}) {
	l.log(INFO, formatArgs(args))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(INFO, templateArgs(format, args))
}

func (l *Logger) Error(args ...interface{}) {
	l.log(ERROR, formatArgs(args))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(ERROR, templateArgs(format, args))
}

func (l *Logger) Notice(args ...interface{}) {
	l.log(NOTICE, formatArgs(args))
}

func (l *Logger) Noticef(format string, args ...interface{}) {
	l.log(NOTICE, templateArgs(format, args))
}

func (l *Logger) Warning(args ...interface{}) {
	l.log(WARNING, formatArgs(args))
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.log(WARNING, templateArgs(format, args))
}

func (l *Logger) Critical(args ...interface{}) {
	l.log(CRITICAL, formatArgs(args))
}

func (l *Logger) Criticalf(format string, args ...interface{}) {
	l.log(CRITICAL, templateArgs(format, args))
}

func formatArgs(args []interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(args...), "\n")
}

func templateArgs(format string, args []interface{}) string {
	msg := format

	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	return msg
}

func NewModuleLogger(module string) *Logger {
	return &Logger{Module: module}
}
