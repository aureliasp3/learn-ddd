package usecase

import (
	"context"

	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/domain/repository"
)

//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./user.go -destination=./mock/user.go

type UserUsecase interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

var _ UserUsecase = (*userUsecase)(nil)

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u userUsecase) GetUser(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.userRepo.Select(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userUsecase) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.userRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
