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
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  os.Getenv("MYSQL_PROTOCOL"),
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}

	log.Print("Connecting to database ...")
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Printf("Error connecting to the database: %s\n", err.Error())
		return nil, err
	}

	log.Println("Pinging to database ...")
	if err := db.Ping(); err != nil {
		log.Printf("Error could not ping database: %s\n", err.Error())
		return nil, err
	}
	log.Println("Database connected")

	return db, nil
}
