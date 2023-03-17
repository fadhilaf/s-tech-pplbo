package postgres

import (
	"database/sql"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Start(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error ketika menyambungkan database: %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error ketika membuat driver migrasi dengan instance database yang telah ada: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		// ado yg template folder path ny "internal/template/blabla" ado jg yg "file://internal/template/blabla"
		"file://common/postgres/migration",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Error ketika melakukan migrasi database: %v", err)
	}

	m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	return db
}
