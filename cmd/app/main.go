package main

import (
	"database/sql"
	//"log"
	//"net/http"

	//"github.com/amy/shetalks.tech/routes"
	//"github.com/amy/shetalks.tech/config/database"
	"github.com/amy/shetalks.tech/models"
)

func main() {

	var db *sql.DB

	//router := routes.NewRouter(db)

	//log.Fatal(http.ListenAndServe(":8080", router))

	//fmt.Println("blahblah\n")

	//database.InitDB(db)

	event := models.CreateEvent(db, "Event", "Description", ["speaker1", "speaker2"]string)

}
