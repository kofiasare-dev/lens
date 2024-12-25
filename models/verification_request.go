package models

import (
	"github.com/kofiasare/lens-api/contollers/inputs"
	"github.com/kofiasare/lens-api/db"
)

type VerificationRequest struct {
	*BaseModel
	Type        string                 `gorm:"type:varchar;not null"`                // Type of verification (e.g., "facial_recognition")
	State       string                 `gorm:"type:varchar;default:'pending';index"` // Status of the verification (e.g., "pending", "verified")
	Reference   string                 `gorm:"type:varchar(32);uniqueIndex"`         // Unique reference for the verification request
	Payload     map[string]interface{} `gorm:"type:jsonb"`                           // JSON payload with input data
	Result      map[string]interface{} `gorm:"type:jsonb"`                           // JSON result with output data
	ClientID    uint                   `gorm:"not null"`                             // Foreign key for the client
	CallbackURL string                 `gorm:"not null"`
	Client      *Client
}

func CreateVerificationRequest(i inputs.VerificationRequestInput) (vr *VerificationRequest, err error) {
	pg := db.GetPgClient()

	baseInput := i.GetBaseInput()

	switch baseInput.Type {
	case "facial_recognition":
		fi := i.(*inputs.FacialRecognitionInput)

		vr = &VerificationRequest{
			Type:        baseInput.Type,
			CallbackURL: baseInput.CallbackURL,
			Reference:   baseInput.Reference,
			Payload: map[string]interface{}{
				"selfieUrl":           fi.Payload.SelfieURL,
				"documentPhotoUrl":    fi.Payload.DocumentPhotoURL,
				"confidenceThreshold": fi.Payload.ConfidenceThreshold,
			},
		}
	}

	if err = pg.Create(&vr).Error; err != nil {
		return
	}

	return

}
