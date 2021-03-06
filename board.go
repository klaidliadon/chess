package chess

import (
	"bytes"
	"fmt"
)

var (
	errEmpty   = "cannot remove last piece: empty board"
	errInvalid = "cannot place piece: invalid position"
)

// Size is the dimensions of a board
type Size struct {
	Width  int
	Height int
}

// SquareCount returns the amount of squares
func (s Size) SquareCount() int {
	return s.Width * s.Height
}

// Contains check if the position fits the size
func (s Size) Contains(pos Position) bool {
	return !(pos.X < 0 || pos.Y < 0 || pos.X >= s.Width || pos.Y >= s.Height)
}

// NewBoard creates a board of the given dimensions
func NewBoard(w, h int) *Board {
	return &Board{Size: Size{w, h}}
}

// Board represents a playing board
type Board struct {
	Size
	placements PlacementStack
}

func (b *Board) String() string {
	var cache = make(map[Position]rune, b.SquareCount()-len(b.placements))
	for _, p := range b.placements {
		cache[p.Position] = p.Piece.Rune()
	}

	w := bytes.NewBuffer(nil)
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			var r = '☐'
			if v, ok := cache[Position{x, y}]; ok {
				r = v
			}
			fmt.Fprintf(w, " %c", r)
		}
		if y != b.Height-1 {
			w.WriteRune('\n')
		}
	}
	s := b.placements.String()
	if s != "" {
		w.WriteRune('\n')
	}
	w.WriteString(s)
	return w.String()
}

// Positions returns all the positions of the board
func (b *Board) Positions() []Position {
	p := make([]Position, 0, b.SquareCount())
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			p = append(p, Position{x, y})
		}
	}
	return p
}

// Free tells if a square is occupied
func (b *Board) Free(pos Position) bool {
	for _, p := range b.placements {
		if p.Position == pos {
			return false
		}
	}
	return true
}

// Free tells if a square is occupied
func (b *Board) IsSafe(p Placement) bool {
	return p.Menaces(b.placements...)
}

// Place puts a piece in a square
func (b *Board) Place(p Placement) {
	if !b.Contains(p.Position) {
		panic(errInvalid)
	}
	b.placements.Push(p)
}

// RemoveLast pops the last piece placed, panics if there are no pieces
func (b *Board) RemoveLast() Placement {
	if len(b.placements) == 0 {
		panic(errEmpty)
	}
	return b.placements.Pop()
}

// Combination returns the current placements
func (b *Board) Combination() []Placement {
	var r = make([]Placement, len(b.placements))
	for i, p := range b.placements {
		r[i] = p
	}
	return r
}
