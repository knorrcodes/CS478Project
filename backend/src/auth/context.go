package auth

import (
	"context"
	"net/http"

	"koala.pos/src/common"
	"koala.pos/src/models"
)

// GetServerFromContext retrieves the Server from the current request.
func GetServerFromContext(ctx context.Context) *models.Server {
	if rv := ctx.Value(common.AuthServerKey); rv != nil {
		return rv.(*models.Server)
	}
	return nil
}

// SetServerToContext sets a Server for the current request.
func SetServerToContext(r *http.Request, e *models.Server) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), common.AuthServerKey, e))
}
