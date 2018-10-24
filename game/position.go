package game

// Position a position in the board
type Position struct {
	x int
	y int
}

// GetX returns the position in the x axis
func (p *Position) GetX() int {
	return p.x
}

// GetY returns the position in the y axis
func (p *Position) GetY() int {
	return p.y
}

// NewPosition creates a new position
func NewPosition(x int, y int) *Position {
	return &Position{
		x: x,
		y: y,
	}
}
