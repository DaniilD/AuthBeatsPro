package internal

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

// надо вынести все переменные в конфиг
func (server *Server) Run(host string, port string, router *Router) error {
	server.httpServer = &http.Server{
		Addr:           host + ":" + port,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 + time.Second,
		WriteTimeout:   10 + time.Second,
		Handler:        router.muxRouter,
	}

	return server.httpServer.ListenAndServe()
}

func (server *Server) Shutdown(context context.Context) error {
	return server.httpServer.Shutdown(context)
}
