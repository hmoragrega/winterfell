package game

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Zombie a type of enemy the always moves forward!
type Zombie struct {
	name                string
	position            *Position
	previous            *Position
	milliSecondsPerCell int
	ticker              *time.Ticker
	sync.Mutex
}

// NewZombie creates a new zombie
func NewZombie(name string, position *Position, milliSecondsPerCell int) *Zombie {
	return &Zombie{
		name:                name,
		position:            position,
		milliSecondsPerCell: milliSecondsPerCell,
	}
}

// GetName returns the name of the zombie
func (z *Zombie) GetName() string {
	return z.name
}

// GetCurrentPosition returns the current position of the zombie in the board
func (z *Zombie) GetCurrentPosition() *Position {
	return z.position
}

// StartMoving initiates the zombie movement
func (z *Zombie) StartMoving() <-chan time.Time {
	z.ticker = time.NewTicker(time.Duration(z.milliSecondsPerCell) * time.Millisecond)

	return z.ticker.C
}

// StopMoving stops the zombie movement
func (z *Zombie) StopMoving() {
	if z.ticker != nil {
		log.Println("Stop moving zombie")
		z.ticker.Stop()
		z.ticker = nil
	}
}

// IsCurrentPosition checks if the zombie is in the given position
func (z *Zombie) IsCurrentPosition(p *Position) bool {
	z.Lock()
	defer z.Unlock()

	return p != nil && *p == *z.position
}

// Move tries to move the zombie to a valid new location in the board
func (z *Zombie) Move(b *Board) error {
	z.Lock()
	defer z.Unlock()

	positions := z.getPosiblePositions()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(positions), func(i, j int) {
		positions[i], positions[j] = positions[j], positions[i]
	})

	for _, p := range positions {
		if z.isPreviousPosition(p) {
			continue
		}

		if !b.isPositionValid(p) {
			continue
		}

		z.previous = z.position
		z.position = p
		return nil
	}

	return fmt.Errorf("No valid movements have been found for the enemy")
}

// GetPosiblePositions A zombie will
//  - move always
//  - move only forward (right) or up/down
func (z *Zombie) getPosiblePositions() []*Position {
	return []*Position{
		&Position{z.position.x, z.position.y + cellMovement}, // Up
		&Position{z.position.x, z.position.y - cellMovement}, // Down
		&Position{z.position.x + cellMovement, z.position.y}, // Forward
	}
}

func (z *Zombie) isPreviousPosition(p *Position) bool {
	return p != nil && z.previous != nil && *p == *z.previous
}
