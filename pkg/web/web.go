// Package web implements functions for web rest api
package web

import (
	"net/http"

	"github.com/charmbracelet/log"
)

type Web struct {
	addr   string
	server *http.ServeMux
}

func New(port string) *Web {
	return &Web{
		addr:   ":" + port,
		server: http.NewServeMux(),
	}
}

func (w *Web) Start() error {
	log.Info("Start Server", "addr", w.addr)
	return http.ListenAndServe(w.addr, w.server)
}

func (w *Web) Get(pattern string, h HandlerFn) {
	w.server.HandleFunc("GET "+pattern, mainHandler(h))
}

func (w *Web) Post(pattern string, h HandlerFn) {
	w.server.HandleFunc("POST "+pattern, mainHandler(h))
}

func (w *Web) Put(pattern string, h HandlerFn) {
	w.server.HandleFunc("PUT "+pattern, mainHandler(h))
}

func (w *Web) Patch(pattern string, h HandlerFn) {
	w.server.HandleFunc("PATCH "+pattern, mainHandler(h))
}

func (w *Web) Delete(pattern string, h HandlerFn) {
	w.server.HandleFunc("DELETE "+pattern, mainHandler(h))
}

func (w *Web) Option(pattern string, h HandlerFn) {
	w.server.HandleFunc("OPTIONS "+pattern, mainHandler(h))
}

func (w *Web) Group(pattern string, h http.Handler) {
	group := http.NewServeMux()
	group.Handle(pattern+"/", http.StripPrefix(pattern, h))
}

type Server func(w *Web)

func (w *Web) Add(server Server) {
	server(w)
}
