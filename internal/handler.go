package api

import (
	"net/http"
	"time"
)

type Handler interface {
	HandleTest(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) HandleTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	// simulate work
	time.Sleep(100 * time.Millisecond)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
