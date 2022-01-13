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

func GetRegisteredByContestId(id string) ([]models.ContestRegistered, error) {
	db := createConnection()

	defer db.Close()

	var results []models.ContestRegistered
	sqlStatement := `SELECT first_name, dog_name, class_nbr, event_name ` +
		`FROM result JOIN participant ON result.participant_id=participant.id WHERE result.contest_id=$1 ORDER BY result.created_at desc`
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

func GetResultByContestId(id string) ([]models.ContestResult, error) {
	db := createConnection()

	defer db.Close()

	var results []models.ContestResult
	sqlStatement := `SELECT first_name, dog_name, class_nbr, event_name, points, errors, time, sse ` +
		`FROM result JOIN participant ON result.participant_id=participant.id WHERE result.contest_id=$1`
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

func AddProtocol(protocol models.Protocol) {
	participantId := createParticipant(protocol.Participant)
	for _, result := range protocol.EventResultList {
		createResult(participantId, protocol.ContestId, result)
	}
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

func createResult(participant string, contest string, result models.EventResult) string {
	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO result
		(id, created_at, deleted, participant_id, contest_id, 
		event_name, points, errors, time, sse)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	resultId := uuid.New()
	now := time.Now()
	participantId, err := uuid.Parse(participant)
	contestId, err := uuid.Parse(contest)

	_, err = db.Exec(sqlStatement, resultId, now, false, participantId, contestId,
		result.EventName, result.Points, result.Errors, result.Time, result.Sse)
	if err != nil {
		panic(err)
	}

	return resultId.String()
}
