package database

import (
	"database/sql"
	"fmt"
	"log"
	config "queueflow/configs"

	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) *sql.DB {

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open(
		"postgres",
		connectionString,
	)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Database unreachable:", err)
	}

	log.Println("Database connected")

	return db
}
