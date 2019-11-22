package middleware

import (
	"net/http"
	"strconv"
	"strings"

	log "github.com/lfkeitel/verbose/v5"

	"koala.pos/src/auth"
	"koala.pos/src/models"
	"koala.pos/src/models/stores"
)

// CheckAuth is middleware to check if a user is logged in, if not it will redirect to the login page
func CheckAuthGraphQL(stores *stores.StoreCollection, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(authHeader) != 2 {
			writeAuthErrorResponse(w)
			return
		}

		log.Debug(authHeader[0])

		switch authHeader[0] {
		case "Server":
			serverAuth(w, r, authHeader[1], stores.Server, next)
		case "Customer":
			customerAuth(w, r, authHeader[1], stores.CustCode, next)
		default:
			writeAuthErrorResponse(w)
		}
	})
}

func serverAuth(w http.ResponseWriter, r *http.Request, code string, servers stores.ServerStore, next http.Handler) {
	serverID, err := strconv.Atoi(code)
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
}

func customerAuth(w http.ResponseWriter, r *http.Request, code string, custcodeStore stores.CustCodeStore, next http.Handler) {
	custcode, err := custcodeStore.GetCustCodeByCode(code)
	if custcode == nil || err != nil {
		writeAuthErrorResponse(w)
		return
	}

	server := models.NewServer(nil)
	server.Name = "Customer"
	server.CustCode = code

	r = auth.SetServerToContext(r, server)

	// Continue
	next.ServeHTTP(w, r)
}

func writeAuthErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"errors": [{"message": "Authentication failure, are you using the Authorization header?"}], "data": null}`))
}
