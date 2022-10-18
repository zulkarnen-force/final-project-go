package main

import (
	"final-project-go/databases"
	"final-project-go/router"

	"log"
)

func main() {
	r := router.Router()
	databases.Migrate()
	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())		
	}
}