package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type ErrorMessage[Message string | *ErrorResponse | []*ErrorResponse] struct {
	Error Message `json:"error"`
}

func SendBasicError(c *fiber.Ctx, err error, status int) error {
	return c.Status(status).JSON(ErrorMessage[string]{fmt.Sprint(err)})
}

func SendPostError[E *ErrorResponse | []*ErrorResponse](c *fiber.Ctx, err E, status int) error {
	return c.Status(status).JSON(ErrorMessage[E]{err})
}

func ValidatePostBody[PB any](postBody PB) (errors []*ErrorResponse) {
	validate := validator.New()
	err := validate.Struct(postBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
