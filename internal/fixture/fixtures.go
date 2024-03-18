package fixture

import (
	"time"

	"learn-ddd/internal/domain/model"
)

var now = time.Now()

var (
	FakeUser = &model.User{
		ID:        1,
		Name:      "Name1",
		CreatedAt: now,
		UpdatedAt: now,
	}

	FakeUsers = []*model.User{
		{
			ID:        1,
			Name:      "Name1",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        2,
			Name:      "Name2",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
)

var (
	FakeTask = &model.Task{
		ID:         1,
		Title:      "Title1",
		AuthorID:   FakeUser.ID,
		Author:     FakeUser,
		AssigneeID: FakeUser.ID,
		Assignee:   FakeUser,
		Status:     model.TaskStatusOpen,
		DueDate:    now,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	UpdatedFakeTask = &model.Task{
		ID:         1,
		Title:      "Title1",
		AuthorID:   1,
		AssigneeID: 1,
		Status:     model.TaskStatusDone,
		DueDate:    now,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	FakeTasks = []*model.Task{
		{
			ID:         1,
			Title:      "Title1",
			AuthorID:   FakeUsers[0].ID,
			Author:     FakeUsers[0],
			AssigneeID: FakeUsers[1].ID,
			Assignee:   FakeUsers[1],
			Status:     model.TaskStatusOpen,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Title:      "Title2",
			AuthorID:   FakeUsers[1].ID,
			Author:     FakeUsers[1],
			AssigneeID: FakeUsers[0].ID,
			Assignee:   FakeUsers[0],
			Status:     model.TaskStatusDone,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	FakeTaskWithoutUser = &model.Task{
		ID:         1,
		Title:      "Title1",
		AuthorID:   FakeUser.ID,
		AssigneeID: FakeUser.ID,
		Status:     model.TaskStatusOpen,
		DueDate:    now,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	FakeTasksWithoutUser = []*model.Task{
		{
			ID:         1,
			Title:      "Title1",
			AuthorID:   FakeUsers[0].ID,
			AssigneeID: FakeUsers[1].ID,
			Status:     model.TaskStatusOpen,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Title:      "Title2",
			AuthorID:   FakeUsers[1].ID,
			AssigneeID: FakeUsers[0].ID,
			Status:     model.TaskStatusDone,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	FakeTasksForSearch = []*model.Task{
		{
			ID:         1,
			Title:      "Title1",
			AuthorID:   FakeUsers[0].ID,
			Author:     FakeUsers[0],
			AssigneeID: FakeUsers[1].ID,
			Assignee:   FakeUsers[1],
			Status:     model.TaskStatusOpen,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Title:      "Title2",
			AuthorID:   FakeUsers[1].ID,
			Author:     FakeUsers[1],
			AssigneeID: FakeUsers[0].ID,
			Assignee:   FakeUsers[0],
			Status:     model.TaskStatusDone,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         3,
			Title:      "Title3",
			AuthorID:   FakeUsers[0].ID,
			Author:     FakeUsers[0],
			AssigneeID: FakeUsers[1].ID,
			Assignee:   FakeUsers[1],
			Status:     model.TaskStatusOpen,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Title:      "Title4",
			AuthorID:   FakeUsers[1].ID,
			Author:     FakeUsers[1],
			AssigneeID: FakeUsers[0].ID,
			Assignee:   FakeUsers[0],
			Status:     model.TaskStatusDone,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}

	FakeTasksForSearchWithoutUser = []*model.Task{
		{
			ID:         1,
			Title:      "Title1",
			AuthorID:   FakeUsers[0].ID,
			AssigneeID: FakeUsers[1].ID,
			Status:     model.TaskStatusOpen,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         2,
			Title:      "Title2",
			AuthorID:   FakeUsers[1].ID,
			AssigneeID: FakeUsers[0].ID,
			Status:     model.TaskStatusDone,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         3,
			Title:      "Title3",
			AuthorID:   FakeUsers[0].ID,
			AssigneeID: FakeUsers[1].ID,
			Status:     model.TaskStatusOpen,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
		{
			ID:         4,
			Title:      "Title4",
			AuthorID:   FakeUsers[1].ID,
			AssigneeID: FakeUsers[0].ID,
			Status:     model.TaskStatusDone,
			DueDate:    now,
			CreatedAt:  now,
			UpdatedAt:  now,
		},
	}
)
