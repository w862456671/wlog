package wlog

import (
	"bytes"
	"fmt"
	"os"
)

var defaultLog Ilog

type Ilog interface {
	GetLevel(string) Level
	SetLevel(Level, string)
	isEnabled(level Level, module string) bool
	log(level Level, calldepth int, record *Record)
}

func newIlog() {
	defaultLog = &wlog{newModuleLeveled()}
}

type wlog struct {
	Leveled
}

func (w *wlog) isEnabled(level Level, module string) bool {
	return w.isEnabledFor(level, module)
}

func (w *wlog) log(level Level, calldepth int, record *Record) {
	var buf bytes.Buffer

	col := colors[level]
	t := timeFormat()
	f := fileFormat(calldepth)
	m := fmt.Sprintf("[%s] ", record.Module)
	l := fmt.Sprintf("%s: ", level.String())

	buf.Write([]byte(col))
	buf.Write([]byte(f))
	buf.Write([]byte(t))
	buf.Write([]byte(m))
	buf.Write([]byte(l))
	buf.Write([]byte(record.Args))
	buf.Write([]byte("\033[0m\n"))

	out := os.Stdout
	out.Write(buf.Bytes())
}

func (w *wlog) SetLevel(level Level, module string) {
	w.sl(level, module)
}

func (w *wlog) GetLevel(module string) Level {
	return w.gl(module)
}
