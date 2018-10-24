package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetX(t *testing.T) {
	x := 5
	p := NewPosition(x, 100)

	assert.Equal(t, x, p.GetX())
}

func TestGetY(t *testing.T) {
	y := 5
	p := NewPosition(100, y)

	assert.Equal(t, y, p.GetY())
}
