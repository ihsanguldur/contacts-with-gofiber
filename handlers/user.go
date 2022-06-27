package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm/clause"
	"rehber/database"
	"rehber/models"
	"rehber/repository"
	"rehber/utils"
	"strconv"
	"strings"
)

//CreateUser is a function which create one user with required values.
func CreateUser(ctx *fiber.Ctx) error {
	var err error
	user := new(models.User)

	if err = ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Corrupted body.")
	}

	if user.UserName == "" || user.UserSurname == "" || len(user.PhoneNumbers) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Check Your Inputs.")
	}

	if ok, i := utils.IsPhoneValid(user.PhoneNumbers); !ok {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("%d.th Phone Number is Not Valid.", i+1))
	}

	if err = database.DB.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			return fiber.NewError(fiber.StatusBadRequest, "Phone Number is Already in Use.")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Error While Creating.")
	}

	return utils.SuccessPresenter(ctx, "User Created.", user)
}

//DeleteUser is a function which delete one user with id.
//
//It is not hard delete. It will update deletedAt field.
func DeleteUser(ctx *fiber.Ctx) error {
	var err error
	user := new(models.User)
	phone := new(models.PhoneNumber)
	id := ctx.Params("id")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsTokenValid(id, token) {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized Token.")
	}

	if err = repository.DeleteUser(user, phone, id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error While Deleting.")
	}

	return utils.SuccessPresenter(ctx, "User Deleted.", user)
}

//UpdateUser is a function which update user with id.
func UpdateUser(ctx *fiber.Ctx) error {
	var err error
	user := new(models.User)
	id := ctx.Params("id")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsTokenValid(id, token) {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized Token.")
	}

	if err = ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted Body.")
	}

	if user.UserName == "" && user.UserSurname == "" && len(user.PhoneNumbers) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Check Your Inputs.")
	}

	if ok, i := utils.IsPhoneValid(user.PhoneNumbers); !ok {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("%d.th Phone Number is Not Valid.", i+1))
	}

	uid, _ := strconv.Atoi(id)
	user.ID = uint(uid)
	if err = repository.UpdateUser(user); err != nil {
		if strings.Contains(err.Error(), "23505") {
			return fiber.NewError(fiber.StatusBadRequest, "Phone Number is Already in Use.")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Error While Updating.")
	}

	return utils.SuccessPresenter(ctx, "User Updated.", user)
}

//SearchUser is a function which searches users by name, surname, company and phone numbers.
func SearchUser(ctx *fiber.Ctx) error {
	var err error
	users := &[]models.User{}
	searched := ctx.Query("q")
	id := ctx.Params("id")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsTokenValid(id, token) {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized Token.")
	}

	whereStr := `
			((user_name ILIKE ? OR
			user_surname ILIKE ? OR
			users.user_name||' '||users.user_surname ILIKE ? OR
			users.user_surname||' '||users.user_name ILIKE ?) OR
			company ILIKE ?) OR
			(phone_numbers.user_id = users.id AND (phone_numbers.number LIKE ?))`

	if err = database.DB.
		Preload("PhoneNumbers").
		Distinct("users.id", "users.created_at", "users.updated_at", "users.deleted_at", "user_name", "user_surname", "company").
		Table("users, phone_numbers").
		Where(whereStr,
			"%"+searched+"%",
			"%"+searched+"%",
			"%"+searched+"%",
			"%"+searched+"%",
			"%"+searched+"%",
			"%"+searched+"%").
		Clauses(clause.Returning{}).
		Find(users).
		Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error While Searching.")
	}

	return utils.SuccessPresenter(ctx, "User Found.", users)
}

//ListContacts is a function which get users contacts with id.
func ListContacts(ctx *fiber.Ctx) error {
	var err error
	user := new(models.User)
	id := ctx.Params("id")
	token := ctx.Locals("user").(*jwt.Token)

	if !utils.IsTokenValid(id, token) {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized Token.")
	}

	if err = database.DB.Preload("Contacts.PhoneNumbers").First(user, "id = ?", id).Error; err != nil { //nested preload.
		return fiber.NewError(fiber.StatusInternalServerError, "Error While Finding.")
	}

	return utils.SuccessPresenter(ctx, "User's Contacts Listed.", user.Contacts)
}

/*
SELECT DISTINCT "users"."id","users"."created_at","users"."updated_at","users"."deleted_at","users"."user_name","users"."user_surname","users"."company"
FROM users, phone_numbers
WHERE
(
(user_name ILIKE '%%' OR user_surname ILIKE '%%' OR "users"."user_name"||' '||"users"."user_surname" ILIKE '%%'
 OR "users"."user_surname"||' '||"users"."user_name" ILIKE '%%')
 OR company ILIKE '%%')
OR (phone_numbers.user_id = users.id AND (phone_numbers.number ILIKE '%%')) AND "users"."deleted_at"  IS NULL*/
