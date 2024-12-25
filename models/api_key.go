package models

import (
	"crypto/sha512"
	"os"

	"github.com/kofiasare/lens-api/contollers/inputs"
	"github.com/kofiasare/lens-api/db"
	"github.com/kofiasare/lens-api/utils"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	SupportedApiKeyPermissions = []string{
		"Document Verification",
		"Facial Recognition",
	}
)

type ApiKey struct {
	*BaseModel
	Name        string         `gorm:"type:varchar(255);not null"`                        // Name of the API key
	KeyDigest   string         `gorm:"type:varchar(255);not null"`                        // Hashed value of the API key
	State       string         `gorm:"type:varchar(50);not null;default:'enabled';index"` // AASM state for the key
	Permissions pq.StringArray `gorm:"type:text[];default:'{}'"`                          // Array of permissions
	ClientID    uint           `gorm:"not null;index"`                                    // Foreign key for the client
	Client      *Client
	key         string
}

func CreateApiKey(i *inputs.CreateApiKeyInput, rawKey string) (key string, err error) {
	pg := db.GetPgClient()

	apiKey := &ApiKey{
		ClientID:    utils.StringToUint(i.ClientID),
		Permissions: i.Permissions,
		Name:        i.Name,
		key:         rawKey,
	}

	if err = pg.Create(&apiKey).Error; err != nil {
		return
	}

	key = apiKey.key

	return
}

func ApiKeyWithKey(k string) (key *ApiKey, err error) {
	pg := db.GetPgClient()

	digest, err := generateKeyDigest(k)
	if err != nil {
		return nil, err
	}

	if err = pg.Where("key_digest = ?", digest).First(&key).Error; err != nil {
		return nil, err
	}

	return key, nil
}

func (a *ApiKey) BeforeCreate(tx *gorm.DB) (err error) {
	digest, err := generateKeyDigest(a.key)
	if err != nil {
		return
	}

	a.KeyDigest = digest
	return
}

func generateKeyDigest(key string) (digest string, err error) {
	return utils.HmacHexDigest(sha512.New, os.Getenv("SECRET_KEY"), key)
}
