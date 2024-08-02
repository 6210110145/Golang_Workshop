package routes

import (
	// exercise 5.3
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	// MIDDLEWARE
	// exercise 5.0
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	app.Static("/", "./public") // images, CSS, and JavaScript.

	api := app.Group("/api")

	// api/v1
	v1 := api.Group("/v1")
	v1.Get("/", c.Hello)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.ParamsTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidTest)

	// exercise 5.1
	v1.Post("/fact/:num", c.Factorial)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Post("/", c.AddDog)
	dog.Get("/", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogJson)
	// exercise 7.0.2
	dog.Get("/ddog", c.GetDeleteDogs)

	// exercise 7.1
	dog.Get("/sdog", c.GetSomeDogs)

	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	// api/v2
	v2 := api.Group("/v2")
	v2.Get("/", c.HelloV2)

	// api/v3
	// exercise 5.2
	v3 := api.Group("/v3")
	v3.Post("/james", c.Acii)

	// exercise 6
	v1.Post("/register", c.ValidateRegister)

	// exercise 7.0.2
	com := v1.Group("/company")
	com.Post("/", c.AddCompany)
	com.Get("/", c.GetAllCompany)
	com.Get("/filter", c.GetNameCompany)
	com.Put("/:id", c.UpdateCompany)
	com.Delete("/:id", c.DeleteCompany)
}
