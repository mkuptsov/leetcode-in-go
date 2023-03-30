package problem58

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct{
		input string
		want int
	}{
		{
			input: "Hello World",
			want: 5,
		},
		{
			input: "   fly me   to   the moon  ",
			want: 4,
		},
		{
			input: "luffy is still joyboy",
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			l := lengthOfLastWord(test.input)

			require.Equal(t, test.want, l)
		})
	}
}
