package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `form:"id" json:"id" validate:"required"`
	Name     string `form:"name" json:"name" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

// CRUD
func CreateUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadUser(db *gorm.DB, users *[]User) (err error) {
	err = db.Find(users).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadUserById(db *gorm.DB, user *User, id int) (err error) {
	err = db.Where("id=?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateUser(db *gorm.DB, user *User) (err error) {
	db.Save(user)

	return nil
}
func DeleteUserById(db *gorm.DB, user *User, id int) (err error) {
	db.Where("id=?", id).Delete(user)

	return nil
}

func FindByUsername(db *gorm.DB, user *User, username string) (err error) {
	err = db.Where("name=?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
