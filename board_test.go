package checkmate

import (
	. "gopkg.in/check.v1"
)

func (s *CheckmateSuite) TestContains(c *C) {
	size := Size{5, 8}
	var contains = map[[2]int]bool{
		{5, 0}: false, {4, 0}: true, {3, 1}: true,
		{0, 8}: false, {0, 7}: true, {2, 6}: true,
		{5, 8}: false, {0, 0}: true, {3, 3}: true,
	}
	for p, in := range contains {
		c.Assert(size.Contains(Position{p[0], p[1]}), Equals, in)
	}
}

func (s *CheckmateSuite) TestPlace(c *C) {
	b := NewBoard(6, 7)
	c.Assert(func() { b.RemoveLast() }, Panics, errEmpty)
	c.Assert(func() { b.Place(Queen, Position{7, 7}) }, Panics, errInvalid)

	var (
		pos   Position
		piece Piece
		i     int
	)

	pos, piece, i = Position{5, 6}, Queen, i+1
	b.Place(piece, pos)
	c.Assert(b.Squares, HasLen, i)
	c.Assert(b.placements, HasLen, i)
	c.Assert(b.Squares[pos], Equals, piece)

	pos, piece, i = Position{4, 6}, Bishop, i+1
	b.Place(piece, pos)
	c.Assert(b.Squares, HasLen, i)
	c.Assert(b.placements, HasLen, i)
	c.Assert(b.Squares[pos], Equals, piece)

	lastPiece, lastPos := b.RemoveLast()
	c.Assert(b.Squares, HasLen, i-1)
	c.Assert(b.placements, HasLen, i-1)
	c.Assert(lastPos, Equals, pos)
	c.Assert(lastPiece, Equals, piece)

}
