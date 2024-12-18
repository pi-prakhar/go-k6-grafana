package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/pi-prakhar/go-k6-grafana/internal"
	"github.com/pi-prakhar/go-k6-grafana/metrics"
)

// const (
// 	ADDR string = ":8080"
// )

func getPort() string {
	port := os.Getenv("PORT")
	portString := ":" + port
	return portString
}

func main() {
	ADDR := getPort()
	metrics.InitMetrics()

	handler := api.NewHandler()
	router := api.NewRouter(handler)

	srv := http.Server{
		Addr:    ADDR,
		Handler: router.Mux,
	}
	log.Println("server started at port ", ADDR)
	log.Fatal(srv.ListenAndServe().Error())
}
