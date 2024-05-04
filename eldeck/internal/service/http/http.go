package http

import (
	"github.com/gofiber/fiber/v2"
)

type eldeckHandler struct {
	eldeckService interface{}
}

func NeweldeckHandlers(app *fiber.App, eldeckService *interface{}) *eldeckHandler {
	return nil
}
