package models

import (
	"github.com/kofiasare/lens/app/inputs"
	"github.com/kofiasare/lens/db"
)

type VerificationRequest struct {
	*BaseModel
	Type        string `gorm:"type:varchar;not null"`
	State       string `gorm:"type:varchar;default:'pending';index"`
	Reference   string `gorm:"type:varchar(32);uniqueIndex"`
	Payload     JSONB  `gorm:"type:jsonb"`
	Result      JSONB  `gorm:"type:jsonb"`
	ClientID    uint   `gorm:"not null"`
	CallbackURL string `gorm:"not null"`
	Client      *Client
}

func CreateVerificationRequest(client *Client, si inputs.VerificationSpecificInput) (v *VerificationRequest, err error) {
	pg := db.GetPgClient()

	v = &VerificationRequest{Client: client}
	v.buildRequestWithSpecificInput(si)

	if err = pg.Create(&v).Error; err != nil {
		return
	}

	return
}

func FindVerificationRequest(vr *VerificationRequest) (v *VerificationRequest, err error) {
	pg := db.GetPgClient()

	if err = pg.Where("reference = ?", vr.Reference).First(&vr).Error; err != nil {
		return
	}

	return
}

func (v *VerificationRequest) buildRequestWithSpecificInput(si inputs.VerificationSpecificInput) {
	v.Payload = make(map[string]interface{})

	switch i := si.(type) {
	case *inputs.FacialRecognitionInput:
		v.Type = i.Type
		v.Reference = i.Reference
		v.CallbackURL = i.CallbackURL
		v.Payload["referenceImageUrl"] = i.Payload.ReferenceImageURL
		v.Payload["targetImageUrl"] = i.Payload.TargetImageURL
	}
}
