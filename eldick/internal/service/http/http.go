package http

import (
	"github.com/gofiber/fiber/v2"
)

type ElDickHandler struct {
	elDickService interface{}
}

func NewElDickHandlers(app *fiber.App, elDickService *interface{}) *ElDickHandler {
	return nil
}
