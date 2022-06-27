package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"rehber/database"
	"rehber/models"
)

//CreateContact is a function which
func CreateContact(phoneNumber string, user *models.User, phone *models.PhoneNumber, id uint) error {
	var err error
	tx := database.DB.Begin()

	//for recover a panics. A recover can stop a panic from aborting the program and let it continue with execution instead.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return err
	}

	if err = tx.Find(phone, "number = ?", phoneNumber).Error; err != nil {
		tx.Rollback()
		return err
	}

	if phone.UserID == id {
		return errors.New("Are You so Lonely?")
	}

	user.Contacts = append(user.Contacts, models.User{Model: gorm.Model{ID: phone.UserID}})
	user.ID = id
	if err = tx.Model(user).Clauses(clause.Returning{}).Updates(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Preload("PhoneNumbers").Preload("Contacts.PhoneNumbers").First(user).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
