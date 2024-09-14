package main


import(
	"fmt"
	"strconv"
)


type piece string
const (
	Circle piece = "O"
	Cross piece = "X"
	None piece = "-"
)

type turn string
const (
	Player1 turn = "p1"
	Player2 turn = "p2"
)


func main() {
	turn := Player1
	gameBoard := [][]piece {
		{None, None, None},
		{None, None, None},
		{None, None, None},
	}
	
	fmt.Println("Let the game begin!")

	for !gameOver(gameBoard){
		fmt.Println("------")
		playTurn(turn, gameBoard)
		fmt.Println("------")
		turn = changeTurn(turn)
	}
	turn = changeTurn(turn)
	fmt.Println("------")

	if checkColumns(gameBoard) || checkDiagonals(gameBoard) || checkRows(gameBoard) {
		fmt.Println("Game Over!")
		fmt.Printf("%s wins!\n", turn)
	} else {
		fmt.Println("Game Over!")
		fmt.Println("Draw!")
	}
}


func gameOver(gameBoard [][]piece) bool {
	return checkRows(gameBoard) || checkColumns(gameBoard) || checkDiagonals(gameBoard) || boardFull(gameBoard)
}

func playTurn (turn turn, gameBoard [][]piece) {
	for {
		printGameBoard(gameBoard)
		coordinates := getCoordinates(turn)
		row := coordinates[0]
		col := coordinates[1]

		if gameBoard[row-1][col-1] == None {
			if turn == Player1 {
				gameBoard[row-1][col-1] = Cross
			} else {
				gameBoard[row-1][col-1] = Circle
			}
			break
		} else {
        	fmt.Println("Position already taken! Try again.")
        }
	}
}

func changeTurn (turn turn) (turn) {
	if turn == Player1 {
		return Player2
	}
	return Player1
}


func getCoordinates(turn turn) ([] int) {
    var row string
    fmt.Print("Enter a row: ")
    fmt.Scanln(&row)

    var col string
    fmt.Print("Enter a col: ")
    fmt.Scanln(&col)
    
	rowInt, rowErr := strconv.Atoi(row)
	colInt, colErr := strconv.Atoi(col)

	if rowErr != nil || colErr != nil {
		fmt.Println("Invalid input. Please enter valid numbers.")
		return getCoordinates(turn)
	}

	if rowInt > 3 || rowInt < 1 || colInt > 3 || colInt < 1 {
		fmt.Println("Invalid input - rows and cols must be from 1-3")
		return getCoordinates(turn)
	}



    return []int{rowInt, colInt}
}


func printGameBoard(gameBoard [][]piece) {
	for _, row := range gameBoard {
		fmt.Println(row)
	}
}

func boardFull(gameBoard [][]piece) (bool) {
	for _, row := range gameBoard {
		for _, piece := range row {
			if piece == None {
				return false
			}
		}
	}
	return true
}

func checkRows(board [][]piece) bool {
	for _, row := range board {
		if row[0] != None && row[0] == row[1] && row[1] == row[2] {
			return true
		}
	}
	return false
}

func checkColumns(board [][]piece) bool {
	for col := 0; col < 3; col++ {
		if board[0][col] != None && board[0][col] == board[1][col] && board[1][col] == board[2][col] {
			return true
		}
	}
	return false
}

func checkDiagonals(board [][]piece) bool {
	// Check primary diagonal
	if board[0][0] != None && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	// Check secondary diagonal
	if board[0][2] != None && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	return false
}
