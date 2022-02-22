package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"server/models"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nosework"
	password = "secret"
	dbname   = "nosework"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	psqlConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConfig)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func GetAllContest() ([]models.Contest, error) {
	db := createConnection()

	defer db.Close()

	var contests []models.Contest

	sqlStatement := `SELECT * FROM contest`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Printf("Unable to execute query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var contest models.Contest

		err = rows.Scan(&contest.Id, &contest.Name, &contest.Enabled)
		if err != nil {
			log.Printf("Unable to scan row. %v", err)
		}
		contests = append(contests, contest)
	}
	return contests, err
}

func GetContestById(id string) (models.Contest, error) {
	db := createConnection()

	defer db.Close()

	var contest models.Contest
	sqlStatement := `SELECT * FROM contest WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&contest.Id, &contest.Name, &contest.Enabled)
	switch err {
	case sql.ErrNoRows:
		return contest, nil
	case nil:
		return contest, nil
	default:
		log.Printf("Unable to scan row %v", err)
	}
	return contest, err
}

func GetEventByContestId(id string) ([]models.Event, error) {
	db := createConnection()

	defer db.Close()

	var results []models.Event
	sqlStatement := `SELECT id, name, class_nbr, sort_order FROM event WHERE contest_id=$1`
	rows, err := db.Query(sqlStatement, id)

	for rows.Next() {
		var result models.Event

		err = rows.Scan(&result.Id, &result.Name, &result.ClassNbr, &result.SortOrder)
		if err != nil {
			log.Printf("Unable to scan row. %v", err)
		}
		results = append(results, result)
	}
	return results, err
}

func GetRegisteredByContestId(id string) ([]models.ContestRegistered, error) {
	db := createConnection()

	defer db.Close()

	var results []models.ContestRegistered
	sqlStatement := `SELECT first_name, dog_name, participant.class_nbr, event.name ` +
		`FROM result JOIN participant ON result.participant_id=participant.id ` +
		`JOIN event ON result.event_id=event.id WHERE result.contest_id=$1 AND result.deleted = false ORDER BY result.created_at desc`
	rows, err := db.Query(sqlStatement, id)

	for rows.Next() {
		var result models.ContestRegistered

		err = rows.Scan(&result.FirstName, &result.DogName, &result.ClassNbr, &result.EventName)
		if err != nil {
			log.Printf("Unable to scan row. %v", err)
		}
		results = append(results, result)
	}
	return results, err
}

func GetContestParticipants(contestId string) ([]models.Participant, error) {
	db := createConnection()

	defer db.Close()

	var participants []models.Participant
	sqlStatement := `SELECT DISTINCT participant.id, first_name, dog_name, class_nbr FROM result JOIN participant ON result.participant_id=participant.id WHERE result.contest_id=$1`
	rows, err := db.Query(sqlStatement, contestId)

	for rows.Next() {
		var result models.Participant

		err = rows.Scan(&result.Id, &result.FirstName, &result.DogName, &result.ClassNbr)
		if err != nil {
			log.Printf("Unable to scan row. %v", err)
		}
		participants = append(participants, result)
	}
	return participants, err
}

func GetResultByContestId(id string) ([]models.ContestResult, error) {
	db := createConnection()

	defer db.Close()

	var results []models.ContestResult
	sqlStatement := `SELECT first_name, dog_name, participant.class_nbr, event.name, points, errors, time, sse ` +
		`FROM result JOIN participant ON result.participant_id=participant.id JOIN event ON result.event_id=event.id ` +
		`WHERE result.contest_id=$1 AND result.deleted = false`
	rows, err := db.Query(sqlStatement, id)

	for rows.Next() {
		var result models.ContestResult

		err = rows.Scan(&result.FirstName, &result.DogName, &result.ClassNbr, &result.EventName, &result.Points, &result.Errors, &result.Time, &result.Sse)
		if err != nil {
			log.Printf("Unable to scan row. %v", err)
		}
		results = append(results, result)
	}
	return results, err
}

func GetResultByContestIdAndParticipantId(contestId string, participantId string) ([]models.EventResult, error) {
	db := createConnection()

	defer db.Close()

	var results []models.EventResult
	sqlStatement := `SELECT result.id, event_id, points, errors, time, sse ` +
		`FROM result JOIN event ON result.event_id=event.id WHERE result.contest_id=$1 AND participant_id=$2`
	rows, err := db.Query(sqlStatement, contestId, participantId)

	for rows.Next() {
		var result models.EventResult

		err = rows.Scan(&result.ResultId, &result.EventId, &result.Points, &result.Errors, &result.Time, &result.Sse)
		if err != nil {
			log.Printf("Unable to scan row. %v", err)
		}
		results = append(results, result)
	}
	return results, err
}

func AddProtocol(protocol models.Protocol) {
	participantId := AddParticipantIfNotExists(protocol.Participant).Id
	for _, result := range protocol.EventResultList {
		upsertResult(participantId, protocol.ContestId, result)
	}
}

func AddParticipantIfNotExists(participant models.Participant) models.Participant {
	participantByName, err := getParticipantByNames(participant.FirstName, participant.DogName, participant.ClassNbr)
	if err != nil {
		panic(err)
	}
	if participantByName.Id != "" {
		log.Printf("Participant exists with id %s", participantByName.Id)
		return participantByName
	} else {
		log.Printf("Creating participant...")
		participantId := createParticipant(participant)
		participant.Id = participantId
		return participant
	}
}

func getParticipantByNames(firstName string, dogName string, classNbr int8) (models.Participant, error) {
	db := createConnection()

	defer db.Close()

	var participant models.Participant
	sqlStatement := `SELECT id, first_name, dog_name, class_nbr ` +
		`FROM participant WHERE first_name=$1 AND dog_name=$2 AND class_nbr=$3 AND deleted=false`

	row := db.QueryRow(sqlStatement, firstName, dogName, classNbr)
	err := row.Scan(&participant.Id, &participant.FirstName, &participant.DogName, &participant.ClassNbr)
	switch err {
	case sql.ErrNoRows:
		return participant, nil
	case nil:
		return participant, nil
	default:
		log.Printf("Unable to scan row %v", err)
	}
	return participant, err
}

func createParticipant(participant models.Participant) string {
	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO participant (id, first_name, dog_name, class_nbr) VALUES ($1, $2, $3, $4)`

	participantId := uuid.New()
	_, err := db.Exec(sqlStatement, participantId, participant.FirstName, participant.DogName, participant.ClassNbr)
	if err != nil {
		panic(err)
	}

	return participantId.String()
}

func upsertResult(participant string, contest string, result models.EventResult) string {
	// log.Printf("result id %v", result.ResultId)
	if result.ResultId == "null" {
		return createResult(participant, contest, result)
	} else {
		return updateResult(participant, contest, result)
	}
}

func updateResult(participant string, contest string, result models.EventResult) string {
	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE result SET created_at=$4, ` +
		`event_id=$5, points=$6, errors=$7, time=$8, sse=$9` +
		`WHERE id=$1 AND participant_id=$2 AND contest_id=$3`

	participantId, err := uuid.Parse(participant)
	contestId, err := uuid.Parse(contest)

	now := time.Now()
	_, err = db.Exec(sqlStatement, result.ResultId, participantId, contestId, now,
		result.EventId, result.Points, result.Errors, result.Time, result.Sse)
	if err != nil {
		panic(err)
	}

	return result.ResultId
}

func createResult(participant string, contest string, result models.EventResult) string {
	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO result
		(id, created_at, deleted, participant_id, contest_id, 
		event_id, points, errors, time, sse)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	resultId := uuid.New()
	now := time.Now()
	participantId, err := uuid.Parse(participant)
	contestId, err := uuid.Parse(contest)

	_, err = db.Exec(sqlStatement, resultId, now, false, participantId, contestId,
		result.EventId, result.Points, result.Errors, result.Time, result.Sse)
	if err != nil {
		panic(err)
	}

	return resultId.String()
}
