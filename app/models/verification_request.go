package models

import (
	"time"

	"github.com/gofiber/fiber/v3/client"
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/contollers/inputs"
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
	Payload     JSONB     `gorm:"type:jsonb"`
	Result      JSONB     `gorm:"type:jsonb"`
	ClientID    uint      `gorm:"not null"`
	CallbackURL string    `gorm:"not null"`
	Client      *Client
}

func CreateVerificationRequest(c *Client, si inputs.VerificationSpecificInput) (v *VerificationRequest, err error) {
	pg := db.GetPgClient()

	v = &VerificationRequest{Client: c}
	v.buildRequest(si)

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

func (v *VerificationRequest) buildRequest(si inputs.VerificationSpecificInput) {
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

func (v *VerificationRequest) ProcessVerificationRequest() {
	switch v.Type {
	case constants.PERMISSION_FACIAL_RECOGNITION:

		c := client.New()
		c.SetTimeout(20 * time.Second)

		referenceImg := (v.Payload["referenceImageUrl"]).(string)

		r, err := c.Get(referenceImg)
		if err != nil {
			log.Infof(err.Error())
		}

		log.Info(r)
		log.Info(v.Payload["targetImageUrl"])

		// m := face.GetFaceMatcher()

		// m.CompareFaceImages()
	}
}
