package main

import (
	"database/sql"
	"log"
	"os"
	"tusa/internal/router"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgresql://heroku:secrect@localhost:54328/heroku?sslmode=disable"
	}
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "tusa", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	server := router.SetupServer()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":32958"
	} else {
		port = ":" + port
	}
	server.Run(port)
}
