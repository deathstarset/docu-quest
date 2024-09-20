package app

import (
	"fmt"
	"os"

	"github.com/deathstarset/backend-docu-quest/config"
	"github.com/deathstarset/backend-docu-quest/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupAndRunApp() error {
	// Load env
	err := config.LoadEnv()
	if err != nil {
		return err
	}

	// Start database
	err = config.StartPostgres()
	if err != nil {
		return err
	}

	// Create app
	app := fiber.New()

	// Add middlewares
	app.Use(logger.New())

	// Add routes
	router.SetupRoutes(app)

	// Get port ps: Figure this shit out later
	port := os.Getenv("PORT")
	fmt.Println(port)
	app.Listen(":3000")
	return nil
}
