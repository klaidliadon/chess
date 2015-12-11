package checkmate

import (
	"testing"

	. "gopkg.in/check.v1"
)

type CheckmateSuite struct{}

var _ = Suite(&CheckmateSuite{})

func Test(t *testing.T) { TestingT(t) }

type pieceCase struct {
	Type  Piece
	Cases []positionCase
}

type positionCase struct {
	Mine    Position
	Their   Position
	Menaced bool
}

func (s *CheckmateSuite) TestPieces(c *C) {
	var cases = []pieceCase{
		pieceCase{Bishop, []positionCase{
			{Position{0, 0}, Position{0, 0}, true},
			{Position{0, 0}, Position{1, 0}, false},
			{Position{0, 0}, Position{0, 1}, false},
			{Position{0, 0}, Position{1, 1}, true},
			{Position{3, 1}, Position{1, 3}, true},
		}},
		pieceCase{King, []positionCase{
			{Position{0, 0}, Position{0, 0}, true},
			{Position{0, 0}, Position{1, 0}, true},
			{Position{0, 0}, Position{0, 1}, true},
			{Position{0, 0}, Position{2, 0}, false},
			{Position{0, 0}, Position{0, 2}, false},
		}},
		pieceCase{Knight, []positionCase{
			{Position{0, 0}, Position{0, 0}, true},
			{Position{0, 0}, Position{1, 0}, false},
			{Position{0, 0}, Position{0, 1}, false},
			{Position{0, 0}, Position{2, 0}, false},
			{Position{0, 0}, Position{0, 2}, false},
			{Position{0, 0}, Position{1, 2}, true},
			{Position{0, 0}, Position{2, 1}, true},
			{Position{3, 1}, Position{5, 2}, true},
			{Position{3, 1}, Position{2, 3}, true},
		}},
		pieceCase{Rook, []positionCase{
			{Position{0, 0}, Position{0, 0}, true},
			{Position{0, 0}, Position{1, 0}, true},
			{Position{0, 0}, Position{0, 1}, true},
			{Position{0, 0}, Position{2, 0}, true},
			{Position{0, 0}, Position{0, 2}, true},
			{Position{0, 0}, Position{1, 2}, false},
			{Position{0, 0}, Position{2, 1}, false},
			{Position{3, 1}, Position{5, 2}, false},
			{Position{3, 1}, Position{2, 3}, false},
		}},
		pieceCase{Queen, []positionCase{
			{Position{0, 0}, Position{0, 0}, true},
			{Position{0, 0}, Position{1, 0}, true},
			{Position{0, 0}, Position{0, 1}, true},
			{Position{0, 0}, Position{2, 0}, true},
			{Position{0, 0}, Position{0, 2}, true},
			{Position{0, 0}, Position{1, 2}, false},
			{Position{0, 0}, Position{2, 1}, false},
			{Position{3, 1}, Position{5, 2}, false},
			{Position{3, 1}, Position{2, 3}, false},
		}},
		pieceCase{Piece(0), []positionCase{
			{Position{0, 0}, Position{0, 0}, true},
			{Position{0, 0}, Position{0, 1}, false},
		}},
	}

	for _, p := range cases {
		for _, t := range p.Cases {
			var action = "should"
			if !t.Menaced {
				action += "n't"
			}
			c.Logf("%s in %v %s menace %v", p.Type, t.Mine, action, t.Their)
			c.Assert(p.Type.Menaces(t.Mine, t.Their), Equals, t.Menaced)
		}
	}

}
