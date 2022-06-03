package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pittsCourt/Server1/handlers"
)

func main() {
	// Handling the /data/1 as a function
	http.HandleFunc("/data/1", handlers.FirstHandler)

	// Handling the /data/2 as a type
	sHandler := secondHandler{}
	http.Handle("/data/2", sHandler)

	log.Println("Listening on port :8080")

	http.ListenAndServe(":8080", nil)
}

type secondHandler struct{}

func (h secondHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	m := Data{2, "two"}
	b, _ := json.Marshal(m)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(200)
	res.Write(b)
}
