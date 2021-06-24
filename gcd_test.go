package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		name string
		a    int64
		b    int64
		want int64
	}{
		{
			name: "0 0",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "18 0",
			a:    18,
			b:    0,
			want: 18,
		},
		{
			name: "0 1",
			a:    0,
			b:    1,
			want: 1,
		},
		{
			name: "2 4",
			a:    2,
			b:    4,
			want: 2,
		},
		{
			name: "4 2",
			a:    4,
			b:    2,
			want: 2,
		},
		{
			name: "8 12",
			a:    8,
			b:    12,
			want: 4,
		},
		{
			name: "7 13",
			a:    7,
			b:    13,
			want: 1,
		},
		{
			name: "-8 12",
			a:    -8,
			b:    12,
			want: 4,
		},
		{
			name: "8 -12",
			a:    8,
			b:    -12,
			want: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GCD(tc.a, tc.b)
			assert.Equal(t, tc.want, got)
		})
	}
}
