package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCM(t *testing.T) {
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
			name: "0 18",
			a:    0,
			b:    18,
			want: 0,
		},
		{
			name: "18 0",
			a:    18,
			b:    0,
			want: 0,
		},
		{
			name: "6 8",
			a:    6,
			b:    8,
			want: 24,
		},
		{
			name: "-6 8",
			a:    -6,
			b:    8,
			want: 24,
		},
		{
			name: "6 -8",
			a:    6,
			b:    -8,
			want: 24,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LCM(tc.a, tc.b)
			assert.Equal(t, tc.want, got)
		})
	}
}
