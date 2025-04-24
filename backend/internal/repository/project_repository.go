package repository

import (
    "github.com/Masozee/researchopera/internal/domain/project"
    "gorm.io/gorm"
)

type ProjectRepository interface {
    ListByTenant(db *gorm.DB, tenantID uint) ([]project.Project, error)
    GetByID(db *gorm.DB, tenantID, id uint) (*project.Project, error)
    Create(db *gorm.DB, proj *project.Project) error
    Update(db *gorm.DB, proj *project.Project) error
    Delete(db *gorm.DB, tenantID, id uint) error
}
