package entity

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string
	Logo string
}

type UserStory struct {
	gorm.Model
	Number              uint
	As                  string
	Want                string
	Because             string
	Order               uint
	Status              string
	Point               uint
	ProjectID           uint
	Project             Project
	Subtasks            []Subtask
	AcceptanceCriterias []AcceptanceCriteria
	Assignees           []UserStoryAssignee
}

type AcceptanceCriteria struct {
	gorm.Model
	Text        string
	Order       uint
	Status      string
	UserStoryID uint
}

type Subtask struct {
	gorm.Model
	Text        string
	Order       uint
	Status      string
	Point       uint
	UserStoryID uint
	Assignees   []SubtaskAssignee
}

type SubtaskAssignee struct {
	gorm.Model
	SubtaskID uint
	Subtask   Subtask
	MemberID  int
	Member    Member
}

type UserStoryAssignee struct {
	gorm.Model
	UserStoryID uint
	UserStory   UserStory
	MemberID    uint
	Member      Member
}

type Member struct {
	gorm.Model
	Name    string
	Profile string
}
