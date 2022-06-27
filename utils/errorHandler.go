package utils

import "github.com/gofiber/fiber/v2"

//ErrorHandler is a custom Error Handler for gofiber.
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	errCode := fiber.StatusInternalServerError
	errMsg := "Something Went Wrong."

	if e, ok := err.(*fiber.Error); ok {
		errCode = e.Code
		errMsg = e.Message
	}
	return ErrorPresenter(ctx, errMsg, errCode)
}
