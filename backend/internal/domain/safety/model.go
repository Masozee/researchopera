package safety

type IncidentReport struct {
	ID          uint   `gorm:"primaryKey"`
	TenantID    uint   `gorm:"index;not null"`
	ProjectID   uint   `gorm:"index;not null"`
	CreatedByID uint   `gorm:"index;not null"`
	Severity    string `gorm:"size:32"`
	Status      string `gorm:"size:32"`
	Description string `gorm:"type:text"`
}
