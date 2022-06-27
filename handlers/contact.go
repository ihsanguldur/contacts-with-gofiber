package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"rehber/models"
	"rehber/repository"
	"rehber/utils"
	"strconv"
)

//CreateContact is a function which create contact for user.
func CreateContact(ctx *fiber.Ctx) error {
	type request struct {
		PhoneNumber string `json:"phone_number"`
	}

	var err error
	user := new(models.User)
	phone := new(models.PhoneNumber)
	req := new(request)
	id := ctx.Params("id")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsTokenValid(id, token) {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized Token.")
	}

	if err = ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted Body.")
	}

	uid, _ := strconv.Atoi(id)
	if err = repository.CreateContact(req.PhoneNumber, user, phone, uint(uid)); err != nil {
		if err.Error() == "Are You so Lonely?" {
			return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessPresenter(ctx, "Contact Created.", user)
}
