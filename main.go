package main

import (
	controllers "ent-demo/controllers"
	userscontrollers "ent-demo/controllers"
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

    PostApi := app.Group("/posts")

    UserApi.Get("/", userscontrollers.GetUsers)
    UserApi.Post("/", userscontrollers.CreateUser)
    UserApi.Get("/:id", userscontrollers.GetUser)
    UserApi.Put("/:id", userscontrollers.UpdateUser)
    UserApi.Delete("/:id", userscontrollers.DeleteUser)

    PostApi.Get("/", controllers.GetPosts)
    PostApi.Post("/", controllers.CreatePost)
    PostApi.Get("/:id", controllers.GetPost)
    PostApi.Put("/:id", controllers.UpdatePost)
    PostApi.Delete("/:id", controllers.DeletePost)

	app.Listen(":3000")

	utils.MyClient.Close()
}