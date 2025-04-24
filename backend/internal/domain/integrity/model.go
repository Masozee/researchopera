package integrity

type IntegrityDeclaration struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  uint   `gorm:"index;not null"`
	UserID    uint   `gorm:"index;not null"`
	ProjectID uint   `gorm:"index;not null"`
	Statement string `gorm:"type:text"`
	SignedAt  string
}

type ConflictOfInterest struct {
	ID          uint   `gorm:"primaryKey"`
	TenantID    uint   `gorm:"index;not null"`
	UserID      uint   `gorm:"index;not null"`
	ProjectID   *uint  `gorm:"index"`
	Description string `gorm:"type:text"`
}
