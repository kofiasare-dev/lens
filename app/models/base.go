package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	BaseModel struct {
		ID        uint      `gorm:"primaryKey;autoIncrement"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	JSONB map[string]any
)

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

var Migrations = []interface{}{
	&Client{},
	&ApiKey{},
	&VerificationRequest{},
}
