package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const TenantIDContextKey = "tenant_id"

// TenantMiddleware extracts TenantID from JWT or subdomain and injects into context.
func TenantMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tenantID string
		// 1. Try to get from JWT (Authorization: Bearer <token>)
		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
			if err == nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					if tid, ok := claims["tenant_id"].(string); ok {
						tenantID = tid
					}
				}
			}
		}
		// 2. Fallback: extract from subdomain (e.g., tenant1.example.com)
		if tenantID == "" {
			host := c.Hostname()
			parts := strings.Split(host, ".")
			if len(parts) > 2 {
				tenantID = parts[0]
			}
		}
		if tenantID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "TenantID not found"})
		}
		c.Locals(TenantIDContextKey, tenantID)
		return c.Next()
	}
}

// GetTenantID retrieves the TenantID from Fiber context
func GetTenantID(c *fiber.Ctx) string {
	if tid, ok := c.Locals(TenantIDContextKey).(string); ok {
		return tid
	}
	return ""
}
