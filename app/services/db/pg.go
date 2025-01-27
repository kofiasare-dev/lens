package db

import (
	"log"
	"os"
	"sync"
	"time"

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
			log.Fatal(err)
		}

		sqlDB, err := client.DB()
		if err != nil {
			log.Fatal(err)
		}

		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetConnMaxLifetime(time.Hour)

		pgClient = &PgClient{client}
	})

	return pgClient
}

func (pg *PgClient) Close() {
	db, _ := pg.DB.DB()

	db.Close()
}
