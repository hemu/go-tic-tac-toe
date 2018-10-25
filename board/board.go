package board

type Player int

type Board struct {
  Size int
  Grid [][]Player
}

func NewBoard(size int) Board {
  b := Board{Size: size}
  b.Grid = make([][]Player, size)
  for i := 0; i < size; i++ {
    b.Grid[i] = make([]Player, size)
  }
  return b
}

type InvalidOperation struct {}
func (op InvalidOperation) Error() string {
  return "Invalid board operation"
}

func (b *Board) Mark(row int, col int, player Player) error {
  if row >= b.Size || col >= b.Size {
    return InvalidOperation{}
  }
  b.Grid[row][col] = Player(player)
  return nil
}
