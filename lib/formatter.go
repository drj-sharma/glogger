package lib

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Formatter struct {
	LogFormatter log.TextFormatter
}

func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] - [%s] %s\n", entry.Time.Format(f.LogFormatter.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}
