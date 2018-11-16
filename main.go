package main

import (
	"bufio"
	"fmt"
	"github.com/hmuar/go-tic-tac-toe/board"
	"os"
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

func endGame(b *board.Board, winningPlayer board.Player) {
	if winningPlayer == board.Player(-1) {
		fmt.Println("Game over, it was a draw :D")
	} else {
		fmt.Printf("Player %v wins! :D\n", winningPlayer)
	}
}

func gameStep(b *board.Board, player int) int {
	row, col := doPlayerMove(player)
	b.Mark(row, col, board.Player(player))
	displayBoard(b)
	return (player + 1) % 2
}

func displayBoard(b *board.Board) {
	fmt.Println(b)
}

func main() {
	fmt.Println("Tic Tac Toe :D")
	b := board.NewBoard()
	displayBoard(b)
	gameOver := false
	player := 0
	for !gameOver {
		player = gameStep(b, player)
		gameOver, winningPlayer := b.IsGameOver()
		if gameOver {
			endGame(b, winningPlayer)
			break
		}
	}
}
