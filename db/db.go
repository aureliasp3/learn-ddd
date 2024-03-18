package db

import (
	"github.com/morikuni/failure"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"learn-ddd/lib/errctrl"
)

var DB *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return failure.New(errctrl.Internal, failure.Messagef("can not establish database connection: %s", err.Error()))
	}
	DB = db
	return nil
}
