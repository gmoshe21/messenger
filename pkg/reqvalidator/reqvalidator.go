package reqvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}
	
	return nil
}