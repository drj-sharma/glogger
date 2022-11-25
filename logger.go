package logger

import (
	log "github.com/sirupsen/logrus"
	"github.com/wayne9499/glogger/lib"
)

type Logger struct {
	Level log.Level
}

func (l *Logger) GetLogger() *log.Logger {
	logger := log.New()
	logger.Formatter = &lib.Formatter{
		LogFormatter: log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        lib.REFERENCE_TIME,
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
	}
	return logger
}
