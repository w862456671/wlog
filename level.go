package wlog

import (
	"errors"
	"strings"
)

var ErrInvalidLogLevel = errors.New("logger: invalid log level")

type Level int

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

func (p Level) String() string {
	return levelNames[p]
}

func logLevel(level string) (Level, error) {
	for i, name := range levelNames {
		if strings.EqualFold(name, level) {
			return Level(i), nil
		}
	}
	return ERROR, ErrInvalidLogLevel
}

type Leveled interface {
	gl(string) Level
	sl(Level, string)
	isEnabledFor(Level, string) bool
}

type moduleLeveled struct {
	levels 		map[string]Level
}

func newModuleLeveled() Leveled {
	return &moduleLeveled{
		levels: make(map[string]Level),
	}
}

func (l *moduleLeveled) gl(module string) Level {
	level, exists := l.levels[module]
	if exists == false {
		level, exists = l.levels[""]
		if exists == false {
			level = DEBUG
		}
	}
	return level
}

func (l *moduleLeveled) sl(level Level, module string) {
	l.levels[module] = level
}

func (l *moduleLeveled) isEnabledFor(level Level, module string) bool {
	return level <= l.levels[module]
}
