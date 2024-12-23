package models

type Client struct {
	*BaseModel
	Name                 string `gorm:"type:varchar(100);uniqueIndex;not null"`  // Unique name of the client (e.g., organization name)
	Status               string `gorm:"type:varchar(20);default:'active';index"` // Status of the client (e.g., active, inactive)
	VerificationRequests []*VerificationRequest
}
