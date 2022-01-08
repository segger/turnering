package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"server/database"
)

func GetAllContest(w http.ResponseWriter, r *http.Request) {

	contests, err := database.GetAllContest()
	if err != nil {
		log.Fatalf("Unable to get all contests. %v", err)
	}
	json.NewEncoder(w).Encode(contests)
}
