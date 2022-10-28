package controllers

import (
	"rapidtech/shoppingcart/database"
	"rapidtech/shoppingcart/models"
	"time"

	"github.com/gofiber/fiber/v2"
	/*"github.com/gofiber/fiber/v2/middleware/session"
	  session tidak perlu digunakan karena menggunakan
	  API
	*/
	"github.com/golang-jwt/jwt/v4" //JSON Web Token
	"golang.org/x/crypto/bcrypt"   //untuk password

	"gorm.io/gorm" //untuk database
)

type LoginForm struct {
	Username string `form:"name" json:"name" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	//store *session.Store (API tidak menggunakan session)
	Db *gorm.DB
}

func InitAuthController() *AuthController {

	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.User{})
	return &AuthController{Db: db}
}

// get /login

/*(Ini tidak dibutuhkan untuk API, karena rendernya HTML, sedangkan API tidak membutuhkan HTML)
 func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
*/

// POST /register
func (controller *AuthController) Register(c *fiber.Ctx) error {
	// load all user
	var myform models.User
	var convertpass LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(500)
	}
	convertpassword, _ := bcrypt.GenerateFromPassword([]byte(convertpass.Password), 10)
	sHash := string(convertpassword)

	myform.Password = sHash

	// save user
	err := models.CreateUser(controller.Db, &myform)
	if err != nil {
		return c.SendStatus(500)
	}
	// if succeed
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": myform})
}

// POST /login
func (controller *AuthController) LoginPosted(c *fiber.Ctx) error {
	// load all products

	/* (API tidak membutuhkan session)
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	*/

	var user models.User
	var myform LoginForm
	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}

	er := models.FindByUsername(controller.Db, &user, myform.Username)
	if er != nil {
		return c.Redirect("/login") // http 500 internal server error
	}

	// hardcode auth

	/* mycompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	if mycompare != nil {
		sess.Set("username", user.Name)
		sess.Set("userID", user.Id)
		sess.Save()

		return c.Redirect("/profile")
	}
	return c.Redirect("/login")
	*/

	// Untuk login API di Postman
	mycompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	if mycompare != nil {
		exp := time.Now().Add(time.Hour * 72)
		claims := jwt.MapClaims{
			"id":    user.Id,
			"name":  user.Name,
			"admin": true,
			"exp":   exp.Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("mysecretpassword"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"token":   t,
			"expired": exp.Format("2006-01-02 15:04:05"),
		})
	}
	return c.SendStatus(fiber.StatusUnauthorized)

}

// /logout
/* func (controller *AuthController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
*/
