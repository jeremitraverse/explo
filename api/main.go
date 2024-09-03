package main

import (
	"fmt"
	"os"

	"github.com/jeremitraverse/explo/db"
	"github.com/jeremitraverse/explo/server"
)

func main() {
	args := os.Args

	if len(args) > 1 && args[1] == "seed" {
		fmt.Println("Started seeding db")
		db.Seed()
	}

	server.Start()
}
