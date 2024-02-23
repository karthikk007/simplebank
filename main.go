package main

import (
	"database/sql"
	"log"

	"example.com/simplebank/api"
	db "example.com/simplebank/db/sqlc"
	"example.com/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdderss)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}

}
