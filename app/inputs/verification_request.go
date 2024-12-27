package inputs

import (
	"github.com/kofiasare/lens/app/constants"
)

type (
	BaseVerificationInput struct {
		Type        string `json:"type" validate:"required,oneof=FACIAL_RECOGNITION"`
		Reference   string `json:"reference" validate:"required,len=32"`
		CallbackURL string `json:"callbackUrl" validate:"required,url"`
	}

	FacialRecognitionInput struct {
		*BaseVerificationInput
		Payload struct {
			TargetImageURL    string `json:"targetImageUrl" validate:"required,url"`
			ReferenceImageURL string `json:"referenceImageUrl" validate:"required,url"`
		}
	}

	VerificationSpecificInput interface{ BaseInput() }
)

func (b *BaseVerificationInput) LoadSpecificInput() VerificationSpecificInput {
	switch b.Type {
	case constants.PERMISSION_FACIAL_RECOGNITION:
		return &FacialRecognitionInput{BaseVerificationInput: b}
	default:
		return nil
	}
}

func (f FacialRecognitionInput) BaseInput() {}
