package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// AirplaneRoom store the information not change a lot
type AirplaneRoom struct {
	player1     string
	player2     string
	expire      time.Time
	boardLength int
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
	Player1    string
	Player2    string
	Err        string
}

type AirplaneTypeA struct {
}

type Airplane interface {
	isSpaceAvail(i int, j int, direction string, boardLength int) bool
	isAirplaneNotOverlap(i int, j int, direction string, board [][]int) bool
	placeAirplane(i int, j int, direction string, board [][]int) [][]int
	generateAirplane(boardLength int) [][]int
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

func (airplane AirplaneTypeA) isSpaceAvail(i int, j int, direction string, boardLength int) bool {
	switch direction {
	case "up":
		if boardLength-i <= 3 || j < 2 || boardLength-j <= 2 {
			return false
		}
		return true
	case "down":
		if i < 3 || j < 2 || boardLength-j <= 2 {
			return false
		}
		return true
	case "left":
		if boardLength-j <= 3 || i < 2 || boardLength-i <= 2 {
			return false
		}
		return true
	case "right":
		if j < 3 || i < 2 || boardLength-i <= 2 {
			return false
		}
		return true
	}
	return false
}

func (airplane AirplaneTypeA) isAirplaneNotOverlap(i int, j int, direction string, board [][]int) bool {
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

func (airplane AirplaneTypeA) placeAirplane(i int, j int, direction string, board [][]int) [][]int {
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

func (airplane AirplaneTypeA) generateAirplane(boardLength int) [][]int {
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
		if airplane.isSpaceAvail(i, j, direction, boardLength) {
			board = airplane.placeAirplane(i, j, direction, board)
			break
		}
	}

	// Find place for the second airplane and place it
	for index := 0; ; index++ {
		i, j := generateHead(boardLength)
		direction := generateDirection()
		if airplane.isSpaceAvail(i, j, direction, boardLength) && airplane.isAirplaneNotOverlap(i, j, direction, board) {
			board = airplane.placeAirplane(i, j, direction, board)
			break
		}
	}

	return board
}

func checkFAWinner(board1 [][]int, board2 [][]int, boardLength int) []string {
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

func airplaneInitHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	boardLength, err := strconv.Atoi(strings.Join(r.URL.Query()["boardLength"], ""))
	if err != nil {
		boardLength = 8
	}

	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
	}
	var room AirplaneRoom
	room.removeExpiredRoom()
	roomid := room.generateRoomid()
	var airplane AirplaneTypeA
	board := airplane.generateAirplane(boardLength)

	room = AirplaneRoom{player1: user, player2: "", expire: getExpireTime(), boardLength: boardLength}
	game := AirplaneGame{Board1: board, Board2: board, PlayerNow: "player1", Round: 1, Win: []string{}, ThisPlayer: "", Player1: user, Player2: "", Err: ""}
	airplaneRooms[roomid] = room
	airplaneGames[roomid] = game

	b, err := json.Marshal(roomid)
	logError(err)
	w.Write(b)
}

func airplaneGameHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	// Get room and cookie
	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
	}
	var room AirplaneRoom
	roomid := strings.Join(r.URL.Query()["room"], "")
	isExist := room.isRoomExist(roomid)
	if !isExist {
		// TODO can be written as game interface?
		game := AirplaneGame{Err: "Sorry! The room is not existing!"}
		b, err := json.Marshal(game)
		logError(err)
		w.Write(b)
		return
	}

	// Check what is the user now and also if is a new user, update room and game
	room = airplaneRooms[roomid]
	game := airplaneGames[roomid]
	playerNow := room.getPlayerNow(user) // this should decide which is this play instead of this player in the struct
	if playerNow == "" {
		if room.player2 == "" {
			playerNow = "player2"
			var airplane AirplaneTypeA
			board := airplane.generateAirplane(room.boardLength)
			updateRoom := AirplaneRoom{player1: room.player1, player2: user, expire: room.expire, boardLength: room.boardLength}
			updateGame := AirplaneGame{Board1: game.Board1, Board2: board, PlayerNow: game.PlayerNow, Round: game.Round, Win: game.Win, ThisPlayer: playerNow, Player1: room.player1, Player2: user, Err: ""}
			airplaneRooms[roomid] = updateRoom
			airplaneGames[roomid] = updateGame
			room = updateRoom
			game = updateGame
		} else {
			game := AirplaneGame{Err: "Sorry! This room is full!"}
			b, err := json.Marshal(game)
			logError(err)
			w.Write(b)
			return
		}
	}

	// Check if it is post or get
	if r.Method == http.MethodPost {
		//	if post, get new data, update it and then send the data back
		reqBody, err := ioutil.ReadAll(r.Body)
		logError(err)
		var receiveGame AirplaneGame
		err = json.Unmarshal(reqBody, &receiveGame)
		logError(err)

		var newGameState AirplaneGame
		if playerNow == "player1" {
			newGameState = AirplaneGame{Board1: receiveGame.Board1, Board2: game.Board2, PlayerNow: "player2", Round: game.Round, Win: checkFAWinner(receiveGame.Board1, game.Board2, room.boardLength), ThisPlayer: playerNow, Player1: game.Player1, Player2: game.Player2, Err: ""}
		} else {
			newGameState = AirplaneGame{Board1: game.Board1, Board2: receiveGame.Board2, PlayerNow: "player1", Round: game.Round + 1, Win: checkFAWinner(game.Board1, receiveGame.Board2, room.boardLength), ThisPlayer: playerNow, Player1: game.Player1, Player2: game.Player2, Err: ""}
		}
		airplaneGames[roomid] = newGameState
	} else {
		//	if get, then return data
		updateGame := AirplaneGame{Board1: game.Board1, Board2: game.Board2, PlayerNow: game.PlayerNow, Round: game.Round, Win: game.Win, ThisPlayer: playerNow, Player1: game.Player1, Player2: game.Player2, Err: ""}
		b, err := json.Marshal(updateGame)
		logError(err)
		w.Write(b)
	}
}

// TODO restart
// func airplaneRestartHandler(w http.ResponseWriter, r *http.Request) {
// 	// Enabling CORS on a Go Web Server
// 	setupResponse(w)
// 	if r.Method == "OPTIONS" {
// 		return
// 	}

// 	// Get room and cookie
// 	roomid := strings.Join(r.URL.Query()["room"], "")
//  isExist := room.isRoomExist(roomid)
// 	if !isExist {
// 		http.Redirect(w, r, "/NotFound", http.StatusFound)
// 		return
// 	}
// 	fmt.Println("restart! get roomid ", roomid)

// 	room := airplaneRooms[roomid]
// 	board1 := generateAirplane(room.boardLength, airplaneTypes[room.airplaneType].AirplaneBody, airplaneTypes[room.airplaneType].AirplaneWing, airplaneTypes[room.airplaneType].AirplaneTail)
// 	board2 := generateAirplane(room.boardLength, airplaneTypes[room.airplaneType].AirplaneBody, airplaneTypes[room.airplaneType].AirplaneWing, airplaneTypes[room.airplaneType].AirplaneTail)
// 	newGame := AirplaneGame{Board1: board1, Board2: board2, PlayerNow: "player1", Round: 1, Win: []string{}, ThisPlayer: "", Player1: room.player1, Player2: room.player2, Err:""}
// 	airplaneGames[roomid] = newGame
// 	fmt.Println("data changed ", airplaneGames[roomid])
// }
