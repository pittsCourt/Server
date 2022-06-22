package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/digicert/health"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// type Data struct {
// 	Id    int    //`json:"id"`
// 	Value string //`json:"value"`
// }

var db *sql.DB

func FirstHandler(res http.ResponseWriter, req *http.Request) {
	// Commenting out the use of the above struct,
	// this was used previously, but now will be using mariadb
	// to provide json data.
	// m := Data{1, "one"}
	// b, _ := json.Marshal(m)

	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		//Addr:                 "172.17.0.0:3306",
		DBName:               "server_data",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		health.Fatal("This is the error: %v", err)
	}

	var s sql.NullString
	err = db.QueryRow("SELECT attr FROM json_data WHERE id = ?", 1).Scan(&s)
	if err != nil {
		health.Fatal("This is the error: %v", err)
	}
	result, err := json.Marshal(s.String)
	if err != nil {
		health.Fatal("This is the error: %v", err)
	}
	resultStr := string(result)
	resultByte := []byte((strings.Replace(resultStr, "\\", "", 6)))

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(200)
	res.Write(resultByte)
}

type SecondHandler struct{}

func (h SecondHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Commenting out the use of the above struct,
	// this was used previously, but now will be using mariadb
	// to provide json data.
	// m := Data{2, "two"}
	// b, _ := json.Marshal(m)

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "server_data",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	var s sql.NullString
	err = db.QueryRow("SELECT attr FROM json_data WHERE id = ?", 2).Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	result, err := json.Marshal(s.String)
	if err != nil {
		health.Fatal("This is the error: %v", err)
	}
	resultStr := string(result)
	resultByte := []byte((strings.Replace(resultStr, "\\", "", 6)))

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(200)
	res.Write(resultByte)
}
