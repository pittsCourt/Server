package main

import (
	"log"
	"net/http"

	"github.com/pittsCourt/Server/handlers"
)

func main() {
	// Handling the /data/1 as a function
	http.HandleFunc("/data/1", handlers.FirstHandler)

	// Handling the /data/2 as a type
	sHandler := handlers.SecondHandler{}
	http.Handle("/data/2", sHandler)

	log.Println("Listening on port :8080")

	http.ListenAndServe(":8080", nil)
}
