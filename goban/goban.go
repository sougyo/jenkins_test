package goban

import (
	"fmt"
	"strings"
)

const (
	BoardSize = 19
	Empty     = '.'
	Black     = 'B'
	White     = 'W'
)

type Move struct {
	X     int
	Y     int
	Color rune
}

type GoBoard struct {
	board [][]rune
	moves []Move
}

func NewGoBoard() *GoBoard {
	board := make([][]rune, BoardSize)
	for i := range board {
		board[i] = make([]rune, BoardSize)
		for j := range board[i] {
			board[i][j] = Empty
		}
	}
	return &GoBoard{board: board, moves: []Move{}}
}

func (gb *GoBoard) PlaceStone(x, y int, color rune) {
	if x >= 0 && x < BoardSize && y >= 0 && y < BoardSize && gb.board[x][y] == Empty {
		gb.board[x][y] = color
		gb.moves = append(gb.moves, Move{x, y, color})
	}
}

func (gb *GoBoard) Display() string {
	var displayString strings.Builder
	for y := 0; y < BoardSize; y++ {
		for x := 0; x < BoardSize; x++ {
			displayString.WriteString(string(gb.board[x][y]))
		}
		displayString.WriteString("\n")
	}
	displayString.WriteString("\n")
	return displayString.String()
}

func (gb *GoBoard) ToSGF() string {
	var sgf strings.Builder
	sgf.WriteString("(;GM[1]FF[4]SZ[19]")
	for _, move := range gb.moves {
		if move.Color == Black {
			sgf.WriteString(fmt.Sprintf(";B[%c%c]", 'a'+move.X, 'a'+move.Y))
		} else if move.Color == White {
			sgf.WriteString(fmt.Sprintf(";W[%c%c]", 'a'+move.X, 'a'+move.Y))
		}
	}
	sgf.WriteString(")")
	return sgf.String()
}
