package usecase

import (
	"context"

	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/domain/repository"
)

//go:generate mkdir -p mock
//go:generate mockgen -package=mock -source=./task.go -destination=./mock/task.go

type TaskUseCase interface {
	GetTask(ctx context.Context, id int64) (*model.Task, error)
	GetTasks(ctx context.Context) ([]*model.Task, error)
	SearchTasks(ctx context.Context, cond map[model.TaskColumn]any) ([]*model.Task, error)
	CreateTask(ctx context.Context, t *model.Task) (*model.Task, error)
	UpdateTask(ctx context.Context, t *model.Task) (*model.Task, error)
	DeleteTask(ctx context.Context, id int64) error
}

type taskUseCase struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

var _ TaskUseCase = (*taskUseCase)(nil)

func NewTaskUseCase(tr repository.TaskRepository, ur repository.UserRepository) TaskUseCase {
	return &taskUseCase{
		taskRepo: tr,
		userRepo: ur,
	}
}

func (u taskUseCase) GetTask(ctx context.Context, id int64) (*model.Task, error) {
	task, err := u.taskRepo.Select(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = u.embedUserIntoTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (u taskUseCase) GetTasks(ctx context.Context) ([]*model.Task, error) {
	tasks, err := u.taskRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	if err = u.embedUserIntoTasks(ctx, tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u taskUseCase) SearchTasks(ctx context.Context, cond map[model.TaskColumn]any) ([]*model.Task, error) {
	tasks, err := u.taskRepo.Search(ctx, cond)
	if err != nil {
		return nil, err
	}

	if err = u.embedUserIntoTasks(ctx, tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u taskUseCase) CreateTask(ctx context.Context, t *model.Task) (*model.Task, error) {
	task, err := u.taskRepo.Insert(ctx, t)
	if err != nil {
		return nil, err
	}

	if err = u.embedUserIntoTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (u taskUseCase) UpdateTask(ctx context.Context, t *model.Task) (*model.Task, error) {
	task, err := u.taskRepo.Update(ctx, t)
	if err != nil {
		return nil, err
	}

	if err = u.embedUserIntoTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (u taskUseCase) DeleteTask(ctx context.Context, id int64) error {
	if err := u.taskRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (u taskUseCase) embedUserIntoTasks(ctx context.Context, ts []*model.Task) error {
	for _, t := range ts {
		err := u.embedUserIntoTask(ctx, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u taskUseCase) embedUserIntoTask(ctx context.Context, t *model.Task) error {
	author, err := u.userRepo.Select(ctx, t.AuthorID)
	if err != nil {
		return err
	}

	assignee, err := u.userRepo.Select(ctx, t.AssigneeID)
	if err != nil {
		return err
	}

	t.Author = author
	t.Assignee = assignee

	return nil
}
