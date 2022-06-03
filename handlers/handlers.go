package handlers

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Id    int    //`json:"id"`
	Value string //`json:"value"`
}

func firstHandler(res http.ResponseWriter, req *http.Request) {
	m := Data{1, "one"}
	b, _ := json.Marshal(m)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(200)
	res.Write(b)
}
