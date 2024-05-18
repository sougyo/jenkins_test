// goban_test.go
package goban

import (
	"strings"
	"testing"
)

func TestNewGoBoard(t *testing.T) {
	board := NewGoBoard()
	for i := range board.board {
		for j := range board.board[i] {
			if board.board[i][j] != Empty {
				t.Errorf("Expected Empty at (%d, %d), but got %c", i, j, board.board[i][j])
			}
		}
	}
}

func TestPlaceStone(t *testing.T) {
	board := NewGoBoard()
	board.PlaceStone(3, 3, Black)
	if board.board[3][3] != Black {
		t.Errorf("Expected Black at (3, 3), but got %c", board.board[3][3])
	}

	board.PlaceStone(16, 16, White)
	if board.board[16][16] != White {
		t.Errorf("Expected White at (16, 16), but got %c", board.board[16][16])
	}

	board.PlaceStone(19, 19, Black)
	if board.board[18][18] != Empty {
		t.Errorf("Expected Empty at (18, 18) since (19, 19) is out of bounds, but got %c", board.board[18][18])
	}

	board.PlaceStone(3, 3, White)
	if board.board[3][3] != Black {
		t.Errorf("Expected Black at (3, 3) since the spot is already occupied, but got %c", board.board[3][3])
	}
}

func TestToSGF(t *testing.T) {
	board := NewGoBoard()
	board.PlaceStone(3, 5, Black)
	board.PlaceStone(4, 16, White)
	board.PlaceStone(10, 9, Black)
	board.PlaceStone(6, 6, White)

	sgf := board.ToSGF()
	expected := "(;GM[1]FF[4]SZ[19];B[df];W[eq];B[kj];W[gg])"
	if strings.TrimSpace(sgf) != expected {
		t.Errorf("Expected SGF: %s, but got: %s", expected, sgf)
	}
}
