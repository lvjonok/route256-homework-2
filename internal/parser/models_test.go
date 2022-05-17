package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem_AddPart(t *testing.T) {
	tests := []struct {
		Problem  *Problem
		Element  string
		Expected *Problem
	}{
		{
			Problem:  &Problem{Parts: []string{"part1"}},
			Expected: &Problem{Parts: []string{"part1 part2"}},
			Element:  "part2",
		},
		{
			Problem:  &Problem{Parts: []string{"part1"}},
			Expected: &Problem{Parts: []string{"part1", baseURL}},
			Element:  baseURL,
		},
		{
			Problem:  &Problem{Parts: []string{"part1."}},
			Expected: &Problem{Parts: []string{"part1. part2"}},
			Element:  "  part2 ",
		},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test %d", idx), func(t *testing.T) {
			tt.Problem.AddPart(tt.Element)

			require.Equal(t, tt.Problem.Parts, tt.Expected.Parts)
		})
	}
}

func TestProblem_AddAnswer(t *testing.T) {
	tests := []struct {
		Problem  *Problem
		Answer   string
		Expected *Problem
	}{
		{
			Problem:  &Problem{},
			Answer:   "Ответ: 5.",
			Expected: &Problem{Answer: "5"},
		},
		{
			Problem:  &Problem{},
			Answer:   "5.",
			Expected: &Problem{Answer: "5"},
		},
		{
			Problem:  &Problem{},
			Answer:   "5 ",
			Expected: &Problem{Answer: "5"},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test %d", idx), func(t *testing.T) {
			tt.Problem.AddAnswer(tt.Answer)

			require.Equal(t, tt.Problem.Answer, tt.Expected.Answer)
		})
	}

}
