package routes

import (
	"database/sql"

	"github.com/amy/shetalks.tech/handlers"
	"github.com/gorilla/mux"
)

// NewRouter router
func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/event").
		Name("EventIndex").
		Handler(handlers.EventIndex(db))

	router.
		Methods("GET").
		Path("/event/create").
		Name("EventCreate").
		Handler(handlers.EventCreate(db))

	router.
		Methods("POST").
		Path("/event").
		Name("EventStore").
		Handler(handlers.EventStore(db))

	router.
		Methods("GET").
		Path("/event/{eventID}").
		Name("EventShow").
		Handler(handlers.EventShow(db))

	router.
		Methods("GET").
		Path("/event/{eventID}/edit").
		Name("EventEdit").
		Handler(handlers.EventEdit(db))

	router.
		Methods("PUT").
		Path("/event/{eventID}").
		Name("EventUpdate").
		Handler(handlers.EventUpdate(db))

	router.
		Methods("DELETE").
		Path("/event/{eventID}").
		Name("EventDestroy").
		Handler(handlers.EventDestroy(db))

	/*for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc(db)
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}*/
	return router
}

// ConfigureRouter @TODO move router configuration stuff above down here
/*func ConfigureRouter() *mux.Router {

}*/

// StartRouter start the router here
/*func StartRouter() {

}*/
