package common

import (
	"database/sql"
	"os"
	"os/signal"
	"syscall"
)

// EnvironmentEnv is an environment type
type EnvironmentEnv string

// Indicates what environment the software is running in
const (
	EnvTesting EnvironmentEnv = "testing"
	EnvProd    EnvironmentEnv = "production"
	EnvDev     EnvironmentEnv = "development"
)

type subscriber struct {
	c chan bool
}

// Environment holds "global" application information such as a database connection,
// logging, the config, sessions, etc.
type Environment struct {
	DB           *DatabaseAccessor
	Config       *Config
	Env          EnvironmentEnv
	shutdownSubs []*subscriber
	shutdownChan chan os.Signal
}

// NewEnvironment creates an environment.
func NewEnvironment(t EnvironmentEnv) *Environment {
	return &Environment{Env: t}
}

// NewTestEnvironment creates and environment setup for testing.
func NewTestEnvironment() *Environment {
	return &Environment{
		Config: NewEmptyConfig(),
		Env:    EnvTesting,
	}
}

// IsTesting checks if the current environment is testing
func (e *Environment) IsTesting() bool {
	return (e.Env == EnvTesting)
}

// IsProd checks if the current environment is production
func (e *Environment) IsProd() bool {
	return (e.Env == EnvProd)
}

// IsDev checks if the current environment is development
func (e *Environment) IsDev() bool {
	return (e.Env == EnvDev)
}

// SubscribeShutdown returns a channel that the caller can block on. The channel
// will receive a value when the application receives a shutdown signal.
func (e *Environment) SubscribeShutdown() <-chan bool {
	e.shutdownWatcher() // Start the watcher

	sub := &subscriber{
		c: make(chan bool, 1),
	}

	e.shutdownSubs = append(e.shutdownSubs, sub)
	return sub.c
}

func (e *Environment) shutdownWatcher() {
	if e.shutdownChan != nil {
		return
	}

	e.shutdownChan = make(chan os.Signal, 1)
	signal.Notify(e.shutdownChan, os.Interrupt, syscall.SIGTERM)
	go func(env *Environment) {
		<-e.shutdownChan
		//e.Log.Notice("Calling shutdown subscribers")
		for _, sub := range e.shutdownSubs {
			//e.Log.Debugf("Calling shutdown subscriber %d", i)
			sub.c <- true
		}
	}(e)
}

// DatabaseAccessor wraps an sql.DB with a driver string.
type DatabaseAccessor struct {
	*sql.DB
	Driver string
}

// SchemaVersion queries the database and returns the current version.
func (d *DatabaseAccessor) SchemaVersion() int {
	var currDBVer int
	verRow := d.DB.QueryRow(`SELECT "value" FROM "settings" WHERE "id" = 'db_version'`)
	if verRow == nil {
		return 0
	}
	verRow.Scan(&currDBVer)
	return currDBVer
}

// SystemInitFunc is a function to be run at the start of the application.
type SystemInitFunc func(*Environment) error

var systemInitFuncs []SystemInitFunc

// RegisterSystemInitFunc registers a function to be run on system init.
func RegisterSystemInitFunc(f SystemInitFunc) {
	if systemInitFuncs == nil {
		systemInitFuncs = make([]SystemInitFunc, 0, 1)
	}
	systemInitFuncs = append(systemInitFuncs, f)
}

// RunSystemInits executes the registered init functions in the order they were
// registered.
func RunSystemInits(e *Environment) error {
	for _, f := range systemInitFuncs {
		if err := f(e); err != nil {
			return err
		}
	}
	return nil
}
