package models

import (
	"crypto/sha512"
	"os"
	"slices"

	"github.com/kofiasare-dev/utils"
	"github.com/kofiasare/lens/app/contollers/inputs"
	"github.com/kofiasare/lens/app/services/db"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	ApiKey struct {
		*BaseModel
		Name        string         `gorm:"type:varchar(255);not null"`
		KeyDigest   string         `gorm:"type:varchar(255);not null"`
		State       string         `gorm:"type:varchar(50);not null;default:'enabled';index"`
		Permissions pq.StringArray `gorm:"type:text[];default:'{}'"`
		ClientID    uint           `gorm:"not null;index"`
		Client      *Client
		key         string
	}
)

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

	a := &ApiKey{key: k}
	digest, _ := a.generateKeyDigest()

	if err = pg.Where("key_digest = ?", digest).Preload("Client").First(&key).Error; err != nil {
		return nil, err
	}

	return key, nil
}

func (a *ApiKey) BeforeCreate(tx *gorm.DB) (err error) {
	digest, err := a.generateKeyDigest()
	if err != nil {
		return
	}

	a.KeyDigest = digest
	return
}

func (a *ApiKey) Supports(permission string) bool {
	return slices.Contains(a.Permissions, permission)
}

func (a *ApiKey) generateKeyDigest() (digest string, err error) {
	return utils.HmacHexDigest(sha512.New, os.Getenv("SECRET_KEY"), a.key)
}
