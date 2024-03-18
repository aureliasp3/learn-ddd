package main

import (
	"time"

	"learn-ddd/db"
	"learn-ddd/internal/domain/model"
	"learn-ddd/lib/errctrl"
)

func init() {
	time.Local = time.FixedZone("Local", 9*60*60)
}

func main() {
	us := []*model.User{
		{ID: 1, Name: "name1"},
		{ID: 2, Name: "name3"},
		{ID: 3, Name: "name3"},
	}
	errctrl.MustExec(db.Connect())
	errctrl.MustExec(db.DB.Migrator().DropTable(&model.Task{}))
	errctrl.MustExec(db.DB.Migrator().DropTable(&model.User{}))
	errctrl.MustExec(db.DB.AutoMigrate(&model.User{}, &model.Task{}))
	errctrl.MustExec(db.DB.Create(us).Error)
}
