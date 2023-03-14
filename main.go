package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	"simplebank/util"

	db "simplebank/sqlc"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Run(config.ServerAddress)
}
