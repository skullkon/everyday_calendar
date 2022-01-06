package v1

import (
	"calendar/internal/server/v1/handler"
	"calendar/internal/utils"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// post methods for work with events in calendar
	router.HandleFunc("/v1/create_event", utils.Logging(handler.CreateEvent)).Methods("POST")
	router.HandleFunc("/v1/update_event", utils.Logging(handler.UpdateEvent)).Methods("POST")
	router.HandleFunc("/v1/delete_event", utils.Logging(handler.DeleteEvent)).Methods("POST")

	// methods for get events
	router.HandleFunc("/v1/events_for_day", utils.Logging(handler.EventForDay)).Methods("GET")

	return router
}
