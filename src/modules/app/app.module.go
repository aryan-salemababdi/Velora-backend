package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type AppModule struct{}

func New() *AppModule { return &AppModule{} }

func (m *AppModule) Register(container *dig.Container, app *fiber.App) error {
	if err := container.Provide(NewService); err != nil { return err }
	if err := container.Provide(NewController); err != nil { return err }
	return container.Invoke(func(ctrl *Controller) { ctrl.RegisterRoutes(app) })
}

func (m *AppModule) Middlewares() []string {
	return []string{}
}
