package models

type Client struct {
	*BaseModel
	Name                 string `gorm:"type:varchar(255);unique;not null"`                // Client name, must be unique
	State                string `gorm:"type:varchar(20);default:'active';index;not null"` // State for AASM (state machine)
	VerificationRequests []*VerificationRequest
	ApiKeys              []*ApiKey
}
