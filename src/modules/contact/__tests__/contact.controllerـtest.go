package contact

import (
	"testing"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
)

func TestControllerRoute(t *testing.T) {
	s := NewService()
	ctrl := NewController(s)
	app := fiber.New()
	ctrl.RegisterRoutes(app)
	req := httptest.NewRequest("GET", "/contact", nil)
	resp, _ := app.Test(req)
	if resp.StatusCode != 200 {
		t.Fatalf("expected status 200; got %d", resp.StatusCode)
	}
}
