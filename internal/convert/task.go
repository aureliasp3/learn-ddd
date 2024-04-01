package convert

import (
	"fmt"
	"time"

	apiv1 "learn-ddd/gen/api/v1"
	"learn-ddd/internal/domain/model"
)

func ToConnectTask(t *model.Task) *apiv1.Task {
	return &apiv1.Task{
		Id:       t.ID,
		Title:    t.Title,
		Author:   ToConnectUser(t.Author),
		Assignee: ToConnectUser(t.Assignee),
		Status:   ToConnectTaskStatus(t.Status),
		DueDate:  ToConnectDueDate(t.DueDate),
	}
}

func ToConnectTasks(ts []*model.Task) []*apiv1.Task {
	ret := make([]*apiv1.Task, len(ts))
	for i, v := range ts {
		vv := ToConnectTask(v)
		ret[i] = vv
	}
	return ret
}

func ToConnectTaskStatus(s model.TaskStatus) apiv1.TaskStatus {
	switch s {
	case model.TaskStatusOpen:
		return apiv1.TaskStatus_TASK_STATUS_OPEN
	case model.TaskStatusInProgress:
		return apiv1.TaskStatus_TASK_STATUS_IN_PROGRESS
	case model.TaskStatusDone:
		return apiv1.TaskStatus_TASK_STATUS_DONE
	default:
		return apiv1.TaskStatus_TASK_STATUS_UNSPECIFIED
	}
}

func ToConnectDueDate(t time.Time) *apiv1.Date {
	fmt.Println("++++++++++++++++++++++")
	fmt.Println(t.Year(), t.Month(), t.Day())
	fmt.Println("++++++++++++++++++++++")
	return &apiv1.Date{
		Year:  int64(t.Year()),
		Month: int64(t.Month()),
		Day:   int64(t.Day()),
	}
}

func ToModelTask(t *apiv1.Task) *model.Task {
	return &model.Task{
		ID:         t.Id,
		Title:      t.Title,
		AuthorID:   t.Author.Id,
		AssigneeID: t.Assignee.Id,
		Status:     ToModelTaskStatus(t.Status),
		DueDate:    ToModelDueData(t.DueDate),
	}
}

func ToModelTaskStatus(s apiv1.TaskStatus) model.TaskStatus {
	switch s {
	case apiv1.TaskStatus_TASK_STATUS_OPEN:
		return model.TaskStatusOpen
	case apiv1.TaskStatus_TASK_STATUS_IN_PROGRESS:
		return model.TaskStatusInProgress
	case apiv1.TaskStatus_TASK_STATUS_DONE:
		return model.TaskStatusDone
	default:
		return model.TaskStatusUnspecified
	}
}

func ToModelDueData(d *apiv1.Date) time.Time {
	return time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.Local)
}

func ToModelTaskFromRequest(req *apiv1.CreateTaskRequest) *model.Task {
	return &model.Task{
		Title:      req.Title,
		AuthorID:   req.AuthorId,
		AssigneeID: req.AssigneeId,
		Status:     ToModelTaskStatus(req.Status),
		DueDate:    ToModelDueData(req.DueDate),
	}
}

func ToModelTaskConditionFromRequest(req *apiv1.SearchTasksRequest) map[model.TaskColumn]any {
	m := make(map[model.TaskColumn]any, 4)

	if req.Title != nil && *req.Title != "" {
		m[model.TaskColumnTitle] = *req.Title
	}
	if req.AuthorId != nil && *req.AuthorId != 0 {
		m[model.TaskColumnAuthorID] = *req.AuthorId
	}
	if req.AssigneeId != nil && *req.AssigneeId != 0 {
		m[model.TaskColumnAssigneeID] = *req.AssigneeId
	}
	if req.Status != nil {
		m[model.TaskColumnStatus] = ToModelTaskStatus(*req.Status)
	}

	return m
}
