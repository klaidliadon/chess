package solve

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/klaidliadon/chess"
)

type common struct {
	Board  *chess.Board
	Pieces []chess.Piece
	Length int
	Result chan []chess.Placement
}

type State struct {
	common
	previous    *State
	squares     []chess.Position
	piecesIndex int
	squareIndex int
}

func (s State) String() string {
	r := s.Board.String()
	if s.IsComplete() {
		r = color.GreenString(r)
	}
	return r
}

func (s *State) ValidIndex() bool         { return s.squareIndex < len(s.squares) }
func (s *State) IsFirst() bool            { return s.piecesIndex == 0 }
func (s *State) IsComplete() bool         { return s.piecesIndex == s.Length }
func (s *State) Piece() chess.Piece       { return s.Pieces[s.piecesIndex] }
func (s *State) PrevPiece() chess.Piece   { return s.Pieces[s.piecesIndex-1] }
func (s *State) NextPiece() chess.Piece   { return s.Pieces[s.piecesIndex+1] }
func (s *State) Position() chess.Position { return s.squares[s.squareIndex] }

func (s *State) Placement() chess.Placement {
	return chess.Placement{Piece: s.Piece(), Position: s.Position()}
}

func (s *State) RemoveSquare(p chess.Position) {
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
	return !s.Board.Free(s.Position()) || s.Board.IsSafe(s.Placement())
}

func (s *State) Next() *State {
	if s.IsComplete() {
		s.Result <- s.Board.Combination()
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
		s.Board.Place(p)
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
	s.Board.RemoveLast()
	s.previous.squareIndex++
	return s.previous
}

func Solve(w, h int, count map[chess.Piece]int, debug bool) <-chan []chess.Placement {
	ch := make(chan []chess.Placement)
	reader := bufio.NewReader(os.Stdin)
	go func() {
		defer close(ch)
		l := piecelist(count)
		s := &State{common: common{
			Board:  chess.NewBoard(w, h),
			Pieces: l,
			Result: ch,
			Length: len(l),
		}}
		s.squares = s.Board.Positions()
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

func piecelist(count map[chess.Piece]int) []chess.Piece {
	var ps = make([]chess.Piece, 0, len(count))
	for p, n := range count {
		for i := 0; i < n; i++ {
			ps = append(ps, p)
		}
	}
	return ps
}
