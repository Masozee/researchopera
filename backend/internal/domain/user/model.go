package user

type User struct {
	ID       uint   `gorm:"primaryKey"`
	TenantID uint   `gorm:"index;not null"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;uniqueIndex;not null"`
	Role     string `gorm:"size:64;not null"`
}
