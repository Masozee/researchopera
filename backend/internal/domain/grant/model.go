package grant

type GrantApplication struct {
	ID              uint   `gorm:"primaryKey"`
	TenantID        uint   `gorm:"index;not null"`
	ProjectID       uint   `gorm:"index;not null"`
	SponsorID       uint   `gorm:"index"`
	Title           string `gorm:"size:255;not null"`
	Status          string `gorm:"size:32;not null"`
	Deadline        string
	AmountRequested float64
	Document        string `gorm:"size:512"`
}
