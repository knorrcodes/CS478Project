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
	minLevel := log.LogLevelInfo
	if logLevel, ok := logLevels[c.Logging.Level]; ok {
		minLevel = logLevel
	}

	log.ClearTransports()
	setupConsoleLogging(minLevel)
	setupFileLogging(minLevel, c.Logging.Path)
}

func setupConsoleLogging(level log.LogLevel) {
	tt := log.NewTextTransport()
	tt.SetMinLevel(level)
	log.AddTransport(tt)
}

func setupFileLogging(level log.LogLevel, path string) {
	if path == "" {
		return
	}

	ft, err := log.NewFileTransport(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ft.SetMinLevel(level)
	ft.Formatter = log.NewJSONFormatter()
	log.AddTransport(ft)
}
