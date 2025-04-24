package audit

type AuditChecklist struct {
	ID            uint   `gorm:"primaryKey"`
	TenantID      uint   `gorm:"index;not null"`
	ProjectID     uint   `gorm:"index;not null"`
	CompletedByID uint   `gorm:"index;not null"`
	CompletedAt   string
}

type AuditItem struct {
	ID           uint   `gorm:"primaryKey"`
	ChecklistID  uint   `gorm:"index;not null"`
	Description  string `gorm:"size:255"`
	Checked      bool
}
