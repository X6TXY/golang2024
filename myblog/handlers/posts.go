package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x6txy/golang2024/database"
	"github.com/x6txy/golang2024/models"

)

func CreatePost(cp *fiber.Ctx) error {
	post := new(models.Post)
	if err := cp.BodyParser(post); err != nil {
		return cp.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&post)

	return cp.Status(200).JSON(post)
}

func ListPosts(c *fiber.Ctx) error {
	posts := []models.Post{}

	database.DB.Db.Find(&posts)

	return c.Status(200).JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	postID := c.Params("id")
	var post models.Post

	if err := database.DB.Db.First(&post, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	postID := c.Params("id")
	var post models.Post
	if err := database.DB.Db.First(&post, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	database.DB.Db.Delete(&post)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}

func UpdatePost(c *fiber.Ctx) error {
	postID := c.Params("id")
	var existingPost models.Post

	if err := database.DB.Db.First(&existingPost, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post not found",
		})
	}

	newPost := new(models.Post)
	if err := c.BodyParser(newPost); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update only the specified fields (you can customize this based on your requirements)
	database.DB.Db.Model(&existingPost).Updates(newPost)

	return c.Status(fiber.StatusOK).JSON(existingPost)
}
