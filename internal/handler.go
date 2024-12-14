package api

import "net/http"

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

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
