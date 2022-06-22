package main

import (
	"log"
	"net/http"

	"github.com/digicert/health"
	"github.com/pittsCourt/Server/handlers"
)

func main() {
	http.HandleFunc("/health", health.HealthEndpoint)
	http.Handle("/metrics", health.Metrics())
	health.SetLogLevel("info")
	health.SetDebug(true)

	// Handling the /data/ as a function
	http.HandleFunc("/data/", handlers.DataHandler)

	log.Println("Listening on port :8080")

	http.ListenAndServe(":8080", nil)
}
