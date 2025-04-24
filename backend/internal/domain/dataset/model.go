package dataset

type Dataset struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  uint   `gorm:"index;not null"`
	ProjectID uint   `gorm:"index;not null"`
	FilePath  string `gorm:"size:512;not null"`
	Version   string `gorm:"size:64"`
}
