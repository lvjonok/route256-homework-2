package models

import (
	"bytes"
)

type ID int

// Corresponds to submission result
type Result string

const (
	Correct Result = "correct"
	Wrong   Result = "wrong"
	Pending Result = "pending"
	Aborted Result = "aborted"
)

type Problem struct {
	ID           ID
	ProblemID    ID // id from reshuege.ru
	CategoryID   ID // foreign key from Category.ID
	ProblemImage string
	Parts        []string
	Answer       string
}

func (p *Problem) Equal(another *Problem) bool {
	equalParts := true
	for idx := 0; idx < len(p.Parts); idx++ {
		if len(p.Parts) <= idx || p.Parts[idx] != another.Parts[idx] {
			equalParts = false
		}
	}
	if len(p.Parts) != len(another.Parts) {
		equalParts = false
	}

	return p.CategoryID == another.CategoryID &&
		equalParts &&
		p.ProblemID == another.ProblemID &&
		p.Answer == another.Answer &&
		p.ProblemImage == another.ProblemImage
}

type Category struct {
	ID         ID
	CategoryID ID
	TaskNumber int
	Title      string
}

func (c *Category) Equal(another *Category) bool {
	return c.TaskNumber == another.TaskNumber &&
		c.Title == another.Title &&
		c.CategoryID == another.CategoryID
}

type Submission struct {
	SubmissionID ID
	ChatID       ID
	ProblemID    ID // foreign key from Problem.ID
	Result       Result
}

type TaskStat struct {
	TaskNumber int
	Correct    int
	All        int
}

type Statistics struct {
	Stat []TaskStat
}

type Rating struct {
	Position int
	All      int
}

type Image struct {
	ID      ID
	Content []byte
	Href    string
}

func (i *Image) Equal(another *Image) bool {
	return bytes.Equal(i.Content, another.Content) && i.Href == another.Href
}
