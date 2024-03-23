package main

import (
	db "bd/db/sqlc"
	"bd/server"
	"bd/utils"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cnf, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("unable to read config %v", err)
	}

	conn, err := sql.Open(cnf.DBDriver, cnf.DBSource)
	if err != nil {
		log.Fatalf("unable to connect db %v", err)
	}
	store := db.NewStore(conn)
	server, err := server.NewServer(cnf, store)
	if err != nil {
		log.Fatalf("Сервис асааж чадсангүй")
	}
	err = server.StartServer(cnf.ServerAddress)
	if err != nil {
		log.Fatalf("cant start server err: %v", err)
	}

}
