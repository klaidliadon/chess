package checkmate

import (
	. "gopkg.in/check.v1"
)

func (s *CheckmateSuite) TestBefore(c *C) {
	var cases = map[[4]int]bool{
		{0, 0, 0, 1}: true, {0, 2, 0, 1}: false, {1, 1, 0, 1}: false, {0, 1, 1, 1}: true,
	}
	for coord, result := range cases {
		mine, their := Position{coord[0], coord[1]}, Position{coord[2], coord[3]}
		c.Assert(mine.Before(their), Equals, result)
	}

}

func (s *CheckmateSuite) TestPieces(c *C) {
	var cases = map[Piece]map[[4]int]bool{
		Piece(0): {
			{0, 0, 0, 0}: true, {0, 0, 0, 1}: false,
		},
		Bishop: {
			{0, 0, 0, 0}: true, {0, 0, 1, 0}: false, {3, 1, 1, 3}: true,
			{0, 0, 1, 1}: true, {0, 0, 0, 1}: false,
		},
		King: {
			{0, 0, 0, 0}: true, {0, 0, 2, 0}: false, {0, 0, 0, 1}: true,
			{0, 0, 1, 0}: true, {0, 0, 0, 2}: false,
		},
		Knight: {
			{0, 0, 0, 0}: true, {0, 0, 1, 0}: false, {0, 0, 0, 1}: false,
			{3, 1, 5, 2}: true, {0, 0, 0, 2}: false, {0, 0, 1, 2}: true,
			{0, 0, 2, 1}: true, {0, 0, 2, 0}: false, {3, 1, 2, 3}: true,
		},
		Rook: {
			{0, 0, 0, 0}: true, {0, 0, 1, 2}: false, {0, 0, 1, 0}: true,
			{0, 0, 2, 0}: true, {3, 1, 2, 3}: false, {0, 0, 0, 2}: true,
			{0, 0, 0, 1}: true, {3, 1, 5, 2}: false, {0, 0, 2, 1}: false,
		},
		Queen: {
			{0, 0, 0, 0}: true, {3, 1, 2, 3}: false, {0, 0, 1, 0}: true,
			{0, 0, 2, 0}: true, {0, 0, 1, 2}: false, {0, 0, 0, 2}: true,
			{0, 0, 0, 1}: true, {0, 0, 2, 1}: false, {3, 1, 5, 2}: false,
		},
	}
	for piece, list := range cases {
		var splitChecker = make(map[Position]struct {
			safe, unsafe []Position
		})
		for coord, menace := range list {
			var action = "should"
			if !menace {
				action += "n't"
			}
			mine, their := Position{coord[0], coord[1]}, Position{coord[2], coord[3]}
			v := splitChecker[mine]
			if menace {
				v.unsafe = append(v.unsafe, their)
			} else {
				v.safe = append(v.safe, their)

			}
			splitChecker[mine] = v
			c.Logf("%s in %v %s menace %v", piece, mine, action, their)
			c.Assert(piece.Menaces(mine, their), Equals, menace)
		}
		for mine, theirs := range splitChecker {
			safe, unsafe := piece.Split(mine, append(theirs.safe, theirs.unsafe...))
			c.Assert(safe, DeepEquals, theirs.safe)
			c.Assert(unsafe, DeepEquals, theirs.unsafe)
		}
	}

}
