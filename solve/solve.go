package solve

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"
	"github.com/klaidliadon/chess"
)

type sortPiece []chess.Piece

func (s sortPiece) Len() int           { return len(s) }
func (s sortPiece) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortPiece) Less(i, j int) bool { return s[i] > s[j] }

func NewCheckmate(w, h int, pieces chess.PieceCount) *Checkmate {
	var list = pieces.List()
	sort.Sort(sortPiece(list))
	return &Checkmate{
		Board:  chess.NewBoard(w, h),
		Pieces: list,
	}
}

// Checkmate solves the extended n queen problem (any piece excluding pawns, any board size)
type Checkmate struct {
	*state
	Board  *chess.Board
	Pieces []chess.Piece
}

func (c *Checkmate) reset() {
	c.state = &state{squares: c.Board.Positions()}
}

func (c *Checkmate) Count() int {
	c.reset()
	i := 0
	for c.next() {
		if c.isComplete() {
			c.prev()
			i++
		}
	}
	return i
}

func (c *Checkmate) Solve(debug bool) chan []chess.Placement {
	c.reset()
	ch := make(chan []chess.Placement)
	var reader interface {
		ReadString(byte) (string, error)
	}
	if debug {
		reader = bufio.NewReader(os.Stdin)
	}
	go func() {
		defer close(ch)
		for c.next() {
			if reader != nil {
				fmt.Println(c)
				reader.ReadString('\n')
			}
			if c.isComplete() {
				ch <- c.Board.Combination()
				c.prev()
			}
		}
	}()
	return ch
}

func (c *Checkmate) String() string {
	r := c.Board.String()
	if c.isComplete() {
		r = color.GreenString(r)
	}
	return r
}

func (c *Checkmate) isComplete() bool       { return c.piecesIndex == len(c.Pieces) }
func (c *Checkmate) currPiece() chess.Piece { return c.Pieces[c.piecesIndex] }
func (c *Checkmate) prevPiece() chess.Piece { return c.Pieces[c.piecesIndex-1] }
func (c *Checkmate) nextPiece() chess.Piece { return c.Pieces[c.piecesIndex+1] }

func (c *Checkmate) invalid() bool {
	pos := c.position()
	return !c.Board.Free(pos) || c.Board.IsSafe(chess.Placement{
		Piece: c.currPiece(), Position: pos,
	})
}

func (c *Checkmate) next() bool {
	if c.isComplete() {
		return true
	}
	if !c.isFirst() && c.currPiece() == c.prevPiece() {
		prev := c.previous.position()
		for c.validIndex() && c.position().Before(prev) {
			c.squareIndex++
		}
	}
	for c.squareIndex < len(c.squares) {
		p := chess.Placement{Position: c.position(), Piece: c.currPiece()}
		if c.invalid() {
			c.squareIndex++
			continue
		}
		c.Board.Place(p)
		safe, _ := p.Split(c.squares)
		c.state = &state{
			previous:    c.state,
			squares:     safe,
			piecesIndex: c.piecesIndex + 1,
		}
		return true
	}
	return c.prev()
}

func (c *Checkmate) prev() bool {
	if c.previous == nil {
		return false
	}
	c.Board.RemoveLast()
	c.previous.squareIndex++
	c.state = c.previous
	return true
}
