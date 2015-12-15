package solve

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/klaidliadon/chess"
)

func NewCheckmate(w, h int, pieces chess.PieceCount) *Checkmate {
	return &Checkmate{
		Board:  chess.NewBoard(w, h),
		Pieces: pieces.List(),
	}
}

// Checkmate solves the extended n queen problem (any piece excluding pawns, any board size)
type Checkmate struct {
	*state
	Board  *chess.Board
	Pieces []chess.Piece
	Length int
	result chan []chess.Placement
}

func (c *Checkmate) Solve(debug bool) chan []chess.Placement {
	c.state = &state{squares: c.Board.Positions()}
	c.result = make(chan []chess.Placement)
	reader := bufio.NewReader(os.Stdin)
	go func() {
		defer close(c.result)
		for c.Next() {
			if debug {
				fmt.Println(c)
				reader.ReadString('\n')
			}
		}
	}()
	return c.result
}

func (c *Checkmate) String() string {
	r := c.Board.String()
	if c.isComplete() {
		r = color.GreenString(r)
	}
	return r
}

func (s *Checkmate) isComplete() bool       { return s.piecesIndex == len(s.Pieces) }
func (s *Checkmate) currPiece() chess.Piece { return s.Pieces[s.piecesIndex] }
func (s *Checkmate) prevPiece() chess.Piece { return s.Pieces[s.piecesIndex-1] }
func (s *Checkmate) nextPiece() chess.Piece { return s.Pieces[s.piecesIndex+1] }

func (s *Checkmate) RemoveSquare(p chess.Position) {
	var index = -1
	for i, v := range s.squares {
		if v != p {
			continue
		}
		index = i
		break
	}
	if index == -1 {
		return
	}
	s.squares = append(s.squares[:index], s.squares[index+1:]...)
}

func (s *Checkmate) invalid() bool {
	pos := s.position()
	return !s.Board.Free(pos) || s.Board.IsSafe(chess.Placement{
		Piece: s.currPiece(), Position: pos,
	})
}

func (s *Checkmate) Next() bool {
	if s.isComplete() {
		s.result <- s.Board.Combination()
		return s.Previous()
	}
	if !s.isFirst() && s.currPiece() == s.prevPiece() {
		prev := s.previous.position()
		for s.validIndex() && s.position().Before(prev) {
			s.squareIndex++
		}
	}
	for s.squareIndex < len(s.squares) {
		p := chess.Placement{Position: s.position(), Piece: s.currPiece()}
		if s.invalid() {
			s.squareIndex++
			continue
		}
		s.Board.Place(p)
		safe, _ := p.Split(s.squares)
		s.state = &state{
			previous:    s.state,
			squares:     safe,
			piecesIndex: s.piecesIndex + 1,
		}
		return true
	}
	return s.Previous()
}

func (s *Checkmate) Previous() bool {
	if s.previous == nil {
		return false
	}
	s.Board.RemoveLast()
	s.previous.squareIndex++
	s.state = s.previous
	return true
}
