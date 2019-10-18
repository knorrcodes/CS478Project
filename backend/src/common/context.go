package common

import (
	"context"
	"net/http"
)

type Key int

// Session values keys
const (
	SessionEnvKey      Key = 1
	StoreCollectionKey Key = 2
	AuthServerKey      Key = 3
)

// Environment

// GetEnvironmentFromContext retrieves the Environment from the current request.
func GetEnvironmentFromContext(r *http.Request) *Environment {
	if rv := r.Context().Value(SessionEnvKey); rv != nil {
		return rv.(*Environment)
	}
	return nil
}

// SetEnvironmentToContext sets an Environment for the current request.
func SetEnvironmentToContext(r *http.Request, e *Environment) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), SessionEnvKey, e))
}
