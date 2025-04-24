package sponsor

type Sponsor struct {
	ID         uint   `gorm:"primaryKey"`
	TenantID   uint   `gorm:"index;not null"`
	Name       string `gorm:"size:255;not null"`
	Category   string `gorm:"size:64"`
	Region     string `gorm:"size:128"`
	FocusAreas string `gorm:"size:255"`
	Eligibility string `gorm:"size:255"`
}
