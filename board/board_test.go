package board

import (
  "testing"
  // "fmt"
)

func TestBoardGridSize(t *testing.T) {
  b := NewBoard()
  if len(b.Grid) != 3 || 
     len(b.Grid[0]) != 3 ||
     len(b.Grid[1]) != 3 ||
     len(b.Grid[2]) != 3 {
    t.Errorf("Did not find expected %d by %d square grid", 3, 3)
  }
}

func TestBoardMark(t *testing.T) {
  b := NewBoard()
  if b.Grid[0][0] != 0 {
    t.Errorf("Expected cell to be unmarked")
  }
  b.Mark(0, 0, 1)
  if b.Grid[0][0] != 1 {
    t.Errorf("Expected cell to be marked with player 1")
  }
  b.Mark(0, 0, 2)
  if b.Grid[0][0] != 2 {
    t.Errorf("Expected cell to be marked with player 2")
  }
}
