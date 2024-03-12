package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/x6txy/golang2024/handlers"
)

func setupRoutes(app *fiber.App) {

	app.Post("/signup", handlers.Signup)
	app.Post("/signin", handlers.Signin)

	app.Post("/post", handlers.JWTMiddleware, handlers.CreatePost)
	app.Get("/post", handlers.ListPosts)
	app.Get("/post/:id", handlers.GetPost)
	app.Delete("/post/:id", handlers.JWTMiddleware, handlers.DeletePost)
	app.Put("/post/:id", handlers.JWTMiddleware, handlers.UpdatePost)
	app.Post("/post/:id/like", handlers.JWTMiddleware, handlers.LikePost)
	app.Post("/post/:id/unlike", handlers.JWTMiddleware, handlers.UnlikePost)

	app.Get("/users", handlers.JWTMiddleware, handlers.ListUsers)
	app.Get("/users/:id", handlers.JWTMiddleware, handlers.GetUsers)
	app.Delete("/users/:id", handlers.JWTMiddleware, handlers.DeleteUser)
	app.Post("/users/:id/follow", handlers.JWTMiddleware, handlers.FollowUser)
	app.Post("/users/:id/unfollow", handlers.JWTMiddleware, handlers.UnfollowUser)

	
	app.Post("/comment/:id", handlers.JWTMiddleware, handlers.CreateComment)
	app.Get("/comment", handlers.ListComments)
	app.Get("/comment/:id", handlers.GetComment)
	app.Put("/comment/:id", handlers.JWTMiddleware, handlers.UpdateComment)
	app.Delete("/comment/:id", handlers.JWTMiddleware, handlers.DeleteComment)
	app.Post("/comment/:id/like", handlers.JWTMiddleware, handlers.LikeComment)
	app.Post("/comment/:id/unlike", handlers.JWTMiddleware, handlers.UnlikeComment)
}
