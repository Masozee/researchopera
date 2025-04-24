package main

import (
	"net/http"
)

// setupHealthCheck adds a /healthz endpoint for container health checks
func setupHealthCheck(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
