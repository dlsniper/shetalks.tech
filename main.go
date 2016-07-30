package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amy/shetalks.tech/config/database"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("blahblah\n")

	database.InitDB()

}
