package repository

import (

    "github.com/Masozee/researchopera/internal/domain/project"
    "github.com/Masozee/researchopera/internal/util"
    "gorm.io/gorm"
)

type projectRepositoryGorm struct{}

func NewProjectRepository() ProjectRepository {
    return &projectRepositoryGorm{}
}

func (r *projectRepositoryGorm) ListByTenant(db *gorm.DB, tenantID uint) ([]project.Project, error) {
    var projects []project.Project
    err := db.Scopes(util.TenantScope(tenantID)).Find(&projects).Error
    return projects, err
}

func (r *projectRepositoryGorm) GetByID(db *gorm.DB, tenantID, id uint) (*project.Project, error) {
    var proj project.Project
    err := db.Scopes(util.TenantScope(tenantID)).Where("id = ?", id).First(&proj).Error
    if err != nil {
        return nil, err
    }
    return &proj, nil
}

func (r *projectRepositoryGorm) Create(db *gorm.DB, proj *project.Project) error {
    return db.Scopes(util.TenantScope(proj.TenantID)).Create(proj).Error
}

func (r *projectRepositoryGorm) Update(db *gorm.DB, proj *project.Project) error {
    return db.Scopes(util.TenantScope(proj.TenantID)).Save(proj).Error
}

func (r *projectRepositoryGorm) Delete(db *gorm.DB, tenantID, id uint) error {
    return db.Scopes(util.TenantScope(tenantID)).Where("id = ?", id).Delete(&project.Project{}).Error
}
