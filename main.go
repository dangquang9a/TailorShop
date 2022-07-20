package main

import (
	"TailorShop/api"
	db "TailorShop/db/sqlc"
	"TailorShop/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// DBSource := config.DBDriver + "://" + config.DBUser + ":" + config.DBPassword + "@" + config.DBHost + ":" + config.DBPort + "/" + config.DBName + "?sslmode=disable"
	// DBSource := "postgres://merlin:%40merlin123@152.69.215.50:5432/app"
	// conn, err := sql.Open(config.DBDriver, config.DBSource)
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
