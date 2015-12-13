package checkmate

import (
	"runtime"
	"time"

	. "gopkg.in/check.v1"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func (s *CheckmateSuite) baseTest(c *C, w, h int, m map[Piece]int, count int) {
	start := time.Now()
	var i int
	for v := range Solve(w, h, m, false) {
		i++
		s.verify(c, v)
	}
	if count > -1 {
		c.Assert(i, Equals, count)
	}
	c.Logf("%dx%d Board, %d solutions in %.3f seconds", w, h, i, float64(time.Since(start).Seconds()))
}

func (s *CheckmateSuite) verify(c *C, v []Placement) {
	for i, p1 := range v {
		for j, p2 := range v {
			m := p2.Menaces(p1)
			c.Assert(m, Equals, i == j)
		}
	}
}

func (s *CheckmateSuite) TestSolveQueens(c *C) {
	for i, v := range []int{1, 0, 0, 2, 10, 4, 40, 92, 352} {
		s.baseTest(c, i+1, i+1, map[Piece]int{Queen: i + 1}, v)
	}
}

func (s *CheckmateSuite) TestSolveComplex(c *C) {
	s.baseTest(c, 7, 7, map[Piece]int{King: 2, Bishop: 2, Queen: 2, Knight: 1}, 3063828)
}
