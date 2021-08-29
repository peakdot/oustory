package entity

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string
	Logo string
}

type Backlog struct {
	gorm.Model
	Type                string
	Number              uint
	Text                string
	Order               uint
	Status              string
	Point               uint
	ProjectID           uint
	Project             Project
	Subtasks            []Subtask
	AcceptanceCriterias []AcceptanceCriteria
	Assignees           []BacklogAssignee
}

type AcceptanceCriteria struct {
	gorm.Model
	Text      string
	Order     uint
	Status    string
	BacklogID uint
}

type Subtask struct {
	gorm.Model
	Text      string
	Order     uint
	Status    string
	Point     uint
	BacklogID uint
	Backlog   Backlog
	Assignees []SubtaskAssignee
}

type SubtaskAssignee struct {
	gorm.Model
	SubtaskID uint
	Subtask   Subtask
	MemberID  int
	Member    Member
}

type BacklogAssignee struct {
	gorm.Model
	BacklogID uint
	Backlog   Backlog
	MemberID  uint
	Member    Member
}

type Member struct {
	gorm.Model
	Name    string
	Profile string
}
