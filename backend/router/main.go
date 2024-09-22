package router

import (
	"log"
	"os"

	"github.com/deathstarset/backend-docu-quest/handlers"
	"github.com/deathstarset/backend-docu-quest/middlewares"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	secret := os.Getenv("JWT_SECRET")
	isAuth := middlewares.AuthMiddleware(secret)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/protected", isAuth, func(c *fiber.Ctx) error {
		userID := utils.GetUserID(c)
		log.Println(userID)
		return c.SendString("this is a protected route")
	})

	users := app.Group("/api/v1/users")
	users.Post("/", handlers.CreateUser)
	users.Get("/:id", handlers.GetUser)
	users.Get("/", handlers.GetUsers)
	users.Delete("/:id", handlers.DeleteUser)
	users.Put("/:id", handlers.UpdateUser)

	profiles := app.Group("/api/v1/profiles")
	profiles.Get("/", isAuth, handlers.GetProfile)
	profiles.Delete("/", isAuth, handlers.DeleteUser)
	profiles.Put("/", isAuth, handlers.UpdateUser)

	auth := app.Group("/api/v1/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	conversations := app.Group("/api/v1/conversations")
	conversations.Post("/", handlers.CreateConversation)
	conversations.Get("/:id", handlers.GetConversation)
	conversations.Delete("/:id", handlers.DeleteConversation)
	conversations.Get("/", handlers.GetConversations)

	messages := app.Group("/api/v1/messages")
	messages.Post("/", handlers.CreateMessage)
	messages.Get("/:id", handlers.GetMessage)
	messages.Delete("/:id", handlers.DeleteMessage)
	messages.Get("/", handlers.GetMessages)

	documents := app.Group("/api/v1/documents")
	documents.Post("/", handlers.CreateDocument)
	documents.Get("/:id", handlers.GetDocument)
	documents.Get("/", handlers.GetDocuments)
	documents.Delete("/:id", handlers.DeleteDocument)

	extractedContents := app.Group("/api/v1/extracted_contents")
	extractedContents.Post("/", handlers.CreateExtractedContent)
	extractedContents.Get("/:id", handlers.GetExtractedContent)
	extractedContents.Get("/", handlers.GetExtractedContents)
	extractedContents.Delete("/:id", handlers.DeleteExtractedContent)
}
