package board

import (
  "testing"
  "fmt"
)

func TestPlayerMark(t *testing.T) {
  b := NewBoard()
  if b.players[0] != 0 {
    t.Errorf("Expected empty player board")
  }
  b.Mark(2, 2, 0)
  if b.players[0] != 1 {
    t.Errorf("Expected player 0 board to be marked at (2, 2), board:%v", b.players[0])
  }
  b.Mark(2, 1, 0)
  if b.players[0] != 3 {
    t.Errorf("Expected player 0 board to be marked at (2, 1) and (2, 2), board:%v", b.players[0])
  }
  b.Mark(2, 0, 1)
  if b.players[0] != 3 {
    t.Errorf("Expected player 0 board to be only marked at (2, 1) and (2, 2), board:%v", b.players[0])
  }
  if b.players[1] != 4 {
    t.Errorf("Expected player 1 board to be marked at (2, 0) board:%v", b.players[1])
  }
  fmt.Println(b)
}

func TestBoardMark(t *testing.T) {
  b := NewBoard()
  if b.board != 0 {
    t.Errorf("Expected empty board")
  }
  b.Mark(2, 2, 0)
  if b.board != 1 {
    t.Errorf("Expected board with only (2,2) cell marked, board:%v", b.board)
  }
  b.Mark(2, 1, 1)
  if b.board != 3 {
    t.Errorf("Expected board with two cells marked, board:%v", b.board)
  }
}

func TestBoardWin(t *testing.T) {
  tests := []struct {
    board string
    want bool
  }{
    // Win by rows
    { "111000100", true },
    { "110111000", true },
    { "110000111", true },
    // Win by cols
    { "100100101", true },
    { "010010011", true },
    { "001001101", true },
    // Win by diagonals
    { "101010001", true },
    { "101010100", true },
    // Random non-winning board states
    { "101000100", false },
    { "110101000", false },
    { "110000110", false },
    { "000100101", false },
    { "010010001", false },
    { "001001100", false },
    { "101010000", false },
    { "100010100", false },
  }

  for _, test := range tests {
    mask, _ := boardDisplayToMask(test.board)
    if isWinBoard(mask) != test.want {
      t.Errorf("Expected board %v to return isWinBoard:%t", test.board, test.want)
    }
  }
}

