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

func PieceList(count map[Piece]int) []Piece {
	var ps = make([]Piece, 0, len(count))
	for p, n := range count {
		for i := 0; i < n; i++ {
			ps = append(ps, p)
		}
	}
	return ps
}
