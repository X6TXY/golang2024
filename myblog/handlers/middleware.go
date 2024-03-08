package handlers

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gofiber/fiber/v2"
    "strings"
)

// JWTMiddleware verifies the JWT token in the request
func JWTMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No authorization header provided"})
    }

    headerParts := strings.Split(authHeader, " ")
    if len(headerParts) != 2 || headerParts[0] != "Bearer" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid authorization header format"})
    }

    tokenString := headerParts[1]
    token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        // Ensure the token method conforms to "SigningMethodHMAC"
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fiber.ErrUnauthorized
        }

        return jwtSecret, nil
    })

    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired JWT token"})
    }

    if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
        // Optionally, you can add logic here to verify any claims you require, such as UserID or roles
        // For example, attaching user info to the request context
        c.Locals("userID", claims.Subject)
        return c.Next()
    } else {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired JWT token"})
    }
}
