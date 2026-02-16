package main

import (
	"flag"

	"github.com/rcastellotti/loggo/logger"
)

func init() {
	logger.Log = logger.NewLogger(logger.ERROR)
}

var verbosityLevelHelpMsg = `increase verbosity (can be repeated)
-v > warning
-v -v > info
-v -v -v > debug
-v -v -v -v > trace
`

func main() {
	flag.Var(&logger.VerbosityLevel, "v", verbosityLevelHelpMsg)
	flag.Parse()
	logger.Init()

	logger.Log.Error("This is an error message (always shown)")
	logger.Log.Warn("This is a warning message (shown with -v)")
	logger.Log.Info("This is an info message (shown with -v -v)")
	logger.Log.Debug("This is a debug message (shown with -v -v -v)")
	logger.Log.Trace("This is a trace message (shown with -v -v -v -v)")
}
