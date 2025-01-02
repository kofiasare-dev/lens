package models

import (
	"encoding/json"
	"time"

	"github.com/kofiasare-dev/background"
	"github.com/kofiasare/lens/app/services/bg"
	"github.com/kofiasare/lens/app/services/db"
	"gorm.io/gorm"
)

type (
	GuaranteedExecution struct {
		ID                uint      `gorm:"primaryKey;autoIncrement"`
		CreatedAt         time.Time `gorm:"autoCreateTime"`
		UpdatedAt         time.Time `gorm:"autoUpdateTime"`
		GuaranteeableType string    `gorm:"type:varchar;not null"`
		GuaranteeableID   string    `gorm:"type:varchar;not null"`
		CompletedAt       time.Time `gorm:"type:timestamp"`
		Task              string    `gorm:"type:varchar;not null"`
		Queue             string    `gorm:"type:varchar;default:'default';not null"`
		Retry             uint
		Wait              time.Duration
	}

	Guaranteetable struct {
		GuaranteedExecutions []*GuaranteedExecution `gorm:"polymorphic:Guaranteeable"`
	}
)

func (bg *Guaranteetable) Guarantee(ge *GuaranteedExecution) (err error) {
	pg := db.GetPgClient()

	if err = pg.Create(&ge).Error; err != nil {
		return err
	}

	return
}

func (ge *GuaranteedExecution) AfterCreate(tx *gorm.DB) (err error) {
	bg := bg.GetInstance()

	payload, err := json.Marshal(ge)
	if err != nil {
		return
	}

	task := bg.NewTask(ge.Task, payload, background.TaskOptions{
		Retry: ge.Retry,
		Queue: ge.Queue,
		Wait:  ge.Wait,
	})

	if err = bg.EnqueueTask(task); err != nil {
		return
	}

	return
}
