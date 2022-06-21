package database

import (
	"fmt"
	"log"

	"github.com/ja1felipe/user-service/pkg/config/database/migrations"
	"github.com/ja1felipe/user-service/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Connect() {
	password := utils.GetEnv("DB_PASSWORD")
	user := utils.GetEnv("DB_USER")
	port := utils.GetEnv("DB_PORT")
	database := utils.GetEnv("DB_NAME")

	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", user, password, port, database)

	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error on connect to database")
	}

	db.Connection = d

	migrations.RunMigrations(db.Connection)
}

func (db *DB) GetDB() *gorm.DB {
	return db.Connection
}

func Close(db *gorm.DB) error {
	config, err := db.DB()

	if err != nil {
		return err
	}

	err = config.Close()

	if err != nil {
		return err
	}

	return nil
}
