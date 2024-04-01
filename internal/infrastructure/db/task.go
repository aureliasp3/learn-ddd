package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/morikuni/failure"
	"gorm.io/gorm"

	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/domain/repository"
	"learn-ddd/lib/ctxctrl"
	"learn-ddd/lib/errctrl"
)

type taskRepository struct{}

var _ repository.TaskRepository = (*taskRepository)(nil)

func NewTaskRepository() repository.TaskRepository {
	return &taskRepository{}
}

func (r taskRepository) Select(ctx context.Context, id int64) (*model.Task, error) {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return nil, failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	t := &model.Task{}
	if err := tx.First(t, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.New(errctrl.NotFound, failure.Messagef("task id %d is not found", id))
		}
		return nil, failure.New(errctrl.Internal, failure.Message(err.Error()))
	}

	return t, nil
}

func (r taskRepository) Fetch(ctx context.Context) ([]*model.Task, error) {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return nil, failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	var ts []*model.Task
	if err := tx.Find(&ts).Error; err != nil {
		return nil, failure.New(errctrl.Internal, failure.Message(err.Error()))
	}

	return ts, nil
}

func (r taskRepository) Search(ctx context.Context, cond map[model.TaskColumn]any) ([]*model.Task, error) {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return nil, failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	var ts []*model.Task
	if err := buildTaskSearchQuery(tx, cond).Find(&ts).Error; err != nil {
		return nil, failure.New(errctrl.Internal, failure.Message(err.Error()))
	}
	return ts, nil
}

func (r taskRepository) Insert(ctx context.Context, t *model.Task) (*model.Task, error) {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return nil, failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	if err := tx.Create(t).Error; err != nil {
		return nil, failure.New(errctrl.Internal, failure.Message(err.Error()))
	}

	return t, nil
}

func (r taskRepository) Update(ctx context.Context, t *model.Task) (*model.Task, error) {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return nil, failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	if err := tx.Save(t).Error; err != nil {
		return nil, failure.New(errctrl.Internal, failure.Message(err.Error()))
	}

	return t, nil
}

func (r taskRepository) Delete(ctx context.Context, id int64) error {
	tx := ctxctrl.GetTx(ctx)
	if tx == nil {
		return failure.New(errctrl.Internal, failure.Message("tx is nil"))
	}

	t := &model.Task{}
	if err := tx.Delete(t, id).Error; err != nil {
		return failure.New(errctrl.Internal, failure.Message(err.Error()))
	}

	return nil
}

func buildTaskSearchQuery(tx *gorm.DB, cond map[model.TaskColumn]any) *gorm.DB {
	if v, ok := cond[model.TaskColumnTitle]; ok {
		tx = tx.Where("title LIKE ?", fmt.Sprintf("%%%s%%", v))
	}
	if v, ok := cond[model.TaskColumnAuthorID]; ok {
		tx = tx.Where("author_id = ?", v)
	}
	if v, ok := cond[model.TaskColumnAssigneeID]; ok {
		tx = tx.Where("assignee_id = ?", v)
	}
	if v, ok := cond[model.TaskColumnStatus]; ok {
		tx = tx.Where("status = ?", v)
	}

	return tx
}
