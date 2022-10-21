package UserControllers

import (
	"context"
	"ent-demo/ent/user"
	"ent-demo/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
    Name string `json:"name" xml:"name" form:"name"`
    Password string `json:"password" xml:"password" form:"password"`
    Email string `json:"email" xml:"email" form:"email"`
    Age int `json:"age" xml:"age" form:"age"`
}

var(
	ctx = context.Background()
)

func GetUsers(c *fiber.Ctx) error {
	users, entError := utils.MyClient.User.
		Query().
		All(ctx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	if entError == nil {
		log.Println("user Created : ", users)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	u := new(User)

	if er := c.BodyParser(u); er != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "error":   er,
        })
	}

	log.Println("MyClient : " , utils.MyClient)

	createdUser, err := utils.MyClient.User.
		Create().
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetAge(u.Age).
		Save(ctx)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "err":   err,
        })
	}

	if err == nil {
		log.Println("user Created : ", createdUser)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user": createdUser,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	u := new(User)

	userId, paramsErr := strconv.Atoi(c.Params("id"))

	if paramsErr != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "paramsErr":   paramsErr,
        })
	}

	if BodyParserError := c.BodyParser(u); BodyParserError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "BodyParserError":   BodyParserError,
        })
	}

	updatedUser, entError := utils.MyClient.User.
		UpdateOneID(userId).
		SetName(u.Name).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetAge(u.Age).
		Save(ctx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	if entError == nil {
		log.Println("user Created : ", updatedUser)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user": updatedUser,
	})
}

func GetUser(c *fiber.Ctx) error {
	userId, paramsErr := strconv.Atoi(c.Params("id"))

	if paramsErr != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "paramsErr":   paramsErr,
        })
	}

	getUser, entError := utils.MyClient.User.
		Query().
		Where(user.IDEQ(userId)).
		Only(ctx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	if entError == nil {
		log.Println("user Created : ", getUser)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user": getUser,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userId, paramsErr := strconv.Atoi(c.Params("id"))

	if paramsErr != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "paramsErr":   paramsErr,
        })
	}

	entError := utils.MyClient.User.
		DeleteOneID(userId).
		Exec(ctx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg": "user with id : " + c.Params("id"),
	})
}

// func UpdateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 		UpdateOneID(1).
// 		SetName("ateeg").
// 		SetAge(29).
// 		Save(ctx)

// 	if err != nil {
// 		return nil, fmt.Errorf("canot update user : %w", err)
// 	}

// 	log.Panicln("user uodated : ", u)

// 	return u, nil
// }

// func GetUser(ctx context.Context, client *ent.Client, c *fiber.Ctx) (*ent.User, error) {
// 	u, err := client.User.
// 		Query().
// 		Where(user.NameEQ("Ateeg")).
// 		Only(ctx)

// 	if err != nil {
// 		return nil, fmt.Errorf("canot query user : %w", err)
// 	}

// 	log.Panicln("user query : ", u)

// 	return u, nil
// }

// func DeleteUser(ctx context.Context, client *ent.Client) (string, error) {
// 	err := client.User.
// 		DeleteOneID(2).
// 		Exec(ctx)

// 	if err != nil {
// 		return "error", fmt.Errorf("canot query user : %w", err)
// 	}

// 	return "done", nil
// }