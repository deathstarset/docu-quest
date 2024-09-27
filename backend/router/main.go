package router

import (
	"log"
	"os"

	"github.com/deathstarset/backend-docu-quest/handlers"
	"github.com/deathstarset/backend-docu-quest/middlewares"
	userhandlers "github.com/deathstarset/backend-docu-quest/userHandlers"
	"github.com/deathstarset/backend-docu-quest/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	secret := os.Getenv("JWT_SECRET")
	isAuth := middlewares.AuthMiddleware(secret)
	adminRouter := app.Group("/admin")
	adminRouter.Use(isAuth, middlewares.OnlyAdmin())

	userRouter := app.Group("/")
	userRouter.Use(isAuth)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/protected", isAuth, func(c *fiber.Ctx) error {
		userID := utils.GetUserID(c)
		log.Println(userID)
		return c.SendString("this is a protected route")
	})

	users := adminRouter.Group("/api/v1/users")
	users.Post("/", handlers.CreateUser)
	users.Get("/:id", handlers.GetUser)
	users.Get("/", handlers.GetUsers)
	users.Delete("/:id", handlers.DeleteUser)
	users.Put("/:id", handlers.UpdateUser)

	conversations := adminRouter.Group("/api/v1/conversations")
	conversations.Post("/", handlers.CreateConversation)
	conversations.Get("/:id", handlers.GetConversation)
	conversations.Delete("/:id", handlers.DeleteConversation)
	conversations.Get("/", handlers.GetConversations)

	messages := app.Group("/api/v1/messages")
	messages.Post("/", handlers.CreateMessage)
	messages.Get("/:id", handlers.GetMessage)
	messages.Delete("/:id", handlers.DeleteMessage)
	messages.Get("/", handlers.GetMessages)

	documents := adminRouter.Group("/api/v1/documents")
	documents.Post("/", handlers.CreateDocument)
	documents.Get("/:id", handlers.GetDocument)
	documents.Get("/", handlers.GetDocuments)
	documents.Delete("/:id", handlers.DeleteDocument)

	extractedContents := adminRouter.Group("/api/v1/extracted_contents")
	extractedContents.Post("/", handlers.CreateExtractedContent)
	extractedContents.Get("/:id", handlers.GetExtractedContent)
	extractedContents.Get("/", handlers.GetExtractedContents)
	extractedContents.Delete("/:id", handlers.DeleteExtractedContent)

	embeddings := adminRouter.Group("/api/v1/embeddings")
	embeddings.Post("/", handlers.CreateEmbedding)
	embeddings.Post("/similar", handlers.GetSimilarEmbeddings)

	profiles := userRouter.Group("/api/v1/profiles")
	profiles.Get("/", userhandlers.GetProfile)
	profiles.Delete("/", userhandlers.DeleteProfile)
	profiles.Put("/", userhandlers.UpdateProfile)

	auth := app.Group("/api/v1/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	userConversations := userRouter.Group("/api/v1/conversations")
	userConversations.Post("/", userhandlers.CreateUserConversation)
	userConversations.Delete("/:id", userhandlers.DeleteUserConversation)
	userConversations.Post("/:id/message", userhandlers.SendMessage)

	userDocuments := userRouter.Group("/api/v1/documents")
	userDocuments.Post("/", userhandlers.CreateUserDocument)
	userDocuments.Delete("/:id", userhandlers.DeleteUserDocument)
	userDocuments.Get("/:id", userhandlers.GetUserDocument)

	userExtractedContents := userRouter.Group("/api/v1/extracted_content")
	userExtractedContents.Post("/", userhandlers.CreateUserExtractedContent)
	userExtractedContents.Delete("/:id", userhandlers.DeleteUserExtractedContent)

}
