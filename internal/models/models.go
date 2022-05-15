package models

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

type Category struct {
	ID         ID
	CategoryID ID
	TaskNumber int
	Title      string
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
