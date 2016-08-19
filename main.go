package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/amy/shetalks.tech/routes"
)

func main() {

	var db *sql.DB

	router := routes.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))

	//fmt.Println("blahblah\n")

	//database.InitDB()

}
