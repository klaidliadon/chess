// generated by stringer -type=Piece; DO NOT EDIT

package checkmate

import "fmt"

const _Piece_name = "KingKnightBishopRookQueen"

var _Piece_index = [...]uint8{0, 4, 10, 16, 20, 25}

func (i Piece) String() string {
	i -= 1
	if i < 0 || i >= Piece(len(_Piece_index)-1) {
		return fmt.Sprintf("Piece(%d)", i+1)
	}
	return _Piece_name[_Piece_index[i]:_Piece_index[i+1]]
}
