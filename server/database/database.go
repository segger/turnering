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
	port     = 5555
	user     = "nosework"
	password = "secret"
	dbname   = "nosework"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	password := os.Getenv("DB_PASSWORD")
	psqlConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
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
		log.Fatalf("Unable to execute query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var contest models.Contest

		err = rows.Scan(&contest.Id, &contest.Name)
		if err != nil {
			log.Fatalf("Unable to scan row. %v", err)
		}
		contests = append(contests, contest)
	}
	return contests, err
}

func AddProtocol(protocol models.Protocol) {
	contestId := "27eec793-a08d-4b00-969a-e7379e4b2643"
	participantId := createParticipant(protocol.Participant)
	for _, result := range protocol.EventResultList {
		createResult(participantId, contestId, result)
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
