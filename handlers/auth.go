package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"rehber/database"
	"rehber/models"
	"rehber/utils"
)

//Login is a function which login with id.
func Login(ctx *fiber.Ctx) error {
	type request struct {
		ID uint `json:"user_id"`
	}
	var err error
	user := new(models.User)
	req := new(request)

	if err = ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted Body.")
	}

	if err = database.DB.Preload("PhoneNumbers").Preload("Contacts.PhoneNumbers").First(user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "User Not Found.")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Error While Finding.")
	}

	token := utils.GenerateToken(user.ID)

	return utils.SuccessPresenter(ctx, "Logged in.",
		fiber.Map{
			"token": token,
			"user":  user,
		})
}
