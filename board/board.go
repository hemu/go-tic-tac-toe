package board

type Player int

const BOARD_SIZE int = 3

type Board struct {
  Grid [][]Player
}

func NewBoard() Board {
  b := Board{}
  b.Grid = make([][]Player, BOARD_SIZE)
  for i := 0; i < BOARD_SIZE; i++ {
    b.Grid[i] = make([]Player, BOARD_SIZE)
  }
  return b
}

type InvalidOperation struct {}
func (op InvalidOperation) Error() string {
  return "Invalid board operation"
}

func (b *Board) Mark(row int, col int, player Player) error {
  if row >= BOARD_SIZE || col >= BOARD_SIZE {
    return InvalidOperation{}
  }
  b.Grid[row][col] = Player(player)
  return nil
}
