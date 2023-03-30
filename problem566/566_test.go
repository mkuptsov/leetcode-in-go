package problem566

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestMatrixReshape(t *testing.T) {
	tests := []struct{
		name string
		mat [][]int
		r, c int
		want [][]int
	}{
		{
			name: "2x2 to 1x4",
			mat: [][]int{{1, 2}, {3, 4}},
			r: 1,
			c: 4,
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "2x2 to 2x4",
			mat: [][]int{{1, 2}, {3, 4}},
			r: 2,
			c: 4,
			want: [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "2x2 to 4x1",
			mat: [][]int{{1, 2}, {3, 4}},
			r: 4,
			c: 1,
			want: [][]int{{1}, {2}, {3}, {4}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := matrixReshape(test.mat, test.r, test.c)

			require.Equal(t, test.want, l)
		})
	}
}
