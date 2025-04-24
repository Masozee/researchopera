package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/Masozee/researchopera/internal/middleware"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	app := fiber.New()
	// Register tenant extraction middleware globally
	app.Use(middleware.TenantMiddleware())

	// Inject db into handlers via closure
	projectGroup := app.Group("/projects")
	projectGroup.Get("/", ListProjects(db))
	projectGroup.Post("/", CreateProject(db))
	projectGroup.Put("/:id", UpdateProject(db))
	projectGroup.Delete("/:id", DeleteProject(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
