package models

import "time"

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
	ProblemID    ID
	CategoryID   ID
	TaskNumber   int // 1-11
	ProblemImage string
	Parts        []string
	Answer       string
}

type Submission struct {
	ChatID    ID
	ProblemID ID
	Result    Result
	CreatedAt time.Time
	UpdatedAt time.Time
}
