package usecase

import (
	"context"
	"reflect"
	"testing"

	"learn-ddd/lib/ctxctrl"

	"go.uber.org/mock/gomock"

	"learn-ddd/db"
	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/domain/repository/mock"
	"learn-ddd/internal/fixture"
)

func Test_userUsecase_GetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		want    *model.User
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
			want:    fixture.FakeUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mock.NewMockUserRepository(ctrl)
			mockUserRepo.EXPECT().Select(tt.args.ctx, tt.args.id).Return(fixture.FakeUser, nil)

			u := NewUserUseCase(mockUserRepo)

			got, err := u.GetUser(tt.args.ctx, tt.args.id)
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

func Test_userUsecase_GetUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    []*model.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), db.DB)
				}(),
			},
			want:    fixture.FakeUsers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mock.NewMockUserRepository(ctrl)
			mockUserRepo.EXPECT().Fetch(tt.args.ctx).Return(fixture.FakeUsers, nil)

			u := NewUserUseCase(mockUserRepo)

			got, err := u.GetUsers(tt.args.ctx)
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
