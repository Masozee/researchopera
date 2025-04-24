package project

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	TenantID    uint   `gorm:"index;not null"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
}

type ProjectUser struct {
	UserID    uint `gorm:"primaryKey"`
	ProjectID uint `gorm:"primaryKey"`
	Role      string
}
