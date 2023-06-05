package main

import (
	"log"
	"os"
)

func main() {
	args := Args{
		conn: "games.sqlite",
		port: ":8080",
	}

	if conn := os.Getenv("DB_CONN"); conn != "" {
		args.conn = conn
	}

	if port := os.Getenv("PORT"); port != "" {
		args.port = ":" + port
	}

	if err := Run(args); err != nil {
		log.Println(err)
	}

}
