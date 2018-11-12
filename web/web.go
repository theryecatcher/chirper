package web

import (
	"context"
	"net/http"

	"github.com/distsys-project/web/contentd"
)

type Web struct {
	srv *http.Server

	contentDaemon *contentd.Contentd
}

func New(cfg *Config) (*Web, error) {
	mx := http.NewServeMux()
	s := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: mx,
	}

	ws := &Web{
		srv: s,
	}

	mx.Handle("/static", http.FileServer(http.Dir("/static")))
	mx.HandleFunc("/index", ws.Index)
	mx.HandleFunc("/login", ws.Login)
	// mx.HandleFunc("/register", ws.Register)

	return ws, nil
}

func (w *Web) Start() error {
	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}
