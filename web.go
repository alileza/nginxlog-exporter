package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type web struct {
	mux *http.ServeMux

	listenErrChan chan error
	listenAddress string
}

func NewWebServer() *web {
	w := &web{
		mux:           http.NewServeMux(),
		listenErrChan: make(chan error),
		listenAddress: cfg.ListenAddress,
	}

	w.mux.Handle(cfg.MetricPath, promhttp.Handler())

	w.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
      <head><title>Nginx Log Exporter</title></head>
      <body>
      <h1>Nginx Log Exporter</h1>
      <p><a href="` + cfg.MetricPath + `">Metrics</a></p>
      </body>
      </html>`))
	})

	return w
}

func (w *web) Run() {
	log.Printf("Web Service is Listening on %s\n", w.listenAddress)
	w.listenErrChan <- http.ListenAndServe(w.listenAddress, w.mux)
}

func (w *web) ListenError() <-chan error {
	return w.listenErrChan
}
