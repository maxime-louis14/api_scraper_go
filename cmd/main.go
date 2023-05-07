// Package main provides ...
//
//     Schemes: http
//     Host: localhost:3000
//     BasePath: /api/v1
//     Version: 1.0.0
//
// swagger:meta
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awsome API")
}

func setupRoutes(app *fiber.App) {

	app.Get("/api", welcome)
	//User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/user/:id", routes.DeleteUser)
	// Recette endpoints
	app.Post("/api/recettes", routes.PostRecette)
	app.Get("/api/recettes", routes.GetRecettesDetails)
	app.Get("/api/recettes/:id", routes.GetRecette)
	app.Get("/api/recettes/name/:name", routes.GetRecetteByName)
	app.Get("/api/recettes/ingredient/:ingredient", routes.GetRecettesIngredient)
	

}

func main() {


	
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func SawgInit() {
	panic("unimplemented")
}
