package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id int `form:"id" json:"id" validate:"required"`
	//Name      string  `form:"name" json:"name" validate:"required"`
	//Image     string  `form:"image" json:"image" validate:"required"`
	//Deskripsi string  `form:"desc" json:"desc" validate:"required"`
	Quantity  int     `form:"quantity" json:"quantity" validate:"required"`
	Price     float32 `form:"price" json:"price" validate:"required"`
	ProductID int     `form:"productid" json:"productid" validate:"required"`
}

// CRUD
func CreateCart(db *gorm.DB, newCart *Cart) (err error) {
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadCarts(db *gorm.DB, carts *[]Cart) (err error) {
	err = db.Find(carts).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadCartById(db *gorm.DB, carts *Cart, id int) (err error) {
	err = db.Where("id=?", id).First(carts).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateCart(db *gorm.DB, cart *Cart) (err error) {
	db.Save(cart)

	return nil
}
func DeleteCartById(db *gorm.DB, cart *Cart, id int) (err error) {
	db.Where("id=?", id).Delete(cart)

	return nil
}
