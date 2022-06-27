package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string        `json:"user_name" gorm:"not null;size:45" query:"name"`
	UserSurname  string        `json:"user_surname" gorm:"not null;size:45" query:"surname"`
	Company      string        `json:"company" query:"company"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers" query:"phone"`
	Contacts     []User        `json:"contacts" gorm:"many2many:users_contacts"`
}
