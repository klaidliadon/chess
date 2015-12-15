package chess

import "fmt"

// canMove use the delta in x and y axis and return reachability
type canMove func(int, int) bool

// Position is a square in the board
type Position struct {
	X, Y int
}

func (p Position) String() string {
	return fmt.Sprintf("[%v,%v]", p.X, p.Y)
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
	return p.Y < o.Y || p.Y == o.Y && p.X < o.X
}

type PieceCount map[Piece]int

func (p PieceCount) List() []Piece {
	var ps = make([]Piece, 0, len(p))
	for k, n := range p {
		for i := 0; i < n; i++ {
			ps = append(ps, k)
		}
	}
	return ps
}
