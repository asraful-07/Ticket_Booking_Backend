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
// git commit -m "middlewares validation done"
// git push origin main