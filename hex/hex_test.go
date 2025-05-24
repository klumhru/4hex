package hex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition_String(t *testing.T) {
	tests := []struct {
		name string
		pos  Position
		want string
	}{
		{
			name: "Zero position",
			pos:  Position{Q: 0, R: 0},
			want: "Pos(q:0, r:0)",
		},
		{
			name: "Positive coordinates",
			pos:  Position{Q: 1, R: 2},
			want: "Pos(q:1, r:2)",
		},
		{
			name: "Negative coordinates",
			pos:  Position{Q: -1, R: -2},
			want: "Pos(q:-1, r:-2)",
		},
		{
			name: "Mixed coordinates",
			pos:  Position{Q: 5, R: -3},
			want: "Pos(q:5, r:-3)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.want, tt.pos.String())
		})
	}
}

func TestNewPosition(t *testing.T) {
	tests := []struct {
		name string
		q    int
		r    int
		want Position
	}{
		{
			name: "Zero position",
			q:    0,
			r:    0,
			want: Position{Q: 0, R: 0},
		},
		{
			name: "Positive coordinates",
			q:    1,
			r:    2,
			want: Position{Q: 1, R: 2},
		},
		{
			name: "Negative coordinates",
			q:    -1,
			r:    -2,
			want: Position{Q: -1, R: -2},
		},
		{
			name: "Mixed coordinates",
			q:    5,
			r:    -3,
			want: Position{Q: 5, R: -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.want, NewPosition(tt.q, tt.r))
		})
	}
}
