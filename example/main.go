package main

import (
	"fmt"
	"github.com/w862456671/wlog"
)

func main() {
	logger := wlog.NewModuleLogger("example")
	err := logger.SetLevel("debug")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wlog.GetLevel("example"))

	logger.Debug("debug")
	logger.Debugf("debugf: %s", "test")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("error")
	logger.Errorf("errorf: %d", 123)
	logger.Critical("critical")
}
