package models

type VerificationRequest struct {
	*BaseModel
	Type             string
	ConfidenceScore  float32
	DocumentPhotoUrl string
	SelfieUrl        string
	Status           string
}
