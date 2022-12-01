package main

import (
	log "github.com/wayne9499/glogger"
)

var (
	logging = log.Logger{}
	logger  = logging.GetLogger()
)

func main() {
	logger.Info("This is the logger to test")
}
