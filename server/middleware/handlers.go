package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"server/database"
	"server/models"

	"github.com/gorilla/mux"
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

func GetAllEnabledContest(w http.ResponseWriter, r *http.Request) {
	contests, err := database.GetAllContest()
	if err != nil {
		log.Printf("Unable to get all contests. %v", err)
	}
	var enabledContests []models.Contest
	for _, contest := range contests {
		if contest.Enabled {
			enabledContests = append(enabledContests, contest)
		}
	}

	json.NewEncoder(w).Encode(enabledContests)
}

func GetContestById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contestId := params["contestId"]

	contest, err := database.GetContestById(contestId)
	if err != nil {
		log.Printf("Unable to get contest. %v", err)
	}
	json.NewEncoder(w).Encode(contest)
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

func GetRegistered(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contestId := params["contestId"]

	results, err := database.GetRegisteredByContestId(contestId)
	if err != nil {
		log.Printf("Unable to get result. %v", err)
	}
	json.NewEncoder(w).Encode(results)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contestId := params["contestId"]

	results, err := database.GetResultByContestId(contestId)
	if err != nil {
		log.Printf("Unable to get result. %v", err)
	}
	json.NewEncoder(w).Encode(results)
}
