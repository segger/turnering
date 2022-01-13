package router

import (
	"server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/contests", middleware.GetAllContest).Methods("GET")
	router.HandleFunc("/api/contests/enabled", middleware.GetAllEnabledContest).Methods("GET")
	router.HandleFunc("/api/contests/{contestId:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}}", middleware.GetContestById).Methods("GET")

	router.HandleFunc("/api/register", middleware.RegisterResult).Methods("POST")

	router.HandleFunc("/api/registered/{contestId:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}}", middleware.GetRegistered).Methods("GET")
	return router
}
