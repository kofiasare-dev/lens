package models

import (
	"time"

	"github.com/gofiber/fiber/v3/client"
	"github.com/google/uuid"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/services/db"
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
	ClientID    uint      `gorm:"not null"`
	CallbackURL string    `gorm:"not null"`
	Client      *Client
}

func CreateVerificationRequest(v *VerificationRequest) (err error) {
	pg := db.GetPgClient()

	if err = pg.Create(&v).Error; err != nil {
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

func (v *VerificationRequest) ProcessVerificationRequest() {
	switch v.Type {
	case constants.PERMISSION_FACIAL_MATCH:

		c := client.New()
		c.SetTimeout(20 * time.Second)

		// referenceImg := (v.Payload["referenceImageUrl"]).(string)

		// r, err := c.Get(referenceImg)
		// if err != nil {
		// 	log.Infof(err.Error())
		// }

		// log.Info(r)
		// log.Info(v.Payload["targetImageUrl"])

		// m := face.GetFaceMatcher()

		// m.CompareFaceImages()
	}
}
