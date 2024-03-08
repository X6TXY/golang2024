package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x6txy/golang2024/database"
	"github.com/x6txy/golang2024/models"
)

func CreateComment(cp *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := cp.BodyParser(comment); err != nil {
		return cp.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&comment)

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
