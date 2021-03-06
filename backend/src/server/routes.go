package server

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime"

	"github.com/99designs/gqlgen/handler"
	"github.com/julienschmidt/httprouter"
	log "github.com/lfkeitel/verbose/v5"

	"koala.pos/src/common"
	"koala.pos/src/graphql"
	"koala.pos/src/models/stores"
	mid "koala.pos/src/server/middleware"
)

func LoadRoutes(e *common.Environment, stores *stores.StoreCollection) http.Handler {
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(notFoundHandler)

	// GraphQL routes
	if e.IsDev() {
		r.HandlerFunc("GET", "/", handler.Playground("GraphQL playground", "/graphql"))
		log.Info("GraphQL playground enabled")
	}

	r.Handler("POST", "/graphql",
		mid.CheckAuthGraphQL(stores,
			mid.SetSessionInfo(e, stores,
				handler.GraphQL(
					graphql.NewExecutableSchema(graphql.NewConfig()),
				))))

	if e.IsDev() {
		r.Handler("GET", "/debug/*a", debugRouter(e))
		log.Info("Profiling enabled")
	}

	h := mid.Logging(r, e) // Logging
	h = mid.Panic(h, e)    // Panic catcher
	return h
}

func debugRouter(e *common.Environment) http.Handler {
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(notFoundHandler)

	r.HandlerFunc("GET", "/debug/pprof", pprof.Index)
	r.HandlerFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)
	r.HandlerFunc("GET", "/debug/pprof/profile", pprof.Profile)
	r.HandlerFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
	r.HandlerFunc("GET", "/debug/pprof/trace", pprof.Trace)
	// Manually add support for paths linked to by index page at /debug/pprof/
	r.Handler("GET", "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handler("GET", "/debug/pprof/heap", pprof.Handler("heap"))
	r.Handler("GET", "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handler("GET", "/debug/pprof/block", pprof.Handler("block"))

	r.HandlerFunc("GET", "/debug/heap-stats", heapStats)

	return r
}

func heapStats(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Fprintf(w,
		"HeapSys: %d, HeapAlloc: %d, HeapIdle: %d, HeapReleased: %d\n",
		m.HeapSys,
		m.HeapAlloc,
		m.HeapIdle,
		m.HeapReleased,
	)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	common.NewEmptyAPIResponse().WriteResponse(w, http.StatusNotFound)
}
