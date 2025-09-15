package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	middleware "backend/src/common/middlewares"
	"backend/src/modules/admin"
	"backend/src/modules/app"
	"backend/src/modules/contact"

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

	velora.RegisterMiddleware("adminGuard", middleware.AdminGuard)

	if err := a.RegisterModule("app", app.New()); err != nil {
		log.Fatal(err)
	}

	//registering another module
	if err := a.RegisterModule("admin", admin.New()); err != nil {
		log.Fatal(err)
	}

	if err := a.RegisterModule("contact", contact.New()); err != nil {
		log.Fatal(err)
	}

	a.WatchModules(".", 2*time.Second)

	if err := a.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
