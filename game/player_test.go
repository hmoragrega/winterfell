package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAUser(t *testing.T) {
	name := "foo"
	p, err := NewPlayer(name)

	assert.Nil(t, err)
	assert.Equal(t, name, p.GetName())
}

func TestCreateAUserWithAnEmptyName(t *testing.T) {
	name := ""
	_, err := NewPlayer(name)

	assert.NotNil(t, err)
}
