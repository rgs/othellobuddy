package main

import(
	"fmt"
	"bufio"
	"os"
	"regexp"
	"github.com/rgs/othellobuddy/board"
)

// Commands
// - play at x, y
// - pass
// - go back
// - find play that maximizes/minimizes counts for a given player, relative or differential
// - find play that maximizes/minimizes possibilities for a given player, relative or differential

func main() {
	fmt.Println("Othello Buddy")
	curBoard := board.InitialBoard()
	curPlayer := board.BLACK
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("-----------------")
		fmt.Println(curBoard)
		if curPlayer == board.BLACK {
			fmt.Println("Black \u25CF to play")
		} else {
			fmt.Println("White \u25CB to play")
		}
		fmt.Println("Enter position, or 'exit':")
		cmd, _ := reader.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		if cmd == "exit" {
			break
		}
		if isPosition, _ := regexp.MatchString("^[a-h][1-8]$", cmd) ; isPosition {
			var x, y int
			x = int(cmd[1]) - int('1')
			y = int(cmd[0]) - int('a')
			var position board.Coordinates = board.Coordinates{x,y}
			newBoard, err := curBoard.PlayAt(curPlayer, position)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				curBoard = newBoard
				if curBoard.IsFinal() {
					fmt.Println("End of game")
					fmt.Println(curBoard)
					break
				}
				// select new player. Check if we need to pass
				blackP, whiteP := curBoard.PossibilitiesByPlayer()
				switch curPlayer {
				case board.BLACK:
					if whiteP == 0 {
						fmt.Println("White \u25CB passes!")
					} else {
						curPlayer = board.WHITE
					}
				case board.WHITE:
					if blackP == 0 {
						fmt.Println("Black \u25CF passes!")
					} else {
						curPlayer = board.BLACK
					}
				}
			}
		} else {
			fmt.Println("?")
		}
	}
}
