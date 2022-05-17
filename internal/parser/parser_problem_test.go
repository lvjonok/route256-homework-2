package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseProblem(t *testing.T) {
	tests := []struct {
		ProblemID int
		Expected  *Problem
	}{
		{
			ProblemID: 100,
			Expected:  &Problem{},
		},
		{
			ProblemID: 200,
			Expected:  &Problem{},
		},
		{
			ProblemID: 300,
			Expected:  &Problem{},
		},
		{
			ProblemID: 27151,
			Expected:  &Problem{},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test %d, idx", idx), func(t *testing.T) {
			got, err := ParseProblem(tt.ProblemID)
			require.NoError(t, err)
			require.True(t, reflect.DeepEqual(got, tt.Expected))
		})
	}
}
