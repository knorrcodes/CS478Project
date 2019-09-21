package common

import (
	"fmt"

	log "github.com/lfkeitel/verbose/v5"
)

var logLevels = map[string]log.LogLevel{
	"debug":     log.LogLevelDebug,
	"info":      log.LogLevelInfo,
	"notice":    log.LogLevelNotice,
	"warning":   log.LogLevelWarning,
	"error":     log.LogLevelError,
	"critical":  log.LogLevelCritical,
	"alert":     log.LogLevelAlert,
	"emergency": log.LogLevelEmergency,
	"fatal":     log.LogLevelFatal,
}

// SetupLogging creates a standard logger with console and file logging.
func SetupLogging(c *Config) {
	setupFileLogging(c.Logging.Level, c.Logging.Path)
}

func setupFileLogging(level, path string) {
	if path == "" {
		return
	}

	fh, err := log.NewFileTransport(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fh.Formatter = log.NewJSONFormatter()
	log.AddTransport(fh)
}
