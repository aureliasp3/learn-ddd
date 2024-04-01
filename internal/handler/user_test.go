package handler

import (
	"context"
	"reflect"
	"testing"

	"learn-ddd/lib/ctxctrl"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"

	"learn-ddd/db"
	apiv1 "learn-ddd/gen/api/v1"
	"learn-ddd/internal/convert"
	"learn-ddd/internal/fixture"
	"learn-ddd/internal/usecase/mock"
	"learn-ddd/lib/errctrl"
)

func Test_userServiceHandler_GetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[apiv1.GetUserRequest]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[apiv1.GetUserResponse]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[apiv1.GetUserRequest]{
					Msg: &apiv1.GetUserRequest{Id: 1},
				},
			},
			want:    connect.NewResponse(&apiv1.GetUserResponse{User: convert.ToConnectUser(fixture.FakeUser)}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserUsecase := mock.NewMockUserUseCase(ctrl)
			mockUserUsecase.EXPECT().GetUser(tt.args.ctx, tt.args.c.Msg.Id).Return(fixture.FakeUser, nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewUserServiceHandler(mockUserUsecase, validator)

			got, err := h.GetUser(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userServiceHandler_GetUsers(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *connect.Request[emptypb.Empty]
	}

	tests := []struct {
		name    string
		args    args
		want    *connect.Response[apiv1.GetUsersResponse]
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), db.DB)
			}(),
				c: &connect.Request[emptypb.Empty]{},
			},
			want:    connect.NewResponse(&apiv1.GetUsersResponse{Users: convert.ToConnectUsers(fixture.FakeUsers)}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserUsecase := mock.NewMockUserUseCase(ctrl)
			mockUserUsecase.EXPECT().GetUsers(tt.args.ctx).Return(fixture.FakeUsers, nil)

			validator := errctrl.Must(protovalidate.New())

			h := NewUserServiceHandler(mockUserUsecase, validator)

			got, err := h.GetUsers(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
