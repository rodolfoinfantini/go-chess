package board

import (
	"fmt"
	"piece"
)

type Castling struct {
	WhiteKingSide  bool
	WhiteQueenSide bool
	BlackKingSide  bool
	BlackQueenSide bool
}
type Board struct {
	Pieces   [8][8]*piece.Piece
	turn     piece.Color
	Castling Castling
}

func (b *Board) fillDefaults() {
	b.Castling = Castling{true, true, true, true}
}
func (b *Board) GetTurn() piece.Color {
	return b.turn
}
func (b *Board) ChangeTurn() {
	b.turn = b.turn.Opposite()
}
func (b *Board) PieceAt(x uint, y uint) *piece.Piece {
	if x > 7 || y > 7 {
		return nil
	}

	return b.Pieces[y][x]
}
func (b *Board) MovePiece(from piece.Position, to piece.Position) {
	if from == to {
		panic("Can't move to the same position")
	}

	fromPiece := b.PieceAt(from.X, from.Y)
	if fromPiece == nil {
		panic("No piece at position")
	}

	if fromPiece.Color != b.turn {
		panic("Not your turn")
	}

	toPiece := b.PieceAt(to.X, to.Y)
	if toPiece != nil && toPiece.Color == b.turn {
		panic("Can't capture your own piece")
	}

	fromPiece.Move(to)
	b.Pieces[to.Y][to.X] = fromPiece
	b.removePiece(from.X, from.Y)
}
func (b *Board) removePiece(x uint, y uint) {
	b.Pieces[y][x] = nil
}

func IndexToNotation(x uint, y uint) string {
	if x > 7 || y > 7 {
		return ""
	}
	letters := "abcdefgh"
	return string(letters[x]) + fmt.Sprint(8-y)
}

func InitialPositionBoard() Board {
	initialFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	return BoardFromFen(initialFen)
}
