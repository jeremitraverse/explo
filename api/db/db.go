package db

import (
	"fmt"
	"os"

	"github.com/jeremitraverse/explo/scripts"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateConnection() (*sqlx.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", dbUser, dbName, dbPassword, dbHost)

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		fmt.Println("Error while connecting to db: %w", err)
		return nil, err
	}

	return db, nil
}

func Seed() {
	db, err := CreateConnection()

	if err != nil {
		fmt.Println("Error seeding database: ", err)
	}

	tx := db.MustBegin()

	tx.MustExec(scripts.DropTables())
	tx.MustExec(scripts.CreateTables())
	tx.MustExec(scripts.SeedTables())

	tx.Commit()

	db.Close()
}
