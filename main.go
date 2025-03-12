package main

import (
	"database/sql"
	"log"

	// Package Db and API
	_ "github.com/lib/pq" // Package for connecting to Postgres
	api "github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
	util "github.com/techschool/simplebank/util"
)


func main() {
	// Load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Connect to the database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}


}



