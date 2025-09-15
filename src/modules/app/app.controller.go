package app

import "github.com/gofiber/fiber/v2"

type Controller struct {
	Service *Service
}

func NewController(s *Service) *Controller {
	return &Controller{Service: s}
}

func (c *Controller) RegisterRoutes(app *fiber.App) {
	app.Get("/ping-db", func(ctx *fiber.Ctx) error {
		if err := c.Service.PingDB(); err != nil {
			return ctx.Status(500).SendString("DB not reachable")
		}
		return ctx.SendString("✅ DB is reachable")
	})
}
