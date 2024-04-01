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

type userServiceHandler struct {
	useCase   usecase.UserUseCase
	validator *protovalidate.Validator
}

var _ apiv1connect.UserServiceHandler = (*userServiceHandler)(nil)

func NewUserServiceHandler(u usecase.UserUseCase, v *protovalidate.Validator) apiv1connect.UserServiceHandler {
	return &userServiceHandler{
		useCase:   u,
		validator: v,
	}
}

func (h userServiceHandler) GetUser(ctx context.Context, c *connect.Request[apiv1.GetUserRequest]) (*connect.Response[apiv1.GetUserResponse], error) {
	if err := h.validator.Validate(c.Msg); err != nil {
		return nil, failure.New(errctrl.InvalidArgument, failure.Message(err.Error()))
	}

	user, err := h.useCase.GetUser(ctx, c.Msg.Id)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetUserResponse{
		User: convert.ToConnectUser(user),
	})

	return res, nil
}

func (h userServiceHandler) GetUsers(ctx context.Context, c *connect.Request[emptypb.Empty]) (*connect.Response[apiv1.GetUsersResponse], error) {
	users, err := h.useCase.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetUsersResponse{
		Users: convert.ToConnectUsers(users),
	})

	return res, nil
}
