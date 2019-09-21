package middleware

import (
	"net/http"

	"koala.pos/src/common"
)

func SetSessionInfo(next http.Handler, e *common.Environment) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = common.SetEnvironmentToContext(r, e)

		// If running behind a proxy, set the RemoteAddr to the real address
		if r.Header.Get("X-Real-IP") != "" {
			r.RemoteAddr = r.Header.Get("X-Real-IP")
		}

		next.ServeHTTP(w, r)
	})
}
