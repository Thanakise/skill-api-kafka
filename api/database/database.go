package database

import (
	"database/sql"
)

type Db struct{
	DB *sql.DB
}

type DatabaseInterface interface{
	InitDatabase()
	CloseDatabase()
}