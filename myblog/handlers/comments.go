package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/x6txy/golang2024/database"
	"github.com/x6txy/golang2024/models"
)

func CreateComment(cp *fiber.Ctx) error {
	comment := new(models.Comment)
	postIDParam := cp.Params("id")
	if err := cp.BodyParser(comment); err != nil {
		return cp.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing request body"})
	}

	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		return cp.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid post ID format"})
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

	var post models.Post
	if err := database.DB.Db.First(&post, postID).Error; err != nil {
		return cp.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Post not found"})
	}

	comment.UserID = userID
	comment.PostID = uint(postID)

	if err := database.DB.Db.Create(&comment).Error; err != nil {
		return cp.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error saving the comment to the database", "error": err.Error()})
	}

	database.DB.Db.First(&comment.User, comment.UserID)
	database.DB.Db.First(&comment.Post, comment.PostID)

	return cp.Status(200).JSON(comment)
}

func ListComments(c *fiber.Ctx) error {
	comment := []models.Comment{}

	database.DB.Db.Find(&comment)

	return c.Status(200).JSON(comment)
}

func GetComment(c *fiber.Ctx) error {
	commentID := c.Params("id")
	var comment models.Comment

	if err := database.DB.Db.First(&comment, commentID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Comment not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(comment)
}

func DeleteComment(c *fiber.Ctx) error {
	commentID := c.Params("id")
	var comment models.Comment
	if err := database.DB.Db.First(&comment, commentID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Comment not found",
		})
	}

	database.DB.Db.Delete(&comment)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Comment deleted successfully",
	})
}

func UpdateComment(c *fiber.Ctx) error {
	commentID := c.Params("id")
	var existingComment models.Comment

	if err := database.DB.Db.First(&existingComment, commentID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Comment not found",
		})
	}

	newComment := new(models.Comment)
	if err := c.BodyParser(newComment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update only the specified fields (you can customize this based on your requirements)
	database.DB.Db.Model(&existingComment).Updates(newComment)

	return c.Status(fiber.StatusOK).JSON(existingComment)
}
