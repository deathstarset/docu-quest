package router

import (
	"log"
	"os"

	"github.com/deathstarset/backend-docu-quest/handlers"
	"github.com/deathstarset/backend-docu-quest/middlewares"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func SetupRoutes(app *fiber.App) {

	secret := os.Getenv("JWT_SECRET")
	jwt := middlewares.AuthMiddleware(secret)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/protected", jwt, func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jtoken.Token)
		claims := user.Claims.(jtoken.MapClaims)
		userID := claims["id"].(string)
		log.Println(userID)
		return c.SendString("this is a protected route")
	})

	users := app.Group("/api/v1/users")
	users.Post("/", handlers.CreateUser)
	users.Get("/:id", handlers.GetUser)
	users.Get("/", handlers.GetUsers)
	users.Delete("/:id", handlers.DeleteUser)
	users.Put("/:id", handlers.UpdateUser)

	auth := app.Group("/api/v1/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

}
