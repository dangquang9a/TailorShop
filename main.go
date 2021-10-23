package main

import (
	"TailorShop/api"
	db "TailorShop/db/sqlc"
	"TailorShop/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://merlin:@merlin123@10.0.0.223:5432/tailor_shop?sslmode=disable"
// 	serverAddress = "localhost:8080"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
