package checkmate

// Piece is a chess piece type
type Piece int

//go:generate stringer -type=Piece

// List of all pieces
const (
	_ Piece = iota
	King
	Knight
	Bishop
	Rook
	Queen
)

var menaces = map[string]func(int, int) bool{
	"same":       func(dX, dY int) bool { return dX == 0 && dY == 0 },
	"adjacent":   func(dX, dY int) bool { return dX < 2 && dY < 2 },
	"special":    func(dX, dY int) bool { return dX == 1 && dY == 2 || dX == 2 && dY == 1 },
	"diagonal":   func(dX, dY int) bool { return dX == dY },
	"orthogonal": func(dX, dY int) bool { return dX == 0 || dY == 0 },
}

// Menaces tells if the piece is menacing a position
func (p Piece) Menaces(self, other Position) bool {
	dX, dY := self.Distance(other)
	if menaces["same"](dX, dY) {
		return true
	}
	switch p {
	case Knight:
		return menaces["special"](dX, dY)
	case King:
		return menaces["adjacent"](dX, dY)
	case Bishop:
		return menaces["diagonal"](dX, dY)
	case Rook:
		return menaces["orthogonal"](dX, dY)
	case Queen:
		return menaces["orthogonal"](dX, dY) || menaces["diagonal"](dX, dY)
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
	if len(safe) == 0 {
		safe = nil
	}
	if len(unsafe) == 0 {
		unsafe = nil
	}
	return safe, unsafe
}
