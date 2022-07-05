package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type JSONTyperesponse struct {
	Message string
}

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        //return c.SendString("Hello, World!")
		return c.JSON(JSONTyperesponse{Message: "Hello, Chutiya!"})
    })

    log.Fatal(app.Listen(":8080"))
}