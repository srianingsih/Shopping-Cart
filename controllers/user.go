package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"rapidtech/shoppingcart/database"
	"rapidtech/shoppingcart/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type UserController struct {
	// declare variables
	Db *gorm.DB
}

func InitUserController() *UserController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.User{})

	return &UserController{Db: db}
}

// routing
// GET /users
func (controller *UserController) IndexUser(c *fiber.Ctx) error {
	// load all products
	var users []models.User
	err := models.ReadUser(controller.Db, &users)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("users", fiber.Map{
		"Title": "Users",
		"Users": users,
	})
}

// GET /register
func (controller *UserController) AddUser(c *fiber.Ctx) error {
	return c.Render("adduser", fiber.Map{
		"Title": "Register",
	})
}

// POST /user/create
func (controller *UserController) AddPostedUser(c *fiber.Ctx) error {
	// load all products
	var myform models.User
	//var convertpass LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}
	//convertpassword, _ := bcrypt.GenerateFromPassword([]byte(convertpass.Password), 10)
	//sHash := string(convertpassword)

	//myform.Password = sHash

	// save product
	err := models.CreateUser(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/login")
	}
	// if succeed
	return c.Redirect("/login")
}

// GET /user/productdetail?id=xxx
func (controller *UserController) GetDetailUser(c *fiber.Ctx) error {
	id := c.Query("id")
	idn, _ := strconv.Atoi(id)

	var user models.User
	err := models.ReadUserById(controller.Db, &user, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("userdetail", fiber.Map{
		"Title": "User Detail",
		"User":  user,
	})
}

// GET /user/detail/xxx
func (controller *UserController) GetDetailUser2(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var user models.User
	err := models.ReadUserById(controller.Db, &user, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("userdetail", fiber.Map{
		"Title": "Detail User",
		"User":  user,
	})
}

/*
// / GET products/editproduct/xx
func (controller *ProductController) EditProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editproduct", fiber.Map{
		"Title":   "Edit Produk",
		"Product": product,
	})
}


// / POST products/editproduct/xx
func (controller *ProductController) EditPostedProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Product

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/products")
	}
	product.Name = myform.Name
	product.Quantity = myform.Quantity
	product.Price = myform.Price
	// save product
	models.UpdateProduct(controller.Db, &product)

	return c.Redirect("/products")

}
*/
// / GET /products/deleteproduct/xx
func (controller *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var user models.User
	models.DeleteUserById(controller.Db, &user, idn)
	return c.Redirect("/users")
}
