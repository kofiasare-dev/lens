package models

type VerificationRequest struct {
	*BaseModel
	Type      string                 `gorm:"type:varchar;not null"`                // Type of verification (e.g., "facial_recognition")
	Status    string                 `gorm:"type:varchar;default:'pending';index"` // Status of the verification (e.g., "pending", "verified")
	Reference string                 `gorm:"type:varchar(32);uniqueIndex"`         // Unique reference for the verification request
	Payload   map[string]interface{} `gorm:"type:jsonb"`                           // JSON payload with input data
	Result    map[string]interface{} `gorm:"type:jsonb"`                           // JSON result with output data
	ClientID  uint                   `gorm:"not null"`                             // Foreign key for the client
	Client    *Client
}
