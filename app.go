package main

import (
	"github.com/ArifulProtik/arifulprotik-api/config"
	"github.com/ArifulProtik/arifulprotik-api/db"
	"github.com/ArifulProtik/arifulprotik-api/server"
)

func main() {
	config.LoadConfig("./", "app", "env")
	db.DbClient()
	server.RunServer()

}
