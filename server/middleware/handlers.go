package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"server/database"
	"server/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllContest(w http.ResponseWriter, r *http.Request) {

	contests, err := database.GetAllContest()
	if err != nil {
		log.Printf("Unable to get all contests. %v", err)
	}
	json.NewEncoder(w).Encode(contests)
}

func RegisterResult(w http.ResponseWriter, r *http.Request) {

	var register models.Protocol

	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		log.Printf("Unable to decode request body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	log.Print(register)

	database.AddProtocol(register)

	res := response{
		Message: "Protokoll registrerat",
	}

	json.NewEncoder(w).Encode(res)
}
