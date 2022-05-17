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
	}{}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test %d", idx), func(t *testing.T) {
			got, err := ParseProblemsIds(tt.categoryID)
			require.NoError(t, err)
			require.True(t, reflect.DeepEqual(got, tt.problems))
		})
	}
}
