package main

import (
	"tickets/internal/config"
	"tickets/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := config.ConnectDatabase(cfg)

	server.StartServer( db, cfg)
}


// git add .
// git commit -m "added update event handler and register route"
// git push origin main