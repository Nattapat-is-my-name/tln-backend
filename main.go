package main

import (
	"fmt"
	"log"
	"tln-backend/App"
	_ "tln-backend/docs"
)

// @title Talardnad API
// @version 1.0
// @description API user management Server by Fiber | Doc by Swagger.
// @contact.name admin
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host api.talardnad.com
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http https
func main() {
	config, err := App.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := App.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	allHandlers, userRepo, providerRepo, err := App.InitializeHandlers(db)
	if err != nil {
		log.Fatal(err)
	}

	server := App.InitializeServer(userRepo, providerRepo)
	server.MapHandlers(allHandlers)

	address := fmt.Sprintf("%s:%s", config.App.Host, config.App.Port)
	App.StartServer(server, address)
}
