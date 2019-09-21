package middleware

import (
	"net/http"

	"koala.pos/src/auth"
	"koala.pos/src/common"
)

// CheckAuth is middleware to check if a user is logged in, if not it will redirect to the login page
func CheckAuthAPI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Add("Authorization", "Basic realm=\"Koala POS\"")
			common.NewAPIResponse("Login required", nil).WriteResponse(w, http.StatusUnauthorized)
			return
		}

		if !auth.CheckLogin(username, password, r) {
			w.Header().Add("Authorization", "Basic realm=\"Koala POS\"")
			common.NewAPIResponse("Invalid username or password", nil).WriteResponse(w, http.StatusUnauthorized)
			return
		}

		// Continue
		next.ServeHTTP(w, r)
		return
	})
}
