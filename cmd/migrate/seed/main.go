package main

import (
	"log"

	"github.com/studysoros/go-social-media/internal/db"
	"github.com/studysoros/go-social-media/internal/store"
)

func main() {
	addr := "postgres://admin:admin@localhost/social?sslmode=disable"
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
