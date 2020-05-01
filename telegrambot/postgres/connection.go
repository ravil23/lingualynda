package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
)

type Connection struct {
	*pg.DB
}

func NewConnection() *Connection {
	options := &pg.Options{
		Addr:     os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DATABASE"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
	log.Printf("PostgreSQL connection options: address=%s database=%s user=%s", options.Addr, options.Database, options.User)
	return &Connection{
		DB: pg.Connect(options),
	}
}
