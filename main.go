package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"

	APIShoppingCart "rapidtech/shoppingcart/APIShoppingCart"
	"rapidtech/shoppingcart/controllers"
)

func main() {
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public", "./public")

	// controllers
	//helloController := controllers.InitHelloController(store)
	prodController := controllers.InitProductController()
	authController := controllers.InitAuthController(store)
	userController := controllers.InitUserController()
	//cartController := controllers.InitCartController()
	apiprodController := APIShoppingCart.InitProductController()
	apiuserController := APIShoppingCart.InitUserController()
	apiauthController := APIShoppingCart.InitAuthController()

	/*=======================================================
							WEB
	  =======================================================
	*/
	prod := app.Group("/products")
	prod.Get("/", prodController.IndexProduct)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/productdetail", prodController.GetDetailProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct2)
	prod.Get("/editproduct/:id", prodController.EditProduct)
	prod.Post("/editproduct/:id", prodController.EditPostedProduct)
	prod.Get("/deleteproduct/:id", prodController.DeleteProduct)

	//cart.Get("/deletecart/:id", cartController.DeleteCart)

	user := app.Group("/users")
	user.Get("/", userController.IndexUser)
	user.Get("/create", userController.AddUser)
	user.Post("/create", userController.AddPostedUser)

	///
	user.Get("/userdetail", userController.GetDetailUser)
	user.Get("/detail/:id", userController.GetDetailUser2)
	///
	/*
		user.Get("/editproduct/:id", userController.EditProduct)
		user.Post("/editproduct/:id", userController.EditPostedProduct)
		user.Get("/deleteproduct/:id", userController.DeleteProduct)
	*/

	app.Get("/login", authController.Login)
	app.Post("/login", authController.LoginPosted)
	app.Get("/logout", authController.Logout)
	/////
	app.Get("/register", userController.AddUser)
	app.Post("/register", userController.AddPostedUser)
	/////
	//app.Get("/profile",authController.Profile)

	// app.Use("/profile", func(c *fiber.Ctx) error {
	// 	sess,_ := store.Get(c)
	// 	val := sess.Get("username")
	// 	if val != nil {
	// 		return c.Next()
	// 	}

	// 	return c.Redirect("/login")

	// })
	app.Get("/profile", func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		val := sess.Get("username")
		if val != nil {
			return c.Next()
		}

		return c.Redirect("/login")

	}, authController.Profile)

	/*
		==================================================
								API
		===================================================
	*/
	apiprod := app.Group("/api/products")
	apiprod.Get("/", apiprodController.IndexProduct)
	apiprod.Get("/create", apiprodController.AddProduct)
	apiprod.Post("/create", apiprodController.AddPostedProduct)
	//apiprod.Get("/productdetail", apiprodController.GetDetailProduct)
	apiprod.Get("/detail/:id", apiprodController.GetDetailProduct2)
	apiprod.Get("/editproduct/:id", apiprodController.EditProduct)
	apiprod.Put("/editproduct/:id", apiprodController.EditPostedProduct)
	apiprod.Delete("/deleteproduct/:id", apiprodController.DeleteProduct)

	apiuser := app.Group("/api/users")
	apiuser.Get("/", apiuserController.IndexUser)
	apiuser.Get("/create", apiuserController.AddUser)
	apiuser.Post("/create", apiuserController.AddPostedUser)
	apiuser.Get("/detail/:id", apiuserController.GetDetailUser2)
	apiuser.Delete("/deleteuser/:id", apiuserController.DeleteUser)

	//app.Get("/login", apiauthController.Login)
	app.Post("/api/login", apiauthController.LoginPosted)
	//app.Get("/logout", apiauthController.Logout)
	app.Post("/api/register", apiauthController.Register)

	app.Listen(":8080")
}
