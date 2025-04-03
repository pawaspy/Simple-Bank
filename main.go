package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/pawaspy/simple_bank/api"
	db "github.com/pawaspy/simple_bank/db/sqlc"
	"github.com/pawaspy/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot open the config file")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(config.Address)
	if err != nil {
		log.Fatal("Cannot start the server")
	}
}
