package db

import (
	"context"
	"errors"

	"github.com/morikuni/failure"
	"gorm.io/gorm"

	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/domain/repository"
	"learn-ddd/lib/ctxctrl"
	"learn-ddd/lib/errctrl"
)

type userRepository struct{}

var _ repository.UserRepository = (*userRepository)(nil)

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r userRepository) Select(ctx context.Context, id int64) (*model.User, error) {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return nil, failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	ret := &model.User{}
	if err := tx.First(ret, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.New(errctrl.NotFound, failure.Messagef("user id %d is not found", id))
		}
		return nil, failure.New(errctrl.Internal, failure.Message(err.Error()))
	}

	return ret, nil
}

func (r userRepository) Fetch(ctx context.Context) ([]*model.User, error) {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return nil, failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	var ret []*model.User
	if err := tx.Find(&ret).Error; err != nil {
		return nil, failure.New(errctrl.Internal, failure.Message(err.Error()))
	}

	return ret, nil
}
