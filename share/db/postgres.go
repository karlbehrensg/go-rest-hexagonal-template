package share

import (
	"database/sql"
	"fmt"
)

func NewPostgresClient(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		db.Close()
		panic(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		db.Close()
		panic(err)
	}

	fmt.Printf("Connected to %s:%d/%s successfully!\n", host, port, dbname)

	return db, nil
}
