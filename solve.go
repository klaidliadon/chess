package checkmate

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type commonState struct {
	Board  *Board
	Pieces []Piece
	Length int
	Result chan []Placement
}

type State struct {
	common      commonState
	previous    *State
	squares     []Position
	piecesIndex int
	squareIndex int
}

func (s State) String() string {
	r := s.common.Board.String()
	if s.IsComplete() {
		r = color.GreenString(r)
	}
	return r
}

func (s *State) ValidIndex() bool         { return s.squareIndex < len(s.squares) }
func (s *State) IsFirst() bool            { return s.piecesIndex == 0 }
func (s *State) IsLast() bool             { return s.piecesIndex+1 >= s.common.Length }
func (s *State) IsComplete() bool         { return s.piecesIndex == s.common.Length }
func (s *State) Piece() Piece             { return s.common.Pieces[s.piecesIndex] }
func (s *State) PrevPiece() Piece         { return s.common.Pieces[s.piecesIndex-1] }
func (s *State) NextPiece() Piece         { return s.common.Pieces[s.piecesIndex+1] }
func (s *State) Position() Position       { return s.squares[s.squareIndex] }
func (s *State) Placement() Placement     { return Placement{Piece: s.Piece(), Position: s.Position()} }
func (s *State) Combination() []Placement { return s.common.Board.Combination() }

func (s *State) RemoveSquare(p Position) {
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

func (s *State) Invalid() bool {
	return !s.common.Board.Free(s.Position()) ||
		s.Placement().Menaces(s.common.Board.placements...)
}

func (s *State) AddCurrent() {
	s.common.Board.Place(s.Placement())
}

func (s *State) RemoveLast() Placement {
	return s.common.Board.RemoveLast()
}

func (s *State) Next() *State {
	if s.IsComplete() {
		s.common.Result <- s.common.Board.Combination()
		return s.Previous()
	}
	if !s.IsFirst() && s.Piece() == s.PrevPiece() {
		prev := s.previous.Position()
		for s.ValidIndex() && s.Position().Before(prev) {
			s.squareIndex++
		}
	}
	for s.squareIndex < len(s.squares) {
		p := s.Placement()
		if s.Invalid() {
			s.squareIndex++
			continue
		}
		s.AddCurrent()
		safe, _ := p.Split(s.squares)
		return &State{
			common:      s.common,
			previous:    s,
			squares:     safe,
			piecesIndex: s.piecesIndex + 1,
		}
	}
	return s.Previous()
}

func (s *State) Previous() *State {
	if s.previous == nil {
		return nil
	}
	s.RemoveLast()
	s.previous.squareIndex++
	return s.previous
}

//
func Solve(w, h int, count map[Piece]int, debug bool) <-chan []Placement {
	ch := make(chan []Placement)
	reader := bufio.NewReader(os.Stdin)
	go func() {
		defer close(ch)
		l := PieceList(count)
		s := &State{common: commonState{
			Board: NewBoard(w, h), Pieces: l, Result: ch, Length: len(l),
		}}
		s.squares = s.common.Board.Positions()
		for s != nil {
			s = s.Next()
			if debug {
				fmt.Println(s)
				reader.ReadString('\n')
			}
		}
	}()
	return ch
}
