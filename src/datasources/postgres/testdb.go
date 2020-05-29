package testdb

import (
	"database/sql"
	"deploy-sample/src/logger"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgresql_username    = "postgresql_username"
	postgresql_password    = "postgresql_password"
	postgresql_host        = "postgresql_host"
	postgresql_port        = "postgresql_port"
	postgresql_cliqtale_db = "postgresql_cliqtale_db"
)

var (
	Client *sql.DB

	username = os.Getenv(postgresql_username)
	password = os.Getenv(postgresql_password)
	host     = os.Getenv(postgresql_host)
	schema   = os.Getenv(postgresql_cliqtale_db)
	port     = os.Getenv(postgresql_port)
)

func init() {
	dbSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, schema)

	var err error

	Client, err = sql.Open("postgres", dbSourceName)
	if err != nil {
		logger.Error("Error while connecting to db", err)
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	logger.Info("Database successfully configured")
}
