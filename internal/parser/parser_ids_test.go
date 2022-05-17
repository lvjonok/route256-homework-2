package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseProblemsIds(t *testing.T) {
	tests := []struct {
		categoryID int
		problems   []int
	}{
		{
			categoryID: 13,
			problems:   []int{26669, 77376, 77377},
		},
		{
			categoryID: 63,
			problems:   []int{77415, 77416, 77417, 98471},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test %d", idx), func(t *testing.T) {
			got, err := ParseProblemsIds(tt.categoryID)
			require.NoError(t, err)
			require.Truef(t, reflect.DeepEqual(got, tt.problems), "Category=%d\nExpected %v\ngot %v", tt.categoryID, tt.problems, got)
		})
	}
}
