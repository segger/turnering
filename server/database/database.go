package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"server/models"

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
