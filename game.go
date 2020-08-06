package tic_tac_toe

import "errors"

var DuplicateLocationError = errors.New("duplicate location")
var OutOfRangeError = errors.New("out of range")
var GameIsJustFinishedError = errors.New("game is just finished")
var GameResultIsDrawn = errors.New("game result is drawn")

const ROWS = 3

type Game struct {
	current byte
	board   [ROWS][ROWS]byte
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Play(x, y uint) (byte, error) {
	player := g.nextPlayer()

	if x >= ROWS || y >= ROWS {
		return 0, OutOfRangeError
	}

	if g.board[x][y] != 0 {
		return 0, DuplicateLocationError
	}
	g.board[x][y] = player

	if g.isWin(x, y) {
		return player, GameIsJustFinishedError
	} else if g.isDrawn() {
		return player, GameResultIsDrawn
	}

	return player, nil
}

func (g *Game) nextPlayer() byte {
	if g.current == 'x' {
		g.current = 'o'
	} else {
		g.current = 'x'
	}
	return g.current
}

func (g *Game) isWin(x, y uint) bool {
	return g.checkOnX(x, y) ||
		g.checkOnY(x, y) ||
		g.checkOnCross(x, y) ||
		g.checkOnReversedCross(x, y)
}

func (g *Game) checkOnX(_, y uint) bool {
	var c byte
	for i := 0; i < ROWS; i++ {
		t := g.board[i][y]
		if t == 0 {
			return false
		}
		if c == 0 {
			c = t
		}
		if c != t {
			return false
		}
	}
	return true
}

func (g *Game) checkOnY(x, _ uint) bool {
	var c byte
	for i := 0; i < ROWS; i++ {
		t := g.board[x][i]
		if t == 0 {
			return false
		}
		if c == 0 {
			c = t
		}
		if c != t {
			return false
		}
	}
	return true
}

func (g *Game) checkOnCross(x, y uint) bool {
	if x == y {
		var c byte
		for i := 0; i < ROWS; i++ {
			t := g.board[i][i]
			if t == 0 {
				return false
			}
			if c == 0 {
				c = t
			}
			if c != t {
				return false
			}
		}
		return true
	}
	return false
}

func (g *Game) checkOnReversedCross(x, y uint) bool {
	if x+y == ROWS-1 {
		var c byte
		for i := 0; i < ROWS; i++ {
			t := g.board[i][ROWS-1-i]
			if t == 0 {
				return false
			}
			if c == 0 {
				c = t
			}
			if c != t {
				return false
			}
		}
		return true
	}
	return false
}

func (g *Game) isDrawn() bool {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < ROWS; j++ {
			if g.board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
