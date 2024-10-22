package internalsql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatalf("Error Connecting to Database: %v\n", err)
		return nil, err
	}
	return db, nil
}
