package lib

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Formatter struct {
	LogFormatter log.TextFormatter
}

// A custom formatter for the logger.
func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	// example [2022-11-25 13:25:15] - [INFO] log-message...
	return []byte(fmt.Sprintf("[%s] - [%s] %s\n", entry.Time.Format(f.LogFormatter.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}
