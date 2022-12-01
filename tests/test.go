package main

import (
	log "github.com/wayne9499/glogger"
)

var (
	logging = log.Logger{}
	logger  = logging.GetLogger()
)

func main() {
	logger.Info("This is info logger")
	logger.Debug("This is debug logger")
	logger.Fatal("This is fatal logger")
}
