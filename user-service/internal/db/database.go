package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error){
	connStr := "postgres://postgres:admin@localhost:5432/user_service"
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatalf("Error connecting to database: %s\n", err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatalf("DB connection is not alive: %s", err)
	}
	log.Println("Connected to the database!")
	return db, nil
}