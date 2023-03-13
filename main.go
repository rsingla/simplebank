package main

import (
	"database/sql"
	"log"
	"simplebank/api"

	db "simplebank/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:simplebankpass@localhost:5432/postgres?sslmode=disable"
	serverAddress = ":8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Run(serverAddress)
}
