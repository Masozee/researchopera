package tenant

type Tenant struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null"`
}
