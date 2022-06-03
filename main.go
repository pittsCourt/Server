package main

import (
	"net/http"
)

func main() {
	// Handling the /data/1
	http.HandleFunc("/data/1", firstHandler)

	// Handling the /data/2
	sHandler := secondHandler{}
	http.Handle("/data/2", sHandler)

	http.ListenAndServe(":8080", nil)

}

func firstHandler(res http.ResponseWriter, req *http.Request) {
	data := []byte("Data 1")
	res.Write(data)
}

type secondHandler struct{}

func (h secondHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Data 2")
	res.Write(data)
}
