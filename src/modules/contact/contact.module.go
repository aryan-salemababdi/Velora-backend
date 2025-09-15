package contact

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type ContactModule struct{}

func New() *ContactModule { return &ContactModule{} }

func (m *ContactModule) Register(container *dig.Container, app *fiber.App) error {
	if err := container.Provide(NewContactService); err != nil {
		return err
	}
	if err := container.Provide(NewController); err != nil {
		return err
	}
	return container.Invoke(func(ctrl *Controller) { ctrl.RegisterRoutes(app) })
}

func (m *ContactModule) Middlewares() []string {
	return []string{}
}
