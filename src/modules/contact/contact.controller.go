package contact

import (
	dto "backend/src/modules/contact/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Service *ContactService
}

func NewController(s *ContactService) *Controller {
	return &Controller{Service: s}
}

func (c *Controller) RegisterRoutes(app *fiber.App) {
	app.Post("/contact", func(ctx *fiber.Ctx) error {
		var request dto.CreateContactDTO

		if err := ctx.BodyParser(&request); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		if err := dto.ValidateDTO(request); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := c.Service.SendContactEmail(request); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
		}

		return ctx.JSON(fiber.Map{"message": "Contact message sent successfully"})
	})
}
