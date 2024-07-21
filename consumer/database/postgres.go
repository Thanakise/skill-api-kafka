package database

import (
	"database/sql"
	"log"
	"os"
)

func InitDatabase() Db{
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	return Db{
		DB: db,
	}
}

func (database Db) CloseDatabase(){
	database.DB.Close()
}

