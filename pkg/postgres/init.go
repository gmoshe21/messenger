package postgres

import (
	"Messege/config"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	queryCreateTableUser = `CREATE TABLE IF NOT EXISTS user (
		id 				SERIAL PRIMARY KEY,
		uid 			VARCHAR(50),
    	name 			VARCHAR(50),
    	lastname 		VARCHAR(50),
    	number 			VARCHAR(50),
    	mail 			VARCHAR(50),
	  );`
	queryCreateTableCommunication = `CREATE TABLE IF NOT EXISTS communication (
		id 				SERIAL PRIMARY KEY,
		user1 			VARCHAR(50),
    	room 			VARCHAR(50),
	  );`
	queryCreateTableMessege = `CREATE TABLE IF NOT EXISTS messege (
		id 				SERIAL PRIMARY KEY,
		author			VARCHAR(50),
		recipient 		VARCHAR(50),
		data 			VARCHAR(50),
		time 			VARCHAR(200),
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

	return database, nil
}
