package main

import (
	"final-project-go/databases"
	"final-project-go/router"

	"log"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server order.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  zulkarnen1900016072@webmail.uad.ac.id

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used
// @tokenUrl                    http://localhost:8080/users/login

// @host      localhost:8080
// @BasePath  /

func main() {
	r := router.Router()
	databases.Migrate()
	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())		
	}
}