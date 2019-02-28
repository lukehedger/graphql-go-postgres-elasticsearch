package api

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

const (
	host    = "localhost"
	port    = 5432
	user    = "luke"
	dbname  = "luke"
	sslmode = "disable"
)

func OpenDB() (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=%s",
		host, port, user, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}
	m.Steps(1)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to db!")

	return &DB{db}, nil
}
