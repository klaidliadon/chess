package checkmate

var (
	errEmpty   = "cannot remove last piece: empty board"
	errInvalid = "cannot place piece: invalid position"
)

// Size is the dimensions of a board
type Size struct {
	Width  int
	Height int
}

// Contains check if the position fits the size
func (s Size) Contains(pos Position) bool {
	return !(pos.X < 0 || pos.Y < 0 || pos.X >= s.Width || pos.Y >= s.Height)
}

// NewBoard creates a board of the given dimensions
func NewBoard(w, h int) Board {
	return Board{
		Size:    Size{w, h},
		Squares: make(map[Position]Piece, w*h),
	}
}

// Board represents a playing board
type Board struct {
	Size
	Squares    map[Position]Piece
	placements positionList
}

// Place puts a piece in a square
func (b *Board) Place(piece Piece, pos Position) {
	if !b.Contains(pos) {
		panic(errInvalid)
	}
	b.Squares[pos] = piece
	b.placements.Push(pos)
}

// RemoveLast pops the last piece placed, panics if there are no pieces
func (b *Board) RemoveLast() (piece Piece, pos Position) {
	if len(b.placements) == 0 {
		panic(errEmpty)
	}
	pos = b.placements.Pop()
	piece = b.Squares[pos]
	delete(b.Squares, pos)
	return
}

func (b *Board) Combination() []Placement {
	var result = make([]Placement, len(b.placements))
	for i, pos := range b.placements {
		result[i] = Placement{Piece: b.Squares[pos], Position: pos}
	}
	return result
}
