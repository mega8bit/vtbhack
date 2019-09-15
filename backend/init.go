package main

import (
	"database/sql"
	"fmt"
	"os"
)

func init() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "195.181.245.183"
	}
	dbsrc := fmt.Sprintf("user=vtbhack password=Koo6ahghok3cahGu9cae dbname=vtbhack host=%s port=5432 sslmode=disable", dbHost)
	db, err = sql.Open("postgres", dbsrc)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
