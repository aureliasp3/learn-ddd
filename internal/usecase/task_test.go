package usecase

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	"learn-ddd/db"
	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/domain/repository/mock"
	"learn-ddd/internal/fixture"
	"learn-ddd/lib/ctxctrl"
)

func Test_taskUsecase_CreateTask(t *testing.T) {
	type args struct {
		ctx  context.Context
		task *model.Task
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), db.DB)
				}(),
				task: fixture.FakeTask,
			},
			want:    fixture.FakeTask,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskRepo := mock.NewMockTaskRepository(ctrl)
			mockTaskRepo.EXPECT().Insert(tt.args.ctx, tt.args.task).Return(fixture.FakeTask, nil)

			mockUserRepo := mock.NewMockUserRepository(ctrl)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTask.AuthorID).Return(fixture.FakeUser, nil)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTask.AssigneeID).Return(fixture.FakeUser, nil)

			u := NewTaskUseCase(mockTaskRepo, mockUserRepo)

			got, err := u.CreateTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskUsecase_DeleteTask(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), db.DB)
				}(),
				id: 1,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskRepo := mock.NewMockTaskRepository(ctrl)
			mockTaskRepo.EXPECT().Delete(tt.args.ctx, tt.args.id).Return(nil)

			mockUserRepo := mock.NewMockUserRepository(ctrl)

			u := NewTaskUseCase(mockTaskRepo, mockUserRepo)

			got := u.DeleteTask(tt.args.ctx, tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskUsecase_GetTask(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), db.DB)
				}(),
				id: 1,
			},
			want:    fixture.FakeTask,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskRepo := mock.NewMockTaskRepository(ctrl)
			mockTaskRepo.EXPECT().Select(tt.args.ctx, tt.args.id).Return(fixture.FakeTask, nil)

			mockUserRepo := mock.NewMockUserRepository(ctrl)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTask.AuthorID).Return(fixture.FakeUser, nil)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTask.AssigneeID).Return(fixture.FakeUser, nil)

			u := NewTaskUseCase(mockTaskRepo, mockUserRepo)

			got, err := u.GetTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskUsecase_GetTasks(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		want    []*model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), db.DB)
				}(),
				id: 1,
			},
			want:    fixture.FakeTasksWithoutUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskRepo := mock.NewMockTaskRepository(ctrl)
			mockTaskRepo.EXPECT().Fetch(tt.args.ctx).Return(fixture.FakeTasksWithoutUser, nil)

			mockUserRepo := mock.NewMockUserRepository(ctrl)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTasksWithoutUser[0].AuthorID).Return(fixture.FakeTasksWithoutUser[0].Author, nil)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTasksWithoutUser[0].AssigneeID).Return(fixture.FakeTasksWithoutUser[1].Assignee, nil)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTasksWithoutUser[1].AuthorID).Return(fixture.FakeTasksWithoutUser[1].Author, nil)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTasksWithoutUser[1].AssigneeID).Return(fixture.FakeTasksWithoutUser[0].Assignee, nil)

			u := NewTaskUseCase(mockTaskRepo, mockUserRepo)

			got, err := u.GetTasks(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskUsecase_SearchTasks(t *testing.T) {
	type args struct {
		ctx  context.Context
		cond map[model.TaskColumn]any
	}

	tests := []struct {
		name    string
		args    args
		want    []*model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), db.DB)
				}(),
				cond: map[model.TaskColumn]any{
					model.TaskColumnTitle:      "2",
					model.TaskColumnAuthorID:   fixture.FakeTasksForSearchWithoutUser[1].AuthorID,
					model.TaskColumnAssigneeID: fixture.FakeTasksForSearchWithoutUser[1].AssigneeID,
					model.TaskColumnStatus:     model.TaskStatusDone,
				},
			},
			want:    []*model.Task{fixture.FakeTasksForSearchWithoutUser[1]},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskRepo := mock.NewMockTaskRepository(ctrl)
			mockTaskRepo.EXPECT().Search(tt.args.ctx, tt.args.cond).Return([]*model.Task{fixture.FakeTasksForSearchWithoutUser[1]}, nil)

			mockUserRepo := mock.NewMockUserRepository(ctrl)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTasksForSearchWithoutUser[1].AuthorID).Return(fixture.FakeUsers[1], nil)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTasksForSearchWithoutUser[1].AssigneeID).Return(fixture.FakeUsers[0], nil)

			u := NewTaskUseCase(mockTaskRepo, mockUserRepo)

			got, err := u.SearchTasks(tt.args.ctx, tt.args.cond)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskUsecase_UpdateTask(t *testing.T) {
	type args struct {
		ctx  context.Context
		task *model.Task
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), db.DB)
				}(),
				task: fixture.FakeTaskWithoutUser,
			},
			want:    fixture.FakeTaskWithoutUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskRepo := mock.NewMockTaskRepository(ctrl)
			mockTaskRepo.EXPECT().Update(tt.args.ctx, tt.args.task).Return(fixture.FakeTaskWithoutUser, nil)

			mockUserRepo := mock.NewMockUserRepository(ctrl)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTaskWithoutUser.AuthorID).Return(fixture.FakeUser, nil)
			mockUserRepo.EXPECT().Select(tt.args.ctx, fixture.FakeTaskWithoutUser.AssigneeID).Return(fixture.FakeUser, nil)

			u := NewTaskUseCase(mockTaskRepo, mockUserRepo)

			got, err := u.UpdateTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}
