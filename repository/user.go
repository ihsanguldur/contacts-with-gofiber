package repository

import (
	"gorm.io/gorm/clause"
	"rehber/database"
	"rehber/models"
)

//UpdateUser is a function which start a transaction for update user and its phone numbers.
func UpdateUser(user *models.User) error {
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

	if err = tx.Model(user).Clauses(clause.Returning{}).Updates(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, v := range user.PhoneNumbers {
		if err = tx.Model(&v).Clauses(clause.Returning{}).Updates(&v).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err = tx.Preload("PhoneNumbers").Preload("Contacts.PhoneNumbers").First(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

//DeleteUser is a function which start a transaction for delete user and its phone numbers.
func DeleteUser(user *models.User, phone *models.PhoneNumber, id string) error {
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

	if err = tx.Where("user_id = ?", id).Unscoped().Delete(phone).Error; err != nil {
		return err
	}

	/*if err = tx.Model(user).Association("Contacts").Delete(user, id); err != nil {
		log.Fatal(err)
		return err
	}*/

	tx.Find(user, id)

	if err = tx.Model(user).Association("Contacts").Clear(); err != nil {
		return err
	}

	if err = tx.Clauses(clause.Returning{}).Delete(user, id).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
