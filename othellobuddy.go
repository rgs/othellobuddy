package main

import "fmt"
import "github.com/rgs/othellobuddy/board"

// Commands
// - play at x, y
// - pass
// - go back
// - find play that maximizes/minimizes counts for a given player, relative or differential
// - find play that maximizes/minimizes possibilities for a given player, relative or differential

func main() {
	fmt.Println("Othello Buddy")
	fmt.Println("-------------")
	fmt.Println(board.InitialBoard())
}
