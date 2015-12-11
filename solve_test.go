package checkmate

import (
	"fmt"
	"runtime"

	. "gopkg.in/check.v1"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func (s *CheckmateSuite) baseTest(c *C, w, h int, m map[Piece]int, count int, echo bool) {
	var k int
	for s := range Solve(w, h, m) {
		k++
		for i := range s {
			for j := range s {
				m := s[j].Menaces(s[i])
				if i == j {
					c.Assert(m, Equals, true)
					continue
				}
				if m {
					c.Errorf("Invalid combination %d %s\nconflict [%d-%d] %v - %v", k, s, i, j, s[i], s[j])
					return
				}
			}
		}
		if echo {
			fmt.Println(s)
		}

	}
	fmt.Println(k, count)
}

func (s *CheckmateSuite) Test1Queens(c *C) {
	s.baseTest(c, 4, 4, map[Piece]int{Queen: 4}, 0, true)
}

func (s *CheckmateSuite) Test2Mix(c *C) {
	s.baseTest(c, 4, 4, map[Piece]int{King: 1, Bishop: 1, Knight: 1}, 0, true)
}

func (s *CheckmateSuite) Test3BigOne(c *C) {
	s.baseTest(c, 7, 7, map[Piece]int{King: 2, Bishop: 2, Queen: 2, Knight: 1}, 3063828, false)
}
