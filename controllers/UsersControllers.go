package UserControllers

import (
	"context"
	"ent-demo/utils"
	"log"

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

func CreateUser(c *fiber.Ctx) error {
	u := new(User)

	if er := c.BodyParser(u); er != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "error":   er,
        })
	}

	log.Println("MyClient : " , utils.MyClient)

	entUser, err := utils.MyClient.User.
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
		log.Println("user Created : ", entUser)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user": entUser,
	})
}

// func SaveUser(ctx context.Context, client *ent.Client, u *User) (*ent.User, error) {
// 	u, err := client.User.
// 		Create().
// 		SetName(u.Name).
// 		SetEmail(u.Email).
// 		SetPassword(u.Password).
// 		SetAge(u.Age).
// 		Save(ctx)

// 	if err != nil {
// 		return nil, fmt.Errorf("canot update user : %w", err)
// 	}

// 	log.Panicln("user uodated : ", u)

// 	return u, nil
// }

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