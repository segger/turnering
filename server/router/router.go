package router

import (
	"server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/contests", middleware.GetAllContest).Methods("GET")

	router.HandleFunc("/api/register", middleware.RegisterResult).Methods("POST")

	return router
}
