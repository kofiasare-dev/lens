package inputs

type (
	BaseVerificationInput struct {
		Type        string `json:"type" validate:"required,oneofci= facial_recognition"`
		Reference   string `json:"reference" validate:"required,len=32"`
		CallbackURL string `json:"callbackUrl" validate:"required,url"`
	}

	FacialRecognitionInput struct {
		*BaseVerificationInput
		Payload struct {
			SelfieURL           string  `json:"selfieUrl" validate:"required,url"`
			DocumentPhotoURL    string  `json:"documentPhotoUrl" validate:"required,url"`
			ConfidenceThreshold float64 `json:"confidenceThreshold" validate:"required,gte=0,lte=1"`
		}
	}

	VerificationRequestInput interface {
		GetBaseInput() *BaseVerificationInput
	}
)

func (i *FacialRecognitionInput) GetBaseInput() *BaseVerificationInput {
	return i.BaseVerificationInput
}
