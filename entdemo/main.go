package main

import (
	"context"
	"log"

	"github.com/foldera/playground/entdemo/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Printf("failed to close ent client: %s", err.Error())
		}
	}()
	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
