package board

import (
  "fmt"
  "strconv"
)

type Player int

const BOARD_SIZE int = 3

var winMasks [8]int64

func init() {
  var winBoards = []string{
    "111000000",
    "000111000",
    "000000111",
    "100100100",
    "010010010",
    "001001001",
    "100010001",
    "001010100",
  }

  for i, winBoard := range winBoards {
    boardMask, err := boardDisplayToMask(winBoard)
    if err != nil {
      continue
    }
    winMasks[i] = boardMask
  }
}

func boardDisplayToMask(display string) (int64, error) {
  boardMask, err := strconv.ParseInt(display, 2, 64)
  return boardMask, err
}


func IsWinBoard(board int64) bool {
  for _, winMask := range winMasks {
    if winMask & board == winMask {
      return true
    }
  }
  return false
}

type Board struct {
  players [2]int64
  board int64
}

func NewBoard() Board {
  b := Board{}
  return b
}

type InvalidOperation struct {}
func (op InvalidOperation) Error() string {
  return "Invalid board operation"
}

func markMask(row, col int) int64 {
  oneDimensionCoord := row * BOARD_SIZE + col
  shiftAmt := uint(BOARD_SIZE*BOARD_SIZE - 1 - oneDimensionCoord)
  return int64(1 << shiftAmt)
}

func (b *Board) Mark(row int, col int, player Player) error {
  if row >= BOARD_SIZE || col >= BOARD_SIZE || player > 1 || player < 0 {
    return InvalidOperation{}
  }

  b.players[player] |= markMask(row, col)
  b.update()
  return nil
}

func (b *Board) update() {
  b.board = 0
  for _, player := range b.players {
    b.board |= player
  }
}

func playerDisp(player Player) string {
  switch player {
    case 0:
      return "O"
    case 1:
      return "X"
    default:
      return "."
  }
}

func (b Board) String() string {
  display := ""
  for i := 0; i < BOARD_SIZE; i++ {
    rowDisp := "|"
    for j := 0; j < BOARD_SIZE; j++ {
      mask := markMask(i, j)
      cellDisplay := "."
      for i, player := range b.players {
        if player & mask != 0 {
          cellDisplay = playerDisp(Player(i))
          break
        }
      }
      rowDisp += fmt.Sprintf(" %s |", cellDisplay) 
    }
    display += rowDisp + "\n"
  }
  return display
}
