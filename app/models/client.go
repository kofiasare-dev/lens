package models

type (
	Client struct {
		*BaseModel
		Name                 string `gorm:"type:varchar(255);unique;not null"`
		State                string `gorm:"type:varchar(20);default:'active';index;not null"`
		VerificationRequests []*VerificationRequest
		ApiKeys              []*ApiKey
	}
)
