package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/digicert/health"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Data struct {
	Id    int    //`db:"id" json:"id"`
	Value string //`db:"value" json:"value"`
}

var db *sqlx.DB

func DataHandler(res http.ResponseWriter, req *http.Request) {
	// Declare s variable to the url path after /data/
	s := req.URL.Path[len("/data/"):]
	if s == "" {
		health.Error("Path after '/data/' is empty.")
		res.WriteHeader(500)
		res.Write([]byte("No data to display, try a number after '/data/' in the URL"))
		return
	}
	health.Debug("Path contains a value %s", s)

	var err error
	db, err = sqlx.Open("mysql", "root:root@tcp(mariadb:3306)/server_data")
	if err != nil {
		health.Error("This is the error: %v", err)
	}
	health.Info("We've reached this point.")

	var myData []Data
	err = db.Select(&myData, "SELECT * FROM json_data WHERE id = ?", s)
	if err != nil {
		health.Error("This is the error: %v", err)
	}

	var result []byte
	if len(myData) > 0 {
		result, err = json.Marshal(myData[0])

		if err != nil {
			health.Error("This is the error: %v", err)
		}
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(200)
	res.Write(result)
}
