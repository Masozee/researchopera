package util

import "gorm.io/gorm"

// TenantScope returns a GORM scope that restricts queries to the given tenant ID.
func TenantScope(tenantID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("tenant_id = ?", tenantID)
	}
}
