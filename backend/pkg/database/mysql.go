package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    os.Getenv("DB_PROTOCOL"),
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName: os.Getenv("DB_NAME"),
	}

	log.Print("Connecting to database ...")
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Printf("Error connecting to the database: %s\n", err.Error())
		return nil, err
	}

	log.Println("Pinging to database")
	if err := db.Ping(); err != nil {
		log.Printf("Error could not ping database: %s\n", err.Error())
		return nil, err
	}
	log.Println("Database connected")

	return db, nil
}
