package solve

import (
	"fmt"

	"github.com/klaidliadon/chess"
)

type state struct {
	previous    *state
	squares     []chess.Position
	piecesIndex int
	squareIndex int
}

func (s state) String() string {
	return fmt.Sprintf("Piece %d, step %d/%d", s.piecesIndex, s.squareIndex, len(s.squares))
}

func (s *state) validIndex() bool         { return s.squareIndex < len(s.squares) }
func (s *state) position() chess.Position { return s.squares[s.squareIndex] }
func (s *state) isFirst() bool            { return s.piecesIndex == 0 }
