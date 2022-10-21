package main

import (
	UserControllers "ent-demo/controllers"
	"ent-demo/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    utils.ConnectToDB()
	
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
            "success": false,
            "msg":   "Hello World",
        })
	})

    UserApi := app.Group("/users")

    UserApi.Get("/", UserControllers.GetUsers)
    UserApi.Post("/", UserControllers.CreateUser)
    UserApi.Get("/:id", UserControllers.GetUser)
    UserApi.Put("/:id", UserControllers.UpdateUser)
    UserApi.Delete("/:id", UserControllers.DeleteUser)

	app.Listen(":3000")

	utils.MyClient.Close()
}