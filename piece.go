package chess

// Piece is a chess piece type
type Piece int

//go:generate stringer -type=Piece

// List of all pieces
const (
	_ Piece = iota
	Knight
	King
	Bishop
	Rook
	Queen
)

var menaces = map[string]canMove{
	"same":       func(x, y int) bool { return x == 0 && y == 0 },
	"adjacent":   func(x, y int) bool { return x < 2 && y < 2 },
	"special":    func(x, y int) bool { return x == 1 && y == 2 || x == 2 && y == 1 },
	"diagonal":   func(x, y int) bool { return x == y },
	"orthogonal": func(x, y int) bool { return x == 0 || y == 0 },
}

func (p Piece) Simbol() string {
	switch p {
	case King:
		return "♚"
	case Queen:
		return "♛"
	case Rook:
		return "♜"
	case Bishop:
		return "♝"
	case Knight:
		return "♞"
	}
	return "?"
}

// Menaces tells if the piece is menacing any of the positions
func (p Piece) Menaces(self Position, other ...Position) bool {
	var moves = []string{"same"}
	switch p {
	case Knight:
		moves = append(moves, "special")
	case King:
		moves = append(moves, "adjacent")
	case Bishop:
		moves = append(moves, "diagonal")
	case Rook:
		moves = append(moves, "orthogonal")
	case Queen:
		moves = append(moves, "orthogonal", "diagonal")
	}
	for _, o := range other {
		x, y := self.Distance(o)
		for _, s := range moves {
			if menaces[s](x, y) {
				return true
			}
		}
	}
	return false
}

// Split divides the positions in menaced, and not menaced
func (p Piece) Split(self Position, others []Position) (safe, unsafe []Position) {
	safe, unsafe = make([]Position, 0, len(others)), make([]Position, 0, len(others))
	for _, o := range others {
		if p.Menaces(self, o) {
			unsafe = append(unsafe, o)
		} else {
			safe = append(safe, o)
		}
	}
	return safe, unsafe
}
