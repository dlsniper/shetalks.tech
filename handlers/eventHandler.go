package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

// EventIndex GET /event
func EventIndex(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HIT HIT HIT")
	})

}

// EventCreate GET /event/create
func EventCreate(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

}

// EventStore POST /event
func EventStore(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventShow GET /event/{id}
func EventShow(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventEdit GET /event/{id}/edit
func EventEdit(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventUpdate PUT/PATCH /event/{id}
func EventUpdate(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventDestroy DELETE /event/{id}
func EventDestroy(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}
