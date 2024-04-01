package model

import "time"

type TaskStatus string

const (
	TaskStatusUnspecified TaskStatus = "Unspecified"
	TaskStatusOpen        TaskStatus = "Open"
	TaskStatusInProgress  TaskStatus = "InProgress"
	TaskStatusDone        TaskStatus = "Done"
)

type TaskColumn string

const (
	TaskColumnTitle      TaskColumn = "title"
	TaskColumnAuthorID   TaskColumn = "author_id"
	TaskColumnAssigneeID TaskColumn = "assignee_id"
	TaskColumnStatus     TaskColumn = "status"
)

type Task struct {
	ID         int64      `gorm:"primaryKey;autoIncrement:true;column:id;"`
	Title      string     `gorm:"column:title;"`
	AuthorID   int64      `gorm:"column:author_id;"`
	Author     *User      `gorm:"foreignKey:AuthorID"`
	AssigneeID int64      `gorm:"column:assignee_id;"`
	Assignee   *User      `gorm:"foreignKey:AssigneeID"`
	Status     TaskStatus `gorm:"column:status;"`
	DueDate    time.Time  `gorm:"column:due_date;"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
}
