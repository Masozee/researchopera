package literature

type Literature struct {
	ID       uint   `gorm:"primaryKey"`
	TenantID uint   `gorm:"index;not null"`
	UserID   uint   `gorm:"index;not null"`
	Title    string `gorm:"size:255;not null"`
	Authors  string `gorm:"size:255"`
	Source   string `gorm:"size:255"`
	Year     int
	File     string `gorm:"size:512"`
}
