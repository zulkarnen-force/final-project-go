package main

import (
	"final-project-go/databases"
	"final-project-go/router"

	"log"
)

func main() {
	databases.StartDB()
	r := router.Router()
	
	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())		
	}
}