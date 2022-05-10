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
	ProblemID    ID
	CategoryID   ID
	ProblemImage string
	Parts        []string
	Answer       string
}

type Category struct {
	CategoryID ID
	TaskNumber int
	Title      string
}

type Submission struct {
	SubmissionID ID
	ChatID       ID
	ProblemID    ID
	Result       Result
	// CreatedAt    time.Time
	// UpdatedAt    time.Time
}

// TODO: add hashed images
