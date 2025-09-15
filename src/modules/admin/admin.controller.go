package admin

import (
	middleware "backend/src/common/middlewares"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Service *Service
}

func NewController(s *Service) *Controller { return &Controller{Service: s} }

func (c *Controller) RegisterRoutes(app *fiber.App) {
	adminGroup := app.Group("/admin", middleware.AuthGuard, middleware.AdminGuard)

	adminGroup.Get("/dashboard", func(ctx *fiber.Ctx) error {
		data := c.Service.FindAll()
		return ctx.JSON(fiber.Map{
			"data": data,
		})
	})
}
