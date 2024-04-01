package dbtest

import (
	"github.com/morikuni/failure"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"learn-ddd/internal/domain/model"
	"learn-ddd/lib/errctrl"
)

var TestDB *gorm.DB

func ConnectTest() error {
	db, err := gorm.Open(sqlite.Open("../../../sqlite_test.db"), &gorm.Config{})
	if err != nil {
		return failure.New(errctrl.Internal, failure.Messagef("can not establish database connection: %s", err.Error()))
	}
	TestDB = db
	return nil
}

func Setup() {
	errctrl.MustExec(ConnectTest())
	errctrl.MustExec(TestDB.Migrator().DropTable(&model.Task{}))
	errctrl.MustExec(TestDB.Migrator().DropTable(&model.User{}))
	errctrl.MustExec(TestDB.AutoMigrate(&model.User{}, &model.Task{}))
}
