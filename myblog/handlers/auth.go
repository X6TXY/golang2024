package handlers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/x6txy/golang2024/database"
	"github.com/x6txy/golang2024/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// JWT secret key (replace with a strong, unique key)
var jwtSecret = []byte("your-secret-key")

func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}
	user.Password = string(hashedPassword)

	database.DB.Db.Create(&user)

	user.Password = ""
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}

func Signin(c *fiber.Ctx) error {
	loginData := new(models.User)
	if err := c.BodyParser(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	var user models.User
	if err := database.DB.Db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error querying the database",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error generating JWT token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Db.Find(&users)
	return c.Status(200).JSON(users)
}

func GetUsers(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User

	if err := database.DB.Db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
