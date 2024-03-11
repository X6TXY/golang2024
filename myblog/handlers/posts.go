package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/x6txy/golang2024/database"
	"github.com/x6txy/golang2024/models"
)

func CreatePost(cp *fiber.Ctx) error {
    post := new(models.Post)
    if err := cp.BodyParser(post); err != nil {
        return cp.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing request body"})
    }

    userIDInterface := cp.Locals("userID")
    fmt.Printf("Extracted userIDInterface: %#v\n", userIDInterface)

    if userIDInterface == nil {
        fmt.Println("userIDInterface is nil. User ID was not set in context.")
        return cp.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User not authenticated"})
    }

    userID, ok := userIDInterface.(uint)
    if !ok || userID == 0 {
        fmt.Printf("Failed to assert userIDInterface to uint or userID is 0. userIDInterface actual value: %#v\n", userIDInterface)
        return cp.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID format or userID is 0"})
    }

    var user models.User
    if err := database.DB.Db.First(&user, userID).Error; err != nil {
        return cp.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User does not exist", "userID": userID})
    }

    post.UserID = userID
    if err := database.DB.Db.Create(&post).Error; err != nil {
        return cp.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error saving the post to the database", "error": err.Error()})
    }

	database.DB.Db.First(&post.User, post.UserID)

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
