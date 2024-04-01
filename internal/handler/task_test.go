package handler

import (
	"context"
	"reflect"
	"testing"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"

	"learn-ddd/db"
	apiv1 "learn-ddd/gen/api/v1"
	"learn-ddd/internal/convert"
	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/fixture"
	"learn-ddd/internal/usecase/mock"
	"learn-ddd/lib/ctxctrl"
	"learn-ddd/lib/errctrl"
)

func Test_taskServiceHandler_CreateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[apiv1.CreateTaskRequest]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[apiv1.CreateTaskResponse]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[apiv1.CreateTaskRequest]{
					Msg: &apiv1.CreateTaskRequest{
						Title:      "Title1",
						AuthorId:   1,
						AssigneeId: 1,
						Status:     apiv1.TaskStatus_TASK_STATUS_OPEN,
						DueDate:    convert.ToConnectDueDate(fixture.FakeTask.DueDate),
					},
				},
			},
			want:    connect.NewResponse(&apiv1.CreateTaskResponse{Task: convert.ToConnectTask(fixture.FakeTask)}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskUsecase := mock.NewMockTaskUseCase(ctrl)
			mockTaskUsecase.EXPECT().CreateTask(tt.args.ctx, convert.ToModelTaskFromRequest(tt.args.c.Msg)).Return(fixture.FakeTask, nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewTaskServiceHandler(mockTaskUsecase, validator)

			got, err := h.CreateTask(tt.args.ctx, tt.args.c)
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

func Test_taskServiceHandler_DeleteTask(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[apiv1.DeleteTaskRequest]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[emptypb.Empty]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[apiv1.DeleteTaskRequest]{
					Msg: &apiv1.DeleteTaskRequest{Id: 1},
				},
			},
			want:    connect.NewResponse(&emptypb.Empty{}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskUsecase := mock.NewMockTaskUseCase(ctrl)
			mockTaskUsecase.EXPECT().DeleteTask(tt.args.ctx, tt.args.c.Msg.Id).Return(nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewTaskServiceHandler(mockTaskUsecase, validator)

			got, err := h.DeleteTask(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskServiceHandler_GetTask(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[apiv1.GetTaskRequest]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[apiv1.GetTaskResponse]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[apiv1.GetTaskRequest]{
					Msg: &apiv1.GetTaskRequest{Id: 1},
				},
			},
			want:    connect.NewResponse(&apiv1.GetTaskResponse{Task: convert.ToConnectTask(fixture.FakeTask)}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskUsecase := mock.NewMockTaskUseCase(ctrl)
			mockTaskUsecase.EXPECT().GetTask(tt.args.ctx, tt.args.c.Msg.Id).Return(fixture.FakeTask, nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewTaskServiceHandler(mockTaskUsecase, validator)

			got, err := h.GetTask(tt.args.ctx, tt.args.c)
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

func Test_taskServiceHandler_GetTasks(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[emptypb.Empty]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[apiv1.GetTasksResponse]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[emptypb.Empty]{},
			},
			want:    connect.NewResponse(&apiv1.GetTasksResponse{Tasks: convert.ToConnectTasks(fixture.FakeTasks)}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskUsecase := mock.NewMockTaskUseCase(ctrl)
			mockTaskUsecase.EXPECT().GetTasks(tt.args.ctx).Return(fixture.FakeTasks, nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewTaskServiceHandler(mockTaskUsecase, validator)

			got, err := h.GetTasks(tt.args.ctx, tt.args.c)
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

func Test_taskServiceHandler_SearchTasks(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[apiv1.SearchTasksRequest]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[apiv1.SearchTasksResponse]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[apiv1.SearchTasksRequest]{
					Msg: &apiv1.SearchTasksRequest{
						Title:      toPtr("Title"),
						AuthorId:   toPtr(int64(2)),
						AssigneeId: toPtr(int64(1)),
						Status:     nil,
					},
				},
			},
			want:    connect.NewResponse(&apiv1.SearchTasksResponse{Tasks: convert.ToConnectTasks([]*model.Task{fixture.FakeTasksForSearch[1]})}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskUsecase := mock.NewMockTaskUseCase(ctrl)
			mockTaskUsecase.EXPECT().SearchTasks(tt.args.ctx, convert.ToModelTaskConditionFromRequest(tt.args.c.Msg)).Return([]*model.Task{fixture.FakeTasksForSearch[1]}, nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewTaskServiceHandler(mockTaskUsecase, validator)

			got, err := h.SearchTasks(tt.args.ctx, tt.args.c)
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

func Test_taskServiceHandler_UpdateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[apiv1.UpdateTaskRequest]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[apiv1.UpdateTaskResponse]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[apiv1.UpdateTaskRequest]{
					Msg: &apiv1.UpdateTaskRequest{
						Task: &apiv1.Task{
							Id:       1,
							Title:    "Title1",
							Author:   convert.ToConnectUser(fixture.FakeTask.Author),
							Assignee: convert.ToConnectUser(fixture.FakeTask.Assignee),
							Status:   apiv1.TaskStatus_TASK_STATUS_OPEN,
							DueDate:  convert.ToConnectDueDate(fixture.FakeTask.DueDate),
						},
					},
				},
			},
			want:    connect.NewResponse(&apiv1.UpdateTaskResponse{Task: convert.ToConnectTask(fixture.FakeTask)}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTaskUsecase := mock.NewMockTaskUseCase(ctrl)
			mockTaskUsecase.EXPECT().UpdateTask(tt.args.ctx, convert.ToModelTask(tt.args.c.Msg.Task)).Return(fixture.FakeTask, nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewTaskServiceHandler(mockTaskUsecase, validator)

			got, err := h.UpdateTask(tt.args.ctx, tt.args.c)
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

func toPtr[T any](v T) *T {
	return &v
}
