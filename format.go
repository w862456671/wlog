package wlog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type color int

const (
	ColorBlack = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

var (
	colors = []string{
		CRITICAL: colorSeq(ColorMagenta),
		ERROR:    colorSeq(ColorRed),
		WARNING:  colorSeq(ColorYellow),
		NOTICE:   colorSeq(ColorGreen),
		DEBUG:    colorSeq(ColorCyan),
	}
)

func colorSeq(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

func timeFormat() string {
	return time.Now().Format(" 2006-01-02 15:04:05 ")
}

func fileFormat(calldepth int) string {
	_, file, line, ok := runtime.Caller(calldepth + 1)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}
