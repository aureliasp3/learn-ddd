package repository

import (
	"context"

	"learn-ddd/internal/domain/model"
)

//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./user.go -destination=./mock/user.go

type UserRepository interface {
	Select(ctx context.Context, id int64) (*model.User, error)
	Fetch(ctx context.Context) ([]*model.User, error)
}
