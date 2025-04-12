package main

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"log"

	"golinks/internal/db"
	"golinks/internal/rest"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	defer db.ClosePool()

	app := fiber.New()
	app.Use(swagger.New(swagger.Config{FilePath: "./docs/swagger/swagger.json"}))

	rest.RegisterLinksHandlers(app)

	log.Fatalln(app.Listen("localhost:8080"))
}
