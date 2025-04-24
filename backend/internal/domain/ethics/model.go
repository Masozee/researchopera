package ethics

type EthicsApproval struct {
	ID           uint   `gorm:"primaryKey"`
	TenantID     uint   `gorm:"index;not null"`
	ProjectID    uint   `gorm:"index;not null"`
	ApproverName string `gorm:"size:255"`
	DocumentPath string `gorm:"size:512"`
	Date         string
}
