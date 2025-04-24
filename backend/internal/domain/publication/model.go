package publication

type Publication struct {
	ID         uint   `gorm:"primaryKey"`
	TenantID   uint   `gorm:"index;not null"`
	Title      string `gorm:"size:255;not null"`
	Journal    string `gorm:"size:255"`
	DOI        string `gorm:"size:128"`
	PublishDate string
}

type PublicationAuthor struct {
	PublicationID uint `gorm:"primaryKey"`
	UserID        uint `gorm:"primaryKey"`
}
