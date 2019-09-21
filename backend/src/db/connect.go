package db

import (
	"errors"
	"time"

	"koala.pos/src/common"

	log "github.com/lfkeitel/verbose/v5"
)

const DBVersion = 1

type dbInit interface {
	init(*common.DatabaseAccessor, *common.Config) error
}

var dbInits = make(map[string]dbInit)

type migrateFunc func(*common.DatabaseAccessor, *common.Config) error

func RegisterDatabaseAccessor(name string, db dbInit) {
	dbInits[name] = db
}

func NewDatabaseAccessor(e *common.Environment) (*common.DatabaseAccessor, error) {
	da := &common.DatabaseAccessor{}

	var err error
	retries := 0
	dur, err := time.ParseDuration(e.Config.Database.RetryTimeout)
	if err != nil {
		return nil, errors.New("Invalid RetryTimeout")
	}

	// This loop will break when no error occurs when connecting to a database
	// Or when the number of attempted retries is greater than configured
	shutdownChan := e.SubscribeShutdown()

	connector := newDBConnector()

	for {
		err = connector.init(da, e.Config)

		// If no error occurred, break
		// If an error occurred but retries is not set to inifinite and we've tried
		// too many times already, break
		if err == nil || (e.Config.Database.Retry != 0 && retries >= e.Config.Database.Retry) {
			break
		}

		retries++
		log.WithFields(log.Fields{
			"Attempts":    retries,
			"MaxAttempts": e.Config.Database.Retry,
			"Timeout":     e.Config.Database.RetryTimeout,
			"Error":       err,
		}).Error("Failed to connect to database. Retrying after timeout.")

		select {
		case <-shutdownChan:
			return nil, err
		case <-time.After(dur):
		}
	}

	da.SetConnMaxLifetime(time.Minute)
	return da, err
}
