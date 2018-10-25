package main

import (
  "fmt"
  "bufio"
  "os"
  "./board"
  "strconv"
)

func playerInput(prompt string) int {
  for {
    fmt.Printf(prompt)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    playerMove, err := strconv.ParseInt(scanner.Text(), 10, 32)
    if err != nil {
      fmt.Println("Invalid move")
    } else {
      return int(playerMove)
    }
  }
}

func doPlayerMove(player int) (int, int) {
  fmt.Printf("-- Player %v turn --\n", player)
  row := playerInput("row:")
  col := playerInput("col:")
  return row, col
}

func main() {
  fmt.Println("Tic Tac Toe :D\n")
  b := board.NewBoard()
  fmt.Println(b)
  gameOver := false
  player := 0
  for !gameOver {
    row, col := doPlayerMove(player)
    b.Mark(row, col, board.Player(player))
    player = (player + 1) % 2
    fmt.Println(b)
    gameOver, winningPlayer := b.IsGameOver()
    if gameOver {
      if winningPlayer == board.Player(-1) {
        fmt.Println("Game over, it was a draw :D")
      } else {
        fmt.Printf("Player %v wins! :D\n", winningPlayer)
      }
      break
    }
  }
}
