package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/src/modules/app"

	velora "github.com/aryan-salemababdi/Velora/app"
)

func main() {
	a := velora.New()

	velora.RegisterMiddleware("auth", func(c *fiber.Ctx) error {
		return c.Next()
	})

	velora.RegisterMiddleware("logging", func(c *fiber.Ctx) error {
		fmt.Println("[LOG]", c.Path())
		return c.Next()
	})

	velora.RegisterMiddleware("cors", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	})

	a.UseGlobalMiddleware("logging", "cors", "auth")

	if err := a.RegisterModule("app", app.New()); err != nil {
		log.Fatal(err)
	}

	a.WatchModules(".", 2*time.Second)

	if err := a.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
