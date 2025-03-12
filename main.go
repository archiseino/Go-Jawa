package main

import (
	"database/sql"
	"log"

	// Package Db and API
	_ "github.com/lib/pq" // Package for connecting to Postgres
	api "github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
)

const (
	dbDriver= "postgres"
	dbSource= "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress= "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}


}



