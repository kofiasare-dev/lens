package db

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgClient struct{ *gorm.DB }

var pgClient *PgClient
var once sync.Once

func GetPgClient() *PgClient {

	once.Do(func() {

		client, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")))

		if err != nil {
			log.Fatal(err.Error())
		}

		pgClient = &PgClient{client}

	})

	return pgClient
}

func (pg *PgClient) Close() {
	db, _ := pg.DB.DB()

	db.Close()
}
