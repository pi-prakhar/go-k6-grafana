package main

import (
	"log"
	"net/http"

	api "github.com/pi-prakhar/go-k6-grafana/internal"
)

const (
	ADDR string = ":8080"
)

func main() {
	handler := api.NewHandler()
	router := api.NewRouter(handler)

	srv := http.Server{
		Addr:    ADDR,
		Handler: router.Mux,
	}
	log.Println("server started at port ", ADDR)
	log.Fatal(srv.ListenAndServe().Error())
}
