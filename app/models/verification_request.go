package models

import (
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/services/db"
	"github.com/kofiasare/lens/app/services/face"
	"gorm.io/gorm"
)

type VerificationRequest struct {
	*BaseModel
	*Guaranteetable
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Type        string    `gorm:"type:varchar;not null"`
	State       string    `gorm:"type:varchar;default:'pending';index"`
	Reference   string    `gorm:"type:varchar(32);uniqueIndex"`
	Result      JSONB     `gorm:"type:jsonb"`
	ClientID    uint      `gorm:"not null" json:"-"`
	CallbackURL string    `gorm:"not null"`
	Client      *Client   `json:"-"`
}

func CreateVerificationRequest(v *VerificationRequest) (err error) {
	pg := db.GetPgClient()

	if err = pg.Create(&v).Error; err != nil {
		return
	}

	return

}

func FirstOrCreateVerificationRequest(v *VerificationRequest) (err error) {
	pg := db.GetPgClient()

	if err = pg.Where(&VerificationRequest{Reference: v.Reference}).FirstOrCreate(&v).Error; err != nil {
		return
	}

	return

}

func FindVerificationRequest(id uuid.UUID) (v *VerificationRequest, err error) {
	pg := db.GetPgClient()

	if err = pg.Where(&VerificationRequest{ID: id}).First(&v).Error; err != nil {
		return
	}

	return
}

func FindVerificationByReference(r string) (v *VerificationRequest, err error) {
	pg := db.GetPgClient()

	if err = pg.Where("reference = ?", r).First(&v).Error; err != nil {
		return
	}

	return
}

func (v *VerificationRequest) AfterCreate(tx *gorm.DB) (err error) {
	return v.Guarantee(&GuaranteedExecution{
		GuaranteeableType: "VerificationRequest",
		GuaranteeableID:   v.ID.String(),
		Task:              constants.VERIFICATION_REQUEST_PROCESSING_TASK,
		Retry:             5,
	})
}

func (v *VerificationRequest) ProcessVerificationRequest() (result any, err error) {
	switch v.Type {
	case constants.PERMISSION_FACIAL_MATCH:
		bc := db.GetBadgerClient()

		ri, err := bc.Get("referenceImage:" + v.Reference)
		if err != nil {
			return nil, err
		}

		ti, err := bc.Get("targetImage:" + v.Reference)
		if err != nil {
			return nil, err
		}

		m := face.GetFaceMatcher()
		result, err = m.CompareFaceImages(ri, ti)
		if err != nil {
			return nil, err
		}

	}

	return result, nil
}

func (v *VerificationRequest) Fail(reason error) {
	data := map[string]any{"State": "failed", "ErrorReason": reason.Error()}
	v.update(data)
}

func (v *VerificationRequest) Complete(result any) {
	data := map[string]any{"State": "completed", "Result": result}
	v.update(data)
}

func (v *VerificationRequest) update(data map[string]any) {
	pg := db.GetPgClient()

	err := pg.Model(v).Updates(data).Error
	if err != nil {
		log.Error("Failed to update verification request:", err)
	}

	pg.First(v, "id = ?", v.ID)
}
