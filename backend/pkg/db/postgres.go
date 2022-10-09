package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	User   string
	Passwd string
	Net    string
	Host   string
	Port   string
	DBName string
}

func Init() (*pgxpool.Pool, error) {
	cfg := Config{
		User:   os.Getenv("POSTGRES_USER"),
		Passwd: os.Getenv("POSTGRES_PASSWORD"),
		Host:   os.Getenv("POSTGRES_HOST"),
		Port:   os.Getenv("POSTGRES_PORT"),
		DBName: os.Getenv("POSTGRES_DB"),
	}
	connString := formatDSN(cfg)

	log.Print("Connecting to database ...")
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	log.Println("Database connected")

	return dbpool, nil
}

func formatDSN(cfg Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Passwd, cfg.Host, cfg.Port, cfg.DBName)
}
