package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// AirplaneRoom store the information not change a lot
type AirplaneRoom struct {
	player1 string
	player2 string
	expire  time.Time
}

// AirplaneGame store the information which change a lot
// and pass into client side
type AirplaneGame struct {
	Board1     [8][8]int
	Board2     [8][8]int
	PlayerNow  string
	Round      int
	Win        []string
	ThisPlayer string
}

type ReceiveData struct {
	WithCredentials bool
	Params          map[string]string
	Data            AirplaneGame
}

func generateHead() (int, int) {
	return rand.Intn(8), rand.Intn(8)
}

func generateDirection() string {
	direction := rand.Intn(4)

	switch direction {
	case 0:
		return "up"
	case 1:
		return "down"
	case 2:
		return "left"
	case 3:
		return "right"
	}
	return ""
}

// checkAirplane : check if the space is enough for the airplane
func checkAirplane(i int, j int, direction string) bool {
	switch direction {
	case "up":
		if 7-i < 3 || j < 2 || 7-j < 2 {
			return false
		}
		return true
	case "down":
		if i < 3 || j < 2 || 7-j < 2 {
			return false
		}
		return true
	case "left":
		if 7-j < 3 || i < 2 || 7-i < 2 {
			return false
		}
		return true
	case "right":
		if j < 3 || i < 2 || 7-i < 2 {
			return false
		}
		return true
	}
	return false
}

// checkAirplaneExisting if there is already an airplane
func checkAirplaneExisting(i int, j int, direction string, board [8][8]int) bool {
	if board[i][j] != 0 {
		return false
	}

	switch direction {
	case "up":
		for index := i + 1; index <= i+3; index++ {
			if board[index][j] != 0 {
				return false
			}
		}
		for index := j - 2; index <= j+2; index++ {
			if board[i+1][index] != 0 {
				return false
			}
		}
		for index := j - 1; index <= j+1; index++ {
			if board[i+3][index] != 0 {
				return false
			}
		}
		return true
	case "down":
		for index := i - 1; index >= i-3; index-- {
			if board[index][j] != 0 {
				return false
			}
		}
		for index := j - 2; index <= j+2; index++ {
			if board[i-1][index] != 0 {
				return false
			}
		}
		for index := j - 1; index <= j+1; index++ {
			if board[i-3][index] != 0 {
				return false
			}
		}
		return true
	case "left":
		for index := j + 1; index <= j+3; index++ {
			if board[i][index] != 0 {
				return false
			}
		}
		for index := i - 2; index <= i+2; index++ {
			if board[index][j+1] != 0 {
				return false
			}
		}
		for index := i - 1; index <= i+1; index++ {
			if board[index][j+3] != 0 {
				return false
			}
		}
		return true
	case "right":
		for index := j - 1; index >= j-3; index-- {
			if board[i][index] != 0 {
				return false
			}
		}
		for index := i - 2; index <= i+2; index++ {
			if board[index][j-1] != 0 {
				return false
			}
		}
		for index := i - 1; index <= i+1; index++ {
			if board[index][j-3] != 0 {
				return false
			}
		}
		return true
	}
	return false
}

// placeAirplane the airplane on the board
func placeAirplane(i int, j int, direction string, board [8][8]int) [8][8]int {
	board[i][j] = 2

	switch direction {
	case "up":
		for index := i + 1; index <= i+3; index++ {
			board[index][j] = 1
		}
		for index := j - 2; index <= j+2; index++ {
			board[i+1][index] = 1
		}
		for index := j - 1; index <= j+1; index++ {
			board[i+3][index] = 1
		}
	case "down":
		for index := i - 1; index >= i-3; index-- {
			board[index][j] = 1
		}
		for index := j - 2; index <= j+2; index++ {
			board[i-1][index] = 1
		}
		for index := j - 1; index <= j+1; index++ {
			board[i-3][index] = 1
		}
	case "left":
		for index := j + 1; index <= j+3; index++ {
			board[i][index] = 1
		}
		for index := i - 2; index <= i+2; index++ {
			board[index][j+1] = 1
		}
		for index := i - 1; index <= i+1; index++ {
			board[index][j+3] = 1
		}
	case "right":
		for index := j - 1; index >= j-3; index-- {
			board[i][index] = 1
		}
		for index := i - 2; index <= i+2; index++ {
			board[index][j-1] = 1
		}
		for index := i - 1; index <= i+1; index++ {
			board[index][j-3] = 1
		}
	}
	return board
}

func generateAirplane() [8][8]int {
	var board [8][8]int

	// Initiate board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = 0
		}
	}

	// Find place for the first airplane and place it
	for index := 0; ; index++ {
		i, j := generateHead()
		direction := generateDirection()
		if checkAirplane(i, j, direction) {
			board = placeAirplane(i, j, direction, board)
			break
		}
	}

	// Find place for the second airplane and place it
	for index := 0; ; index++ {
		i, j := generateHead()
		direction := generateDirection()
		if checkAirplane(i, j, direction) && checkAirplaneExisting(i, j, direction, board) {
			board = placeAirplane(i, j, direction, board)
			break
		}
	}

	return board
}

func checkWin(board1 [8][8]int, board2 [8][8]int) []string {
	player1 := 0
	player2 := 0
	win := []string{}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board1[i][j] == 5 {
				player1++
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board2[i][j] == 5 {
				player2++
			}
		}
	}

	if player1 == 2 {
		win = append(win, "player1")
	}
	if player2 == 2 {
		win = append(win, "player2")
	}

	return win
}

func (room *AirplaneRoom) getUserNow(user string) string {
	switch user {
	case room.player1:
		return "player1"
	case room.player2:
		return "player2"
	default:
		return ""
	}
}

func (game *AirplaneGame) getNewGame(r *http.Request, userNow string) AirplaneGame {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var receiveData ReceiveData
	err = json.Unmarshal(reqBody, &receiveData)

	receiveGame := receiveData.Data
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("original game is ", game)
	fmt.Println("receive game is ", receiveGame)

	var newGame AirplaneGame
	if userNow == "player1" {
		newGame = AirplaneGame{Board1: receiveGame.Board1, Board2: game.Board2, PlayerNow: "player2", Round: game.Round, Win: checkWin(receiveGame.Board1, game.Board2), ThisPlayer: userNow}
	} else {
		newGame = AirplaneGame{Board1: game.Board1, Board2: receiveGame.Board2, PlayerNow: "player1", Round: game.Round + 1, Win: checkWin(game.Board1, receiveGame.Board2), ThisPlayer: userNow}
	}

	return newGame
}
