package dependencies

import (
	"log"
	"database/sql"

	_ "github.com/lib/pq"
)

type Environment struct {
	DBConn *sql.DB  
}

func CreateEnvironment() *Environment {
	dbConn, err := sql.Open("postgres", "host=localhost port=5432 user=bstack_user password=password dbname=bstack_db sslmode=disable")
	if err != nil {
		log.Panicf("Error creating connection to database: %v", err)
	}

	env := Environment{
		DBConn: dbConn,
	}

	return &env
}

