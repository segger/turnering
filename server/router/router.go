package router

import (
	"github.com/gorilla/mux"
	"server/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/contests", middleware.GetAllContest).Methods("GET")

	return router
}
