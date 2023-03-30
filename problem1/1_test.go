package problem1

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	tests := []struct{
		name string
		nums []int
		target int
		want []int
	}{
		{
			name: "length 4",
			nums: []int{2,7,11,15},
			target: 9,
			want: []int{0, 1},
		},
		{
			name: "length 3",
			nums: []int{3,2,4},
			target: 6,
			want: []int{1,2},
		},
		{
			name: "length 2",
			nums: []int{3,3},
			target: 6,
			want: []int{0,1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := TwoSum(test.nums, test.target)

			require.Equal(t, test.want, l)
		})
	}
}