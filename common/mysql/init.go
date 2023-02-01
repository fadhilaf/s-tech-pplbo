package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Start(mysqlConnection string) *sql.DB {
  db, err := sql.Open("mysql", mysqlConnection)
  if err != nil {
    log.Fatalf("Error when connecting to database (%v): %v", mysqlConnection , err)
  }

  return db;
}
