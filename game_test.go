package tic_tac_toe_test

import (
	"testing"

	tic_tac_toe "github.com/jerray/tic-tac-toe"
)

func Test_NewGame(t *testing.T) {
	_ = tic_tac_toe.NewGame()
}

func Test_GameShouldHavePlayFunction(t *testing.T) {
	game := tic_tac_toe.NewGame()
	_, err := game.Play(0, 0)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func Test_GamePlayFirstCallShouldReturnX(t *testing.T) {
	game := tic_tac_toe.NewGame()
	r, _ := game.Play(0, 0)
	if r != 'x' {
		t.Errorf("first play should return x, got %s", string(r))
	}
}

func Test_GamePlayShouldBeTurnByTurn(t *testing.T) {
	game := tic_tac_toe.NewGame()
	cases := []struct {
		x, y     uint
		expected byte
	}{
		{0, 0, 'x'},
		{0, 1, 'o'},
		{1, 1, 'x'},
		{2, 0, 'o'},
	}
	for _, c := range cases {
		r, err := game.Play(c.x, c.y)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if r != c.expected {
			t.Errorf("expected result is %s, actual is %s", string(c.expected), string(r))
		}
	}
}

func Test_GamePlayShouldNotBeDuplicateLocation(t *testing.T) {
	game := tic_tac_toe.NewGame()
	_, err := game.Play(0, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	_, err = game.Play(0, 1)
	if err != tic_tac_toe.DuplicateLocationError {
		t.Errorf("should report duplicate location error")
	}
}

func Test_GamePlayShouldNotBeOutOfRange(t *testing.T) {
	game := tic_tac_toe.NewGame()
	_, err := game.Play(5, 0)
	if err != tic_tac_toe.OutOfRangeError {
		t.Errorf("should report out of range error on X")
	}
	_, err = game.Play(2, 4)
	if err != tic_tac_toe.OutOfRangeError {
		t.Errorf("should report out of range error on Y")
	}
}

func Test_GamePlayWinOnX(t *testing.T) {
	game := tic_tac_toe.NewGame()
	_, _ = game.Play(0, 0)    // x
	_, _ = game.Play(0, 1)    // o
	_, _ = game.Play(1, 0)    // x
	_, _ = game.Play(1, 1)    // o
	r, err := game.Play(2, 0) // x
	if err != tic_tac_toe.GameIsJustFinishedError {
		t.Errorf("should report game is just finished error")
	}
	if r != 'x' {
		t.Errorf("winner should be x, got %s", string(r))
	}
}

func Test_GamePlayWinOnY(t *testing.T) {
	game := tic_tac_toe.NewGame()
	_, _ = game.Play(0, 0)    // x
	_, _ = game.Play(1, 1)    // o
	_, _ = game.Play(0, 1)    // x
	_, _ = game.Play(2, 1)    // o
	r, err := game.Play(0, 2) // x
	if err != tic_tac_toe.GameIsJustFinishedError {
		t.Errorf("should report game is just finished error")
	}
	if r != 'x' {
		t.Errorf("winner should be x, got %s", string(r))
	}
}

func Test_GamePlayWinOnCross(t *testing.T) {
	game := tic_tac_toe.NewGame()
	_, _ = game.Play(0, 0)    // x
	_, _ = game.Play(0, 1)    // o
	_, _ = game.Play(1, 1)    // x
	_, _ = game.Play(0, 2)    // o
	r, err := game.Play(2, 2) // x
	if err != tic_tac_toe.GameIsJustFinishedError {
		t.Errorf("should report game is just finished error")
	}
	if r != 'x' {
		t.Errorf("winner should be x, got %s", string(r))
	}
}

func Test_GamePlayWinOnReversedCross(t *testing.T) {
	game := tic_tac_toe.NewGame()
	_, _ = game.Play(2, 0)    // x
	_, _ = game.Play(0, 1)    // o
	_, _ = game.Play(0, 2)    // x
	_, _ = game.Play(1, 2)    // o
	r, err := game.Play(1, 1) // x
	if err != tic_tac_toe.GameIsJustFinishedError {
		t.Errorf("should report game is just finished error")
	}
	if r != 'x' {
		t.Errorf("winner should be x, got %s", string(r))
	}
}

func Test_GameResultIsDrawn(t *testing.T) {
	game := tic_tac_toe.NewGame()
	// xox
	// oox
	// xxo
	_, _ = game.Play(0, 0)    // x
	_, _ = game.Play(1, 0)    // o
	_, _ = game.Play(2, 0)    // x
	_, _ = game.Play(0, 1)    // o
	_, _ = game.Play(2, 1)    // x
	_, _ = game.Play(1, 1)    // o
	_, _ = game.Play(0, 2)    // x
	_, _ = game.Play(2, 2)    // o
	_, err := game.Play(1, 2) // x
	if err != tic_tac_toe.GameResultIsDrawn {
		t.Errorf("should report game result is drawn")
	}
}
