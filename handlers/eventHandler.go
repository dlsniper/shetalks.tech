package handlers

import (
	"database/sql"
	"net/http"
)

// EventIndex GET /event
func EventIndex(db *sql.DB) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventCreate GET /event/create
func EventCreate(db *sql.DB) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventStore POST /event
func EventStore(db *sql.DB) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventShow GET /event/{id}
func EventShow(db *sql.DB) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventEdit GET /event/{id}/edit
func EventEdit(db *sql.DB) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventUpdate PUT/PATCH /event/{id}
func EventUpdate(db *sql.DB) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}

// EventDestroy DELETE /event/{id}
func EventDestroy(db *sql.DB) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pull up model things
	})

}
