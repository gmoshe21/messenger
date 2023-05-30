package main

import (
	"Messege/config"
	"Messege/internal/server"
	"Messege/pkg/postgres"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
)

func main() {
	cfdFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfdFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	psqlDB, err := postgres.InitPsqlDB(context.Background(), cfg)
	if err != nil {
		log.Fatalf("PostgreSQL init error: %s", err)
	} else {
		log.Println("PostgreSQL connected")
	}

	conn, err := stan.Connect(
		"test-cluster",
		"testcv")
	if err != nil {
		log.Fatalf("Nats-streaming connected error: %s", err)
	} else {
		log.Println("Nats-streaming connected")
	}

	defer func(psqlDB *sqlx.DB) {
		err = psqlDB.Close()
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("PostgreSQL closed properly")
		}
	}(psqlDB)

	s := server.NewServer(cfg, psqlDB, conn)
	err = s.Run(context.Background())
	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}

}
