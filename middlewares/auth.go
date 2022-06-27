package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"rehber/config"
	"rehber/utils"
)

//Protected is a middleware for JWT.
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("JWT_SECRET")),
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return utils.ErrorPresenter(ctx, "Missing or malformed JWT", fiber.StatusBadRequest)
	}

	return utils.ErrorPresenter(ctx, "Invalid or expired JWT", fiber.StatusUnauthorized)
}
