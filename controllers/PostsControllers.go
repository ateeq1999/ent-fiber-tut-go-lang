package controllers

import (
	"ent-demo/ent/post"
	"ent-demo/ent/user"
	"ent-demo/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Post struct {
    Title string `json:"title" xml:"title" form:"title"`
    UserId int `json:"userId" xml:"userId" form:"userId"`
}

func GetPosts(c *fiber.Ctx) error {
	posts, entError := utils.MyClient.Post.
		Query().
		All(utils.MyCtx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	if entError == nil {
		log.Println("Post Created : ", posts)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"posts": posts,
	})
}

func CreatePost(c *fiber.Ctx) error {
	postData := new(Post)

	if er := c.BodyParser(postData); er != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "error":   er,
        })
	}

	userData, userEntError := utils.MyClient.User.
		Query().
		Where(user.IDEQ(postData.UserId)).
		Only(utils.MyCtx)

	if userEntError != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"userEntError":   userEntError,
		})
	}

	log.Println("MyClient : " , utils.MyClient)

	createdPost, err := utils.MyClient.Post.
		Create().
		SetTitle(postData.Title).
		SetUserId(postData.UserId).
		AddUser(userData).
		Save(utils.MyCtx)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "err":   err,
        })
	}

	if err == nil {
		log.Println("Post Created : ", createdPost)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Post": createdPost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	postData := new(Post)

	postId, paramsErr := strconv.Atoi(c.Params("id"))

	if paramsErr != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "paramsErr":   paramsErr,
        })
	}

	if BodyParserError := c.BodyParser(postData); BodyParserError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "BodyParserError":   BodyParserError,
        })
	}

	updatedPost, entError := utils.MyClient.Post.
		UpdateOneID(postId).
		SetTitle(postData.Title).
		SetUserId(postData.UserId).
		Save(utils.MyCtx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	if entError == nil {
		log.Println("Post Created : ", updatedPost)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Post": updatedPost,
	})
}

func GetPost(c *fiber.Ctx) error {
	postId, paramsErr := strconv.Atoi(c.Params("id"))

	if paramsErr != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "paramsErr":   paramsErr,
        })
	}

	getPost, entError := utils.MyClient.Post.
		Query().
		Where(post.IDEQ(postId)).
		Only(utils.MyCtx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	if entError == nil {
		log.Println("Post Created : ", getPost)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Post": getPost,
	})
}

func DeletePost(c *fiber.Ctx) error {
	postId, paramsErr := strconv.Atoi(c.Params("id"))

	if paramsErr != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "paramsErr":   paramsErr,
        })
	}

	entError := utils.MyClient.Post.
		DeleteOneID(postId).
		Exec(utils.MyCtx)

	if entError != nil {
		return c.Status(404).JSON(fiber.Map{
            "success": false,
            "entError":   entError,
        })
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"msg": "Post with id : " + c.Params("id"),
	})
}
