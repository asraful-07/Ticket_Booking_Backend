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
// git commit -m "Implement user registration with validation and error handling"
// git push origin main