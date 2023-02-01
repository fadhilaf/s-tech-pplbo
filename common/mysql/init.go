package mysql

import (
	"database/sql"
	"log"

  "github.com/golang-migrate/migrate/v4"
  "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Start(mysqlConnection string) *sql.DB {
  db, err := sql.Open("mysql", mysqlConnection)
  if err != nil {
    log.Fatalf("Error when connecting to database: %v", err)
  }

  driver, err := mysql.WithInstance(db, &mysql.Config{})
  m, err := migrate.NewWithDatabaseInstance(
    "file:///common/mysql/migration", 
    "postgres", driver )

  m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

  return db;
}
