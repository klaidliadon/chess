package chess

import (
	"bytes"
	"fmt"
)

// Placement represents a piece in a square
type Placement struct {
	Piece
	Position
}

// Menaces checks if the placement menace the other
func (p Placement) Menaces(o ...Placement) bool {
	return p.Piece.Menaces(p.Position, PlacementStack(o).Positions()...)
}

// Split divides the positions in menaced, and not menaced
func (p Placement) Split(pos []Position) (safe, unsafe []Position) {
	return p.Piece.Split(p.Position, pos)
}

// String returns the format "piece position"
func (p Placement) String() string {
	return fmt.Sprintf("%c {%d,%d}", p.Piece.Rune(), p.Position.X, p.Position.Y)
}

// PlacementStack is a stack of Placements
type PlacementStack []Placement

// Push adds an element to the stack
func (p *PlacementStack) Push(i Placement) {
	*p = append(*p, i)
}

// Pop removes the last element from the stack
func (p *PlacementStack) Pop() Placement {
	count := len(*p)
	i := (*p)[count-1]
	*p = (*p)[:count-1]
	return i
}

// Positions returns the positions of the stack
func (p PlacementStack) Positions() []Position {
	var r = make([]Position, len(p))
	for i, v := range p {
		r[i] = v.Position
	}
	return r
}

func (p PlacementStack) String() string {
	w := bytes.NewBuffer(nil)
	var last Piece
	for i, v := range p {
		if last != v.Piece {
			if i > 0 {
				fmt.Fprint(w, "\n")
			}
			last = v.Piece
			fmt.Fprintf(w, "%c ", last.Rune())
		}
		fmt.Fprintf(w, "%v", v.Position)

	}
	return w.String()
}
