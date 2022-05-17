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
			ProblemID: 77415,
			Expected: &Problem{
				ProblemId:    77415,
				ProblemImage: "",
				Parts:        []string{"Найдите значение выражения", "https://ege.sdamgia.ru/formula/svg/9e/9e5acdd0b0eda0bc92a128a2a1d9a0f8.svg", "если", "https://ege.sdamgia.ru/formula/svg/82/82903cf1c84f5c9788fd410cd677298b.svg"},
				Answer:       "22",
			},
		},
		{
			ProblemID: 27151,
			Expected: &Problem{
				ProblemId:    27151,
				ProblemImage: "https://ege.sdamgia.ru/get_file?id=66984",
				Parts:        []string{"Основанием прямой треугольной призмы служит прямоугольный треугольник с катетами 6 и 8. Площадь ее поверхности равна 288. Найдите высоту призмы."},
				Answer:       "10",
			},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("test %d, idx", idx), func(t *testing.T) {
			got, err := ParseProblem(tt.ProblemID)
			require.NoError(t, err)
			require.Truef(t, reflect.DeepEqual(got, tt.Expected), "problem:%d\nexpected %v\ngot%v", tt.ProblemID, tt.Expected, got)
		})
	}
}
