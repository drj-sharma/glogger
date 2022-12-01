package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/wayne9499/glogger/lib"
)

type Logger struct {
	Level log.Level
}

// Parsing the log level from the passing argument.
func (l *Logger) parseLoggerLevel(lvl string) (log.Level, error) {
	level, err := log.ParseLevel(lvl)
	if err != nil {
		return log.DebugLevel, err
	}
	return level, nil
}

// Creating a new logger and setting the log level and formatter.
func (l *Logger) GetLogger() *log.Logger {
	logger := log.New()
	level := os.Getenv("LOG_LEVEL")
	logLevel, _ := l.parseLoggerLevel(level)
	logger.Level = logLevel
	logger.SetOutput(&lib.LogWriter{})
	logger.Formatter = &lib.Formatter{
		LogFormatter: log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        lib.REFERENCE_TIME,
			ForceColors:            false,
			DisableLevelTruncation: false,
		},
	}
	return logger
}
