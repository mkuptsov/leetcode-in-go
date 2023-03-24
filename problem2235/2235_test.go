package problem2235

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{
			name: "sum of 4 and 6",
			num1: 4,
			num2: 6,
			want: 11,
		},
		{
			name: "sum of -1 and 6",
			num1: -1,
			num2: 6,
			want: 56,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := Sum(test.num1, test.num2)

			require.Equal(t, test.want, l)
		})
	}
}
