package server

import (
	"net/http"
	"strconv"
	"time"

	"koala.pos/src/common"

	log "github.com/lfkeitel/verbose/v5"
	"gopkg.in/tylerb/graceful.v1"
)

type Server struct {
	e         *common.Environment
	routes    http.Handler
	address   string
	httpPort  string
	httpsPort string
}

func NewServer(e *common.Environment, routes http.Handler) *Server {
	serv := &Server{
		e:       e,
		routes:  routes,
		address: e.Config.Webserver.Address,
	}

	serv.httpPort = strconv.Itoa(e.Config.Webserver.Port)
	return serv
}

func (s *Server) Run() {
	log.Info("Starting web server...")

	log.WithFields(log.Fields{
		"address": s.address,
		"port":    s.httpPort,
	}).Debug("Starting HTTP server")

	timeout := 5 * time.Second

	if s.e.IsDev() {
		timeout = 1 * time.Millisecond
	}

	srv := &graceful.Server{
		Timeout: timeout,
		Server: &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         s.address + ":" + s.httpPort,
			Handler:      s.routes,
		},
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
