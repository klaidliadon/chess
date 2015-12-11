package checkmate

import "fmt"

// Position is a square in the board
type Position struct {
	X, Y int
}

// Distance returns the absolute distance from another position
func (p Position) Distance(o Position) (int, int) {
	dX := p.X - o.X
	if dX < 0 {
		dX = -dX
	}
	dY := p.Y - o.Y
	if dY < 0 {
		dY = -dY
	}
	return dX, dY
}

func (p Position) Before(o Position) bool {
	return p.Y < o.Y || p.Y == o.Y && p.X < o.Y
}

type positionList []Position

func (p *positionList) Push(pos Position) {
	*p = append(*p, pos)
}

func (p *positionList) Pop() Position {
	count := len(*p)
	pos := (*p)[count-1]
	*p = (*p)[:count-1]
	return pos
}

type Placement struct {
	Piece
	Position
}

func (p Placement) Menaces(o Placement) bool {
	return p.Piece.Menaces(p.Position, o.Position)
}

func (p Placement) String() string {
	return fmt.Sprintf("%s %v", p.Piece, p.Position)
}
