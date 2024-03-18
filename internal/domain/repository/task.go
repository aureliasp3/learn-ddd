package repository

import (
	"context"

	"learn-ddd/internal/domain/model"
)

//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./task.go -destination=./mock/task.go

type TaskRepository interface {
	Select(ctx context.Context, id int64) (*model.Task, error)
	Fetch(ctx context.Context) ([]*model.Task, error)
	Search(ctx context.Context, cond map[model.TaskColumn]any) ([]*model.Task, error)
	Insert(ctx context.Context, t *model.Task) (*model.Task, error)
	Update(ctx context.Context, t *model.Task) (*model.Task, error)
	Delete(ctx context.Context, id int64) error
}
