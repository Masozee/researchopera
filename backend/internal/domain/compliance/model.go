package compliance

type Policy struct {
	ID          uint   `gorm:"primaryKey"`
	TenantID    uint   `gorm:"index;not null"`
	Title       string `gorm:"size:255;not null"`
	DocumentPath string `gorm:"size:512"`
	Active      bool
}

type PolicyAcceptance struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  uint   `gorm:"index;not null"`
	PolicyID  uint   `gorm:"index;not null"`
	UserID    uint   `gorm:"index;not null"`
	AcceptedAt string
}
