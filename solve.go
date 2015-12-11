package checkmate

func newState(w, h int, count map[Piece]int) *state {
	sq := make([]Position, 0, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			sq = append(sq, Position{x, y})
		}
	}
	var ps = make([]Piece, 0, len(count))
	for p, n := range count {
		for i := 0; i < n; i++ {
			ps = append(ps, p)
		}
	}
	b := NewBoard(w, h)
	return &state{Board: &b, Squares: sq, Pieces: ps}
}

type state struct {
	previous *state
	Board    *Board
	Squares  []Position
	Pieces   []Piece
	Index    int
}

func (s *state) Next() *state {
	if len(s.Pieces) == 0 {
		return s.Previous()
	}
	piece := s.Pieces[0]
loop:
	for s.Index < len(s.Squares) {
		pos := s.Squares[s.Index]
		if _, ok := s.Board.Squares[pos]; ok {
			s.Index++
			continue
		}
		for _, p := range s.Board.placements {
			if piece.Menaces(pos, p) {
				s.Index++
				continue loop
			}
		}
		s.Board.Place(piece, pos)
		safe, _ := piece.Split(pos, s.Squares)
		var index int
		if s.previous != nil && s.previous.Pieces[0] == piece {
			for i, v := range safe {
				if pos.Before(v) {
					index = i
					break
				}
			}
		}
		return &state{
			previous: s,
			Board:    s.Board,
			Squares:  safe,
			Pieces:   s.Pieces[1:],
			Index:    index,
		}
	}
	return s.Previous()
}

func (s *state) Previous() *state {
	if s.previous == nil {
		return nil
	}
	s.Board.RemoveLast()
	s.previous.Index++
	return s.previous
}

func Solve(w, h int, count map[Piece]int) <-chan []Placement {
	ch := make(chan []Placement)
	go func() {
		defer close(ch)
		for s := newState(w, h, count); s != nil; s = s.Next() {
			if len(s.Pieces) == 0 {
				ch <- s.Board.Combination()
			}
		}
	}()
	return ch
}
