package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

  "github.com/golang-migrate/migrate/v4"
  "github.com/golang-migrate/migrate/v4/database/mysql"
)

func Start(mysqlConnection string) *sql.DB {
  db, err := sql.Open("mysql", mysqlConnection)
  if err != nil {
    log.Fatalf("Error when connecting to database (%v): %v", mysqlConnection , err)
  }

  driver, err := mysql.WithInstance(db, &mysql.Config{})
  if err != nil {
    log.Fatalf("Error when creating migration driver (%v): %v", mysqlConnection, err)
  }

  m, err := migrate.NewWithDatabaseInstance(
    "common/mysql/migration", 
    "mysql", driver )
  if err != nil {
    log.Fatalf("Error when migration (%v): %v", mysqlConnection, err)
  }

  m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

  return db;
}
