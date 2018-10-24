package game

import (
	"fmt"
)

// Player is the user playing the game
type Player struct {
	name string
}

// NewPlayer creates a new player validating the name
func NewPlayer(name string) (*Player, error) {
	if name == "" {
		return nil, fmt.Errorf("The player name cannot be empty")
	}

	return &Player{name}, nil
}

// GetName returns the name of the player
func (p *Player) GetName() string {
	return p.name
}
