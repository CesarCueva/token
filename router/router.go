package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Controller - Handlers interface
type Controller interface {
	IndexRoute(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
}

// NewRouter - Creates a router
func NewRouter(
	controller Controller,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.IndexRoute)
	router.HandleFunc("/getToken", controller.GetToken).Methods(http.MethodGet)

	return router
}
