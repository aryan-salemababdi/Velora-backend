package admin

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type AdminModule struct{}

func New() *AdminModule { return &AdminModule{} }

func (m *AdminModule) Register(container *dig.Container, app *fiber.App) error {
	if err := container.Provide(NewService); err != nil {
		return err
	}
	if err := container.Provide(NewController); err != nil {
		return err
	}
	return container.Invoke(func(ctrl *Controller) { ctrl.RegisterRoutes(app) })
}

func (m *AdminModule) Middlewares() []string {
	return []string{"adminGuard"}
}
