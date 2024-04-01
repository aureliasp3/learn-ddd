package memory

import (
	"context"
	"time"

	"github.com/morikuni/failure"

	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/domain/repository"
	"learn-ddd/lib/errctrl"
)

type userMemoryRepository struct{}

var _ repository.UserRepository = (*userMemoryRepository)(nil)

func NewUserMemoryRepository() repository.UserRepository {
	return &userMemoryRepository{}
}

var now = time.Now()
var users = []*model.User{
	{
		ID:        1,
		Name:      "Name1 In Memory",
		CreatedAt: now,
		UpdatedAt: now,
	},
	{
		ID:        2,
		Name:      "Name2 In Memory",
		CreatedAt: now,
		UpdatedAt: now,
	},
	{
		ID:        3,
		Name:      "Name3 In Memory",
		CreatedAt: now,
		UpdatedAt: now,
	},
}

func (r userMemoryRepository) Select(ctx context.Context, id int64) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, failure.New(errctrl.NotFound, failure.Messagef("user id %d is not found", id))
}

func (r userMemoryRepository) Fetch(ctx context.Context) ([]*model.User, error) {
	return users, nil
}
