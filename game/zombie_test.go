package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetZombieName(t *testing.T) {
	name := "foo"
	z := NewZombie(name, nil, 1)

	assert.Equal(t, name, z.GetName())
}

func TestGetCurrentPosition(t *testing.T) {
	p := NewPosition(0, 0)
	z := NewZombie("foo", p, 1)

	assert.Equal(t, p, z.GetCurrentPosition())
}

func TestMoving(t *testing.T) {
	z := NewZombie("foo", NewPosition(0, 0), 1)

	z.StartMoving()
	assert.NotNil(t, z.ticker)

	z.StopMoving()
	assert.Nil(t, z.ticker)
}

func TestMoveIfSpaceIsAvailable(t *testing.T) {
	b := NewBoard(5, 5)
	p := NewPosition(0, 0)
	z := NewZombie("foo", p, 1)

	z.Move(b)

	assert.NotEqual(t, p, z.GetCurrentPosition())
	assert.Equal(t, p, z.previous)
}

func TestDoNotMoveIfSpaceIsNotAvailable(t *testing.T) {
	b := NewBoard(0, 0)
	p := NewPosition(0, 0)
	z := NewZombie("foo", p, 1)

	z.Move(b)

	assert.Equal(t, p, z.GetCurrentPosition())
	assert.Nil(t, z.previous)
}

func TestIsCurrentPosition(t *testing.T) {
	z := NewZombie("foo", nil, 1)

	var tests = []struct {
		current  *Position
		compare  *Position
		expected bool
	}{
		{NewPosition(0, 0), nil, false},
		{NewPosition(1, 1), NewPosition(1, 0), false},
		{NewPosition(1, 1), NewPosition(0, 1), false},
		{NewPosition(1, 1), NewPosition(1, 1), true},
	}

	for _, tt := range tests {
		z.position = tt.current
		assert.Equal(t, tt.expected, z.IsCurrentPosition(tt.compare))
	}
}

func TestGetPosiblePositions(t *testing.T) {
	p := NewPosition(5, 5)
	z := NewZombie("foo", p, 1)

	expected := []*Position{
		{5, 6},
		{5, 4},
		{6, 5},
	}

	assert.Equal(t, expected, z.getPosiblePositions())
}

func TestIsPreviousPosition(t *testing.T) {
	z := NewZombie("foo", NewPosition(0, 0), 1)
	var tests = []struct {
		previous *Position
		compare  *Position
		expected bool
	}{
		{NewPosition(1, 1), NewPosition(1, 0), false},
		{NewPosition(1, 1), NewPosition(0, 1), false},
		{NewPosition(1, 1), NewPosition(1, 1), true},
		{nil, nil, false},
		{NewPosition(0, 0), nil, false},
		{nil, NewPosition(0, 0), false},
	}

	for _, tt := range tests {
		z.previous = tt.previous
		assert.Equal(t, tt.expected, z.isPreviousPosition(tt.compare))
	}
}
