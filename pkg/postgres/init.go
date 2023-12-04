package postgres

import (
	"Messege/config"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	queryCreateTableUser = `CREATE TABLE IF NOT EXISTS users (
		uid 			VARCHAR(50),
    	name 			VARCHAR(50),
    	email 			VARCHAR(50),
    	password 		VARCHAR(50)
	  );`
	queryCreateTableCommunication = `CREATE TABLE IF NOT EXISTS communication (
		id 				SERIAL PRIMARY KEY,
		user1 			VARCHAR(50),
    	room 			VARCHAR(50)
	  );`
	queryCreateTableMessege = `CREATE TABLE IF NOT EXISTS messege (
		id 				SERIAL PRIMARY KEY,
		author			VARCHAR(50),
		recipient 		VARCHAR(50),
		data 			VARCHAR(10000),
		time 			timestamp 
	  );`
	queryCreateTableFriendRequest = `CREATE TABLE IF NOT EXISTS friend_request (
		id 				SERIAL PRIMARY KEY,
		user1 			VARCHAR(50),
    	user2 			VARCHAR(50)
	  );`
)

func InitPsqlDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)
	database, err := sqlx.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}
	database.MustExec(queryCreateTableUser)
	database.MustExec(queryCreateTableCommunication)
	database.MustExec(queryCreateTableMessege)
	database.MustExec(queryCreateTableFriendRequest)

	return database, nil
}
