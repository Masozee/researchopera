package api

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/Masozee/researchopera/internal/middleware"
	"github.com/Masozee/researchopera/internal/domain/project"
	"gorm.io/gorm"
)

// ListProjects returns all projects for the authenticated tenant
func ListProjects(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantID := middleware.GetTenantID(c)
		var projects []project.Project
		tid := parseUint(tenantID)
		if err := db.Scopes(util.TenantScope(tid)).Find(&projects).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch projects"})
		}
		return c.JSON(projects)
	}
}

// CreateProject creates a new project for the tenant
func CreateProject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantID := middleware.GetTenantID(c)
		var input project.Project
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}
		tid := parseUint(tenantID)
		input.TenantID = tid
		if err := db.Scopes(util.TenantScope(tid)).Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create project"})
		}
		return c.Status(201).JSON(input)
	}
}

// UpdateProject updates an existing project for the tenant
func UpdateProject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantID := middleware.GetTenantID(c)
		id := c.Params("id")
		var proj project.Project
		tid := parseUint(tenantID)
		if err := db.Scopes(util.TenantScope(tid)).Where("id = ?", id).First(&proj).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Project not found"})
		}
		if err := c.BodyParser(&proj); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}
		if err := db.Scopes(util.TenantScope(tid)).Save(&proj).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update project"})
		}
		return c.JSON(proj)
	}
}

// DeleteProject deletes a project for the tenant
func DeleteProject(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantID := middleware.GetTenantID(c)
		id := c.Params("id")
		tid := parseUint(tenantID)
		if err := db.Scopes(util.TenantScope(tid)).Where("id = ?", id).Delete(&project.Project{}).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete project"})
		}
		return c.SendStatus(204)
	}
}

func parseUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)
	return uint(u)
}
