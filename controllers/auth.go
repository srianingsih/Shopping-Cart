package controllers

import (
	"rapidtech/shoppingcart/database"
	"rapidtech/shoppingcart/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type LoginForm struct {
	Username string `form:"name" json:"name" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	store *session.Store
	Db    *gorm.DB
}

func InitAuthController(s *session.Store) *AuthController {

	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.User{})
	return &AuthController{Db: db, store: s}
}

// get /login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// POST /login
func (controller *AuthController) LoginPosted(c *fiber.Ctx) error {
	// load all products

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}

	var user models.User
	var myform LoginForm
	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}

	er := models.FindByUsername(controller.Db, &user, myform.Username)
	if er != nil {
		return c.Redirect("/register") // http 500 internal server error
	}

	// hardcode auth
	mycompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	if mycompare != nil {
		sess.Set("username", user.Name)
		sess.Set("userID", user.Id)
		sess.Save()

		return c.Redirect("/products")
	}
	return c.Redirect("/register")

}

// /profile
func (controller *AuthController) Profile(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	val := sess.Get("username")

	return c.Render("myview", fiber.Map{
		"Title":    "Profile",
		"username": val,
	})
}

// /logout
func (controller *AuthController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
