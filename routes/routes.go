package routes

import (
	"net/http"

	"github.com/amy/shetalks.tech/handlers/handlers"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},

	/**
	 * Event Routes
	 */

	Route{
		"EventIndex",
		"GET",
		"/event",
		handlers.EventIndex,
	},
	Route{
		"EventCreate",
		"GET",
		"/event/create",
		handlers.EventCreate,
	},
	Route{
		"EventStore",
		"POST",
		"/event",
		handlers.EventStore,
	},
	Route{
		"EventShow",
		"GET",
		"/event/{eventID}",
		handlers.EventShow,
	},
	Route{
		"EventEdit",
		"GET",
		"/event/{eventID}/edit",
		handlers.Eventedit,
	},
	Route{
		"EventUpdate",
		"PUT",
		"/event/{eventID}",
		handlers.EventUpdate,
	},
	Route{
		"EventDestroy",
		"DELETE",
		"/event/{eventID}",
		handlers.EventDestroy,
	},

	/**
	 * Organization Routes
	 */

	/**
	 * Speaker Routes
	 */
}
