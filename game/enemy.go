package game

import (
	"time"
)

// Enemy represents the enemy IA of the game
type Enemy interface {
	GetName() string
	StartMoving() <-chan time.Time
	StopMoving()
	GetCurrentPosition() *Position
	Move(b *Board) error
	IsCurrentPosition(p *Position) bool
}
