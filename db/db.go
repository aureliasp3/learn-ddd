package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/morikuni/failure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"learn-ddd/lib/errctrl"
)

var DB *gorm.DB

func Connect() error {
	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		return failure.New(errctrl.Internal, failure.Messagef("can not establish database connection: %s", err.Error()))
	}
	DB = db
	return nil
}

func dsn() string {
	errctrl.MustExec(godotenv.Load(".env"))
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, pass, database, port)
}
