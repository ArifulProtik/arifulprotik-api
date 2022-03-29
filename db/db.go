package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ArifulProtik/arifulprotik-api/config"
	"github.com/ArifulProtik/arifulprotik-api/ent"
	"github.com/ArifulProtik/arifulprotik-api/ent/migrate"
	_ "github.com/lib/pq"
)

var (
	Ent *ent.Client
)

func DbClient() *ent.Client {
	// "host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.C.DB_HOST, config.C.DB_PORT, config.C.DB_USER, config.C.DB_NAME, config.C.DB_PASS)
	client, err := ent.Open("postgres", psqlInfo)
	if err != nil {
		log.Panicln(err)
	}
	// defer client.Close()
	log.Println("DB Connected")
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); !errors.Is(err, nil) {
		log.Println(err)
	}
	Ent = client
	return client
}
