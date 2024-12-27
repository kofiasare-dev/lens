package inputs

type CreateApiKeyInput struct {
	ClientID    string   `json:"client_id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Permissions []string `json:"permissions"`
}
