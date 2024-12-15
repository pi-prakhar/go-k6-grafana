package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router struct {
	Mux     *http.ServeMux
	handler Handler
}

func NewRouter(handler Handler) *Router {
	router := Router{
		Mux:     http.NewServeMux(),
		handler: handler,
	}
	router.initRouter()
	return &router
}

func (m *Router) initRouter() {
	m.Mux.HandleFunc("/get", m.handler.HandleTest)
	m.Mux.Handle("/metrics", promhttp.Handler())
}
