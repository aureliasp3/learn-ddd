package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/morikuni/failure"
	"google.golang.org/protobuf/types/known/emptypb"

	apiv1 "learn-ddd/gen/api/v1"
	"learn-ddd/gen/api/v1/apiv1connect"
	"learn-ddd/internal/convert"
	"learn-ddd/internal/usecase"
	"learn-ddd/lib/errctrl"
)

type taskServiceHandler struct {
	useCase   usecase.TaskUseCase
	validator *protovalidate.Validator
}

var _ apiv1connect.TaskServiceHandler = (*taskServiceHandler)(nil)

func NewTaskServiceHandler(u usecase.TaskUseCase, v *protovalidate.Validator) apiv1connect.TaskServiceHandler {
	return &taskServiceHandler{
		useCase:   u,
		validator: v,
	}
}

func (h taskServiceHandler) GetTask(ctx context.Context, c *connect.Request[apiv1.GetTaskRequest]) (*connect.Response[apiv1.GetTaskResponse], error) {
	if err := h.validator.Validate(c.Msg); err != nil {
		return nil, failure.New(errctrl.InvalidArgument, failure.Message(err.Error()))
	}

	task, err := h.useCase.GetTask(ctx, c.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetTaskResponse{
		Task: convert.ToConnectTask(task),
	})

	return res, nil
}

func (h taskServiceHandler) GetTasks(ctx context.Context, c *connect.Request[emptypb.Empty]) (*connect.Response[apiv1.GetTasksResponse], error) {
	if err := h.validator.Validate(c.Msg); err != nil {
		return nil, failure.New(errctrl.InvalidArgument, failure.Message(err.Error()))
	}

	tasks, err := h.useCase.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetTasksResponse{
		Tasks: convert.ToConnectTasks(tasks),
	})

	return res, nil
}

func (h taskServiceHandler) SearchTasks(ctx context.Context, c *connect.Request[apiv1.SearchTasksRequest]) (*connect.Response[apiv1.SearchTasksResponse], error) {
	if err := h.validator.Validate(c.Msg); err != nil {
		return nil, failure.New(errctrl.InvalidArgument, failure.Message(err.Error()))
	}

	tasks, err := h.useCase.SearchTasks(ctx, convert.ToModelTaskConditionFromRequest(c.Msg))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.SearchTasksResponse{
		Tasks: convert.ToConnectTasks(tasks),
	})

	return res, nil
}

func (h taskServiceHandler) CreateTask(ctx context.Context, c *connect.Request[apiv1.CreateTaskRequest]) (*connect.Response[apiv1.CreateTaskResponse], error) {
	if err := h.validator.Validate(c.Msg); err != nil {
		return nil, failure.New(errctrl.InvalidArgument, failure.Message(err.Error()))
	}

	task, err := h.useCase.CreateTask(ctx, convert.ToModelTaskFromRequest(c.Msg))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.CreateTaskResponse{
		Task: convert.ToConnectTask(task),
	})

	return res, nil
}

func (h taskServiceHandler) UpdateTask(ctx context.Context, c *connect.Request[apiv1.UpdateTaskRequest]) (*connect.Response[apiv1.UpdateTaskResponse], error) {
	if err := h.validator.Validate(c.Msg); err != nil {
		return nil, failure.New(errctrl.InvalidArgument, failure.Message(err.Error()))
	}

	task, err := h.useCase.UpdateTask(ctx, convert.ToModelTask(c.Msg.Task))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.UpdateTaskResponse{
		Task: convert.ToConnectTask(task),
	})

	return res, nil
}

func (h taskServiceHandler) DeleteTask(ctx context.Context, c *connect.Request[apiv1.DeleteTaskRequest]) (*connect.Response[emptypb.Empty], error) {
	if err := h.validator.Validate(c.Msg); err != nil {
		return nil, failure.New(errctrl.InvalidArgument, failure.Message(err.Error()))
	}

	if err := h.useCase.DeleteTask(ctx, c.Msg.Id); err != nil {
		return nil, err
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}
