package main

import "github.com/gofiber/fiber/v2"

func Client() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home.html", nil)
	})

	app.Listen(":3000")

}
