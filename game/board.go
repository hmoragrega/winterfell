package game

// Board represents an rectangle where the gamne is played
type Board struct {
	width  int
	height int
}

// NewBoard creates a new board
func NewBoard(width int, height int) *Board {
	return &Board{
		width:  width,
		height: height,
	}
}

// Checks if the position is out of bounds
func (b *Board) isPositionValid(p *Position) bool {
	return p.x >= 0 && p.x <= b.width && p.y >= 0 && p.y <= b.height
}
