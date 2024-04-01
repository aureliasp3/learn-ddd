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

// TODO: DBの中身までチェックする
func Test_taskRepository_Delete(t *testing.T) {
	dbtest.Setup()

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    []*model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
				}(),
				id: 1,
			},
			prepare: func() {
				dbtest.TestDB.Create(fixture.FakeUser)
				dbtest.TestDB.Create(fixture.FakeTask)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewTaskRepository()
			err := r.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_taskRepository_Fetch(t *testing.T) {
	dbtest.Setup()
	opt := cmpopts.IgnoreFields(model.Task{}, "CreatedAt", "UpdatedAt")

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    []*model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
				}()},
			prepare: func() {
				dbtest.TestDB.Create(&fixture.FakeUsers)
				dbtest.TestDB.Create(&fixture.FakeTasksWithoutUser)
			},
			want:    fixture.FakeTasksWithoutUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewTaskRepository()
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

// TODO: DBの中身までチェックする
func Test_taskRepository_Insert(t *testing.T) {
	dbtest.Setup()
	opt := cmpopts.IgnoreFields(model.Task{}, "CreatedAt", "UpdatedAt")

	type args struct {
		ctx  context.Context
		task *model.Task
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    *model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
				}(),
				task: fixture.FakeTaskWithoutUser,
			},
			prepare: func() {
				dbtest.TestDB.Create(fixture.FakeUser)
			},
			want:    fixture.FakeTaskWithoutUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewTaskRepository()
			got, err := r.Insert(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskRepository_Search(t *testing.T) {
	dbtest.Setup()
	opt := cmpopts.IgnoreFields(model.Task{}, "CreatedAt", "UpdatedAt")

	type args struct {
		ctx  context.Context
		cond map[model.TaskColumn]any
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    []*model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
				}(),
				cond: map[model.TaskColumn]any{
					model.TaskColumnTitle:      "2",
					model.TaskColumnAuthorID:   fixture.FakeTasksForSearchWithoutUser[1].AuthorID,
					model.TaskColumnAssigneeID: fixture.FakeTasksForSearchWithoutUser[1].AssigneeID,
					model.TaskColumnStatus:     fixture.FakeTasksForSearchWithoutUser[1].Status,
				},
			},
			prepare: func() {
				dbtest.TestDB.Create(&fixture.FakeUsers)
				dbtest.TestDB.Create(&fixture.FakeTasksForSearchWithoutUser)
			},
			want:    []*model.Task{fixture.FakeTasksForSearchWithoutUser[1]},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewTaskRepository()
			got, err := r.Search(tt.args.ctx, tt.args.cond)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskRepository_Select(t *testing.T) {
	dbtest.Setup()
	opt := cmpopts.IgnoreFields(model.Task{}, "CreatedAt", "UpdatedAt")

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    *model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
				}(),
				id: 1,
			},
			prepare: func() {
				dbtest.TestDB.Create(fixture.FakeUser)
				dbtest.TestDB.Create(fixture.FakeTaskWithoutUser)
			},
			want:    fixture.FakeTaskWithoutUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewTaskRepository()
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

// TODO: DBの中身までチェックする
func Test_taskRepository_Update(t *testing.T) {
	dbtest.Setup()
	opt := cmpopts.IgnoreFields(model.Task{}, "CreatedAt", "UpdatedAt")

	type args struct {
		ctx  context.Context
		task *model.Task
	}

	tests := []struct {
		name    string
		args    args
		prepare func()
		want    *model.Task
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: func() (ctx context.Context) {
					return ctxctrl.SetTx(context.Background(), dbtest.TestDB)
				}(),
				task: fixture.UpdatedFakeTask,
			},
			prepare: func() {
				dbtest.TestDB.Create(fixture.FakeUser)
				dbtest.TestDB.Create(fixture.FakeTask)
			},
			want:    fixture.UpdatedFakeTask,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			r := NewTaskRepository()
			got, err := r.Update(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
