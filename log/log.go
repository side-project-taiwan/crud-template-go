package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func SetupLog() {
	log.Infoln("SetupLog()...")
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true,
	})
	// Set output to standard output (default is stderr)
	log.SetOutput(os.Stdout)
	// Set the log level to debug
	log.SetLevel(log.DebugLevel)
	log.Infoln("This is a test log entry")
}
