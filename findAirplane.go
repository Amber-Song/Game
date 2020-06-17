package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"time"
)

// AirplaneRoom store the information not change a lot
type AirplaneRoom struct {
	player1      string
	player2      string
	expire       time.Time
	boardLength  int
	airplaneType int
}

// AirplaneGame store the information which change a lot
// and pass into client side
type AirplaneGame struct {
	Board1     [][]int
	Board2     [][]int
	PlayerNow  string
	Round      int
	Win        []string
	ThisPlayer string // thisplayer is not variable that two players should synchronise
}

// AirplaneType different shape of airplane
type AirplaneType struct {
	AirplaneBody int
	AirplaneWing int
	AirplaneTail int
}

func generateHead(boardLength int) (int, int) {
	return rand.Intn(boardLength), rand.Intn(boardLength)
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
func checkAirplane(i int, j int, direction string, boardLength int, airplaneBody int, airplaneWing int) bool {
	wingSide := int(math.Floor(float64(airplaneWing) / 2))
	switch direction {
	case "up":
		if boardLength-1-i < airplaneBody || j < wingSide || boardLength-1-j < wingSide {
			return false
		}
		return true
	case "down":
		if i < airplaneBody || j < wingSide || boardLength-1-j < wingSide {
			return false
		}
		return true
	case "left":
		if boardLength-1-j < airplaneBody || i < wingSide || boardLength-1-i < wingSide {
			return false
		}
		return true
	case "right":
		if j < airplaneBody || i < wingSide || boardLength-1-i < wingSide {
			return false
		}
		return true
	}
	return false
}

// checkAirplaneExisting if there is already an airplane
func checkAirplaneExisting(i int, j int, direction string, board [][]int, boardLength int, airplaneBody int, airplaneWing int, airplaneTail int) bool {
	wingSide := int(math.Floor(float64(airplaneWing) / 2))
	tailSide := int(math.Floor(float64(airplaneTail) / 2))

	if board[i][j] != 0 {
		return false
	}

	switch direction {
	case "up":
		for index := i + 1; index <= i+airplaneBody; index++ {
			if board[index][j] != 0 {
				return false
			}
		}
		for index := j - wingSide; index <= j+wingSide; index++ {
			if board[i+1][index] != 0 {
				return false
			}
		}
		for index := j - tailSide; index <= j+tailSide; index++ {
			if board[i+airplaneBody][index] != 0 {
				return false
			}
		}
		return true
	case "down":
		for index := i - 1; index >= i-airplaneBody; index-- {
			if board[index][j] != 0 {
				return false
			}
		}
		for index := j - wingSide; index <= j+wingSide; index++ {
			if board[i-1][index] != 0 {
				return false
			}
		}
		for index := j - tailSide; index <= j+tailSide; index++ {
			if board[i-airplaneBody][index] != 0 {
				return false
			}
		}
		return true
	case "left":
		for index := j + 1; index <= j+airplaneBody; index++ {
			if board[i][index] != 0 {
				return false
			}
		}
		for index := i - wingSide; index <= i+wingSide; index++ {
			if board[index][j+1] != 0 {
				return false
			}
		}
		for index := i - tailSide; index <= i+tailSide; index++ {
			if board[index][j+airplaneBody] != 0 {
				return false
			}
		}
		return true
	case "right":
		for index := j - 1; index >= j-airplaneBody; index-- {
			if board[i][index] != 0 {
				return false
			}
		}
		for index := i - wingSide; index <= i+wingSide; index++ {
			if board[index][j-1] != 0 {
				return false
			}
		}
		for index := i - tailSide; index <= i+tailSide; index++ {
			if board[index][j-airplaneBody] != 0 {
				return false
			}
		}
		return true
	}
	return false
}

// placeAirplane the airplane on the board
func placeAirplane(i int, j int, direction string, board [][]int, boardLength int, airplaneBody int, airplaneWing int, airplaneTail int) [][]int {
	wingSide := int(math.Floor(float64(airplaneWing) / 2))
	tailSide := int(math.Floor(float64(airplaneTail) / 2))

	board[i][j] = 2

	switch direction {
	case "up":
		for index := i + 1; index <= i+airplaneBody; index++ {
			board[index][j] = 1
		}
		for index := j - wingSide; index <= j+wingSide; index++ {
			board[i+1][index] = 1
		}
		for index := j - tailSide; index <= j+tailSide; index++ {
			board[i+airplaneBody][index] = 1
		}
	case "down":
		for index := i - 1; index >= i-airplaneBody; index-- {
			board[index][j] = 1
		}
		for index := j - wingSide; index <= j+wingSide; index++ {
			board[i-1][index] = 1
		}
		for index := j - tailSide; index <= j+tailSide; index++ {
			board[i-airplaneBody][index] = 1
		}
	case "left":
		for index := j + 1; index <= j+airplaneBody; index++ {
			board[i][index] = 1
		}
		for index := i - wingSide; index <= i+wingSide; index++ {
			board[index][j+1] = 1
		}
		for index := i - tailSide; index <= i+tailSide; index++ {
			board[index][j+airplaneBody] = 1
		}
	case "right":
		for index := j - 1; index >= j-airplaneBody; index-- {
			board[i][index] = 1
		}
		for index := i - wingSide; index <= i+wingSide; index++ {
			board[index][j-1] = 1
		}
		for index := i - tailSide; index <= i+tailSide; index++ {
			board[index][j-airplaneBody] = 1
		}
	}
	return board
}

func generateAirplane(boardLength int, airplaneBody int, airplaneWing int, airplaneTail int) [][]int {
	var board [][]int

	// Initiate board
	for i := 0; i < boardLength; i++ {
		board = append(board, []int{})
		for j := 0; j < boardLength; j++ {
			board[i] = append(board[i], 0)
		}
	}

	// Find place for the first airplane and place it
	for index := 0; ; index++ {
		i, j := generateHead(boardLength)
		direction := generateDirection()
		if checkAirplane(i, j, direction, boardLength, airplaneBody, airplaneWing) {
			board = placeAirplane(i, j, direction, board, boardLength, airplaneBody, airplaneWing, airplaneTail)
			break
		}
	}

	// Find place for the second airplane and place it
	for index := 0; ; index++ {
		i, j := generateHead(boardLength)
		direction := generateDirection()
		if checkAirplane(i, j, direction, boardLength, airplaneBody, airplaneWing) && checkAirplaneExisting(i, j, direction, board, boardLength, airplaneBody, airplaneWing, airplaneTail) {
			board = placeAirplane(i, j, direction, board, boardLength, airplaneBody, airplaneWing, airplaneTail)
			break
		}
	}

	return board
}

func checkWin(board1 [][]int, board2 [][]int, boardLength int) []string {
	player1 := 0
	player2 := 0
	win := []string{}

	for i := 0; i < boardLength; i++ {
		for j := 0; j < boardLength; j++ {
			if board1[i][j] == 5 {
				player1++
			}
		}
	}

	for i := 0; i < boardLength; i++ {
		for j := 0; j < boardLength; j++ {
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

func (game *AirplaneGame) getUpdateGameData(r *http.Request, userNow string, boardLength int) AirplaneGame {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var receiveGame AirplaneGame
	err = json.Unmarshal(reqBody, &receiveGame)
	if err != nil {
		fmt.Println(err)
	}

	var newGameState AirplaneGame
	if userNow == "player1" {
		newGameState = AirplaneGame{Board1: receiveGame.Board1, Board2: game.Board2, PlayerNow: "player2", Round: game.Round, Win: checkWin(receiveGame.Board1, game.Board2, boardLength), ThisPlayer: userNow}
	} else {
		newGameState = AirplaneGame{Board1: game.Board1, Board2: receiveGame.Board2, PlayerNow: "player1", Round: game.Round + 1, Win: checkWin(game.Board1, receiveGame.Board2, boardLength), ThisPlayer: userNow}
	}

	return newGameState
}