package main

import (
	"fmt"
	"github.com/w862456671/wlog"
)

func main() {
	logger := wlog.NewModuleLogger("example")
	err := logger.SetLevel("notice")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wlog.GetLevel("example"))

	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("error")
	logger.Critical("critical")
}
