package log

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger 
var DebugLogger *log.Logger

func SetLogger(system string, debug bool) () {
	system = "<" + system + "> "
	logger := log.New(os.Stdout, system, log.Ltime)
	debugLogger := log.New(os.Stdout, system, log.Ldate | log.Ltime | log.Lshortfile)
	if !debug {
		debugLogger.SetOutput(io.Discard)
	}
	Logger = logger
	DebugLogger = debugLogger
}