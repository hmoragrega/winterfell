package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardIsPositionValid(t *testing.T) {
	b := NewBoard(5, 5)

	var tests = []struct {
		x        int
		y        int
		expected bool
	}{
		{0, 0, true},
		{3, 3, true},
		{5, 5, true},
		{-1, 0, false},
		{0, -1, false},
		{6, 5, false},
		{5, 6, false},
	}

	for _, tt := range tests {
		p := NewPosition(tt.x, tt.y)
		assert.Equal(t, tt.expected, b.isPositionValid(p), "isPositionValid should be %v for %v", tt.expected, p)
	}
}
