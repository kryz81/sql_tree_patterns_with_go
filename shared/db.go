package shared

import (
	"database/sql"
	"fmt"
)

func Connect(dbSrv string, dbPort int, dbName string, user string, pass string) (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbSrv, dbPort, user, pass, dbName)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
