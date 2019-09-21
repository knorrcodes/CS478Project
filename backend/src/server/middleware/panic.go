package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	log "github.com/lfkeitel/verbose/v5"
	"koala.pos/src/common"
)

func Panic(next http.Handler, e *common.Environment) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(runtime.Error); ok {
					buf := make([]byte, 2048)
					runtime.Stack(buf, false)
					log.WithFields(log.Fields{
						"package": "middleware:panic",
						"stack":   string(buf),
					}).Alert("")
				}
				log.Alert(fmt.Sprintf("%v", r))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
