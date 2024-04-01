package db

import (
	"context"
	"testing"

	"learn-ddd/lib/ctxctrl"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"learn-ddd/db/dbtest"
	"learn-ddd/internal/domain/model"
	"learn-ddd/internal/fixture"
)

func Test_userRepository_Fetch(t *testing.T) {
	dbtest.Setup()
	opt := cmpopts.IgnoreFields(model.User{}, "CreatedAt", "UpdatedAt")

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    []*model.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
			}()},
			prepare: func() {
				dbtest.TestDB.Create(&fixture.FakeUsers)
			},
			want:    fixture.FakeUsers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewUserRepository()
			got, err := r.Fetch(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Select(t *testing.T) {
	dbtest.Setup()
	opt := cmpopts.IgnoreFields(model.User{}, "CreatedAt", "UpdatedAt")

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    *model.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{ctx: func() (ctx context.Context) {
				return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
			}(),
				id: 1},
			prepare: func() {
				dbtest.TestDB.Create(fixture.FakeUser)
			},
			want:    fixture.FakeUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewUserRepository()
			got, err := r.Select(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("Select() got = %v, want %v", got, tt.want)
			}
		})
	}
}
