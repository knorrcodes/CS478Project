package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"koala.pos/src/auth"
	"koala.pos/src/models/stores"
)

// CheckAuth is middleware to check if a user is logged in, if not it will redirect to the login page
func CheckAuthGraphQL(servers stores.ServerStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(authHeader) != 2 || authHeader[0] != "Server" {
			writeAuthErrorResponse(w)
			return
		}

		serverID, err := strconv.Atoi(authHeader[1])
		if err != nil {
			writeAuthErrorResponse(w)
			return
		}

		server, err := servers.GetByCode(serverID)
		if err != nil || server == nil {
			writeAuthErrorResponse(w)
			return
		}

		r = auth.SetServerToContext(r, server)

		// Continue
		next.ServeHTTP(w, r)
		return
	})
}

func writeAuthErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"errors": [{"message": "Authentication failure, are you using the Authorization header?"}], "data": null}`))
}
