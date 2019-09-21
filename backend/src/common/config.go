package common

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

// PageSize is the number of items per page
var PageSize = 30

// Config defines the configuration struct for the application
type Config struct {
	sourceFile string
	Logging    struct {
		Enabled    bool
		EnableHTTP bool
		Level      string
		Path       string
	}
	Database struct {
		Address      string
		Port         int
		Username     string
		Password     string
		Name         string
		Retry        int
		RetryTimeout string
	}
	Webserver struct {
		Address string
		Port    int
	}
	Auth struct {
		SigningKey string
	}
}

// FindConfigFile searches for a configuration file. The order of search is
// environment, current dir, home dir, and /etc.
func FindConfigFile() string {
	filename := ""

	if os.Getenv("PG_CONFIG") != "" && FileExists(os.Getenv("PG_CONFIG")) {
		filename = os.Getenv("PG_CONFIG")
	} else if FileExists("./config.toml") {
		filename = "./config.toml"
	} else if FileExists("./config/config.toml") {
		filename = "./config/config.toml"
	} else if FileExists(os.ExpandEnv("$HOME/.pg/config.toml")) {
		filename = os.ExpandEnv("$HOME/.pg/config.toml")
	} else if FileExists("/etc/packet-guardian/config.toml") {
		filename = "/etc/packet-guardian/config.toml"
	}

	return filename
}

// NewEmptyConfig returns an empty config with type defaults only.
func NewEmptyConfig() *Config {
	return &Config{}
}

// NewConfig reads the given filename into a Config. If filename is empty,
// the config is looked for in the documented order.
func NewConfig(configFile string) (conf *Config, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()

	if configFile == "" {
		configFile = "config.toml"
	}

	var con Config
	if _, err := toml.DecodeFile(configFile, &con); err != nil {
		return nil, err
	}
	con.sourceFile = configFile

	return setSensibleDefaults(&con)
}

func setSensibleDefaults(c *Config) (*Config, error) {
	// Anything not set here implies its zero value is the default

	// Logging
	c.Logging.Level = setStringOrDefault(c.Logging.Level, "notice")

	// Database
	c.Database.Address = setStringOrDefault(c.Database.Address, "localhost")
	c.Database.RetryTimeout = setStringOrDefault(c.Database.RetryTimeout, "10s")

	// Webserver
	c.Webserver.Port = setIntOrDefault(c.Webserver.Port, 80)

	return c, nil
}

// Given string s, if it is empty, return v else return s.
func setStringOrDefault(s, v string) string {
	if s == "" {
		return v
	}
	return s
}

// Given integer s, if it is 0, return v else return s.
func setIntOrDefault(s, v int) int {
	if s == 0 {
		return v
	}
	return s
}

func validateURL(path, description string) (string, error) {
	path = strings.TrimRight(path, "/")
	if _, err := url.Parse(path); err != nil {
		return "", fmt.Errorf("Invalid %s URL", description)
	}
	return path, nil
}
