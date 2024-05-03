package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type TicTacToeBoxRoom struct {
	player1 string
	player2 string
	expire  time.Time
}

type TicTacToeBoxGame struct {
	BoxCollection1 []string
	BoxCollection2 []string
	Board          [][][]string
	BoardPlayer    [][][]string
	ThisPlayer     string
	Player1        string
	Player2        string
	Winner         string
	Round          int
	PlayerNow      string
	Err            string
}

type TicTacToeBoxRound struct {
	BoxCollection1  []string
	BoxCollection2  []string
	Board           [][][]string
	BoardPlayer     [][][]string
	BoardShowPlayer [][]string
}

func generateBoxCollection() []string {
	return []string{"small", "small", "medium", "medium", "large", "large"}
}

// generateBoard generate 2 dimensional box array
func generateBoard() [][][]string {
	var board [][][]string
	for i := 0; i < 3; i++ {
		var line [][]string
		for j := 0; j < 3; j++ {
			cell := make([]string, 0)
			line = append(line, cell)
		}
		board = append(board, line)
	}
	return board
}

func checkTTTWinner(playerBoard [][]string) string {
	if playerBoard[0][0] != "" && playerBoard[0][0] == playerBoard[1][0] && playerBoard[0][0] == playerBoard[2][0] {
		return playerBoard[0][0]
	}
	if playerBoard[0][0] != "" && playerBoard[0][0] == playerBoard[1][1] && playerBoard[0][0] == playerBoard[2][2] {
		return playerBoard[0][0]
	}
	if playerBoard[0][0] != "" && playerBoard[0][0] == playerBoard[0][1] && playerBoard[0][0] == playerBoard[0][2] {
		return playerBoard[0][0]
	}
	if playerBoard[0][1] != "" && playerBoard[0][1] == playerBoard[1][1] && playerBoard[0][1] == playerBoard[2][1] {
		return playerBoard[0][1]
	}
	if playerBoard[0][2] != "" && playerBoard[0][2] == playerBoard[1][1] && playerBoard[0][2] == playerBoard[2][0] {
		return playerBoard[0][2]
	}
	if playerBoard[0][2] != "" && playerBoard[0][2] == playerBoard[1][2] && playerBoard[0][2] == playerBoard[2][2] {
		return playerBoard[0][2]
	}
	if playerBoard[1][0] != "" && playerBoard[1][0] == playerBoard[1][1] && playerBoard[1][0] == playerBoard[1][2] {
		return playerBoard[1][0]
	}
	if playerBoard[2][0] != "" && playerBoard[2][0] == playerBoard[2][1] && playerBoard[2][0] == playerBoard[2][2] {
		return playerBoard[2][0]
	}
	return ""
}

func tictactoeBoxInitHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	// Get the user
	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
	}
	// Get Room
	var room TicTacToeBoxRoom
	room.removeExpiredRoom()
	roomid := generateRoomid("TicTacToeBoxRoom")
	// Set room and game data
	room = TicTacToeBoxRoom{player1: user, player2: "", expire: getExpireTime()}
	game := TicTacToeBoxGame{BoxCollection1: generateBoxCollection(), BoxCollection2: generateBoxCollection(), Board: generateBoard(), BoardPlayer: generateBoard(), ThisPlayer: "player1", Player1: user, Player2: "", Winner: "", Round: 1, PlayerNow: "player1", Err: ""}
	tictactoeRooms[roomid] = room
	tictactoeGames[roomid] = game

	// Return roomid
	b, err := json.Marshal(roomid)
	logError(err)
	w.Write(b)
}

func tictactoeBoxWaitHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
	}

	var room TicTacToeBoxRoom
	roomid := strings.Join(r.URL.Query()["room"], "")
	isExist := isRoomExist(roomid, "TicTacToeBoxRoom")
	if !isExist {
		fmt.Print("sorry! THis room is not existing!")
		game := TicTacToeBoxGame{Err: "Sorry! The room is not existing!"}
		b, err := json.Marshal(game)
		logError(err)
		w.Write(b)
		return
	}

	room = tictactoeRooms[roomid]
	game := tictactoeGames[roomid]
	playerNow := getPlayerNow(user, room.player1, room.player2)

	if playerNow == "" {
		if room.player2 == "" {
			playerNow = "player2"
			updateRoom := TicTacToeBoxRoom{player1: room.player1, player2: user, expire: room.expire}
			updateGame := TicTacToeBoxGame{BoxCollection1: game.BoxCollection1, BoxCollection2: game.BoxCollection2, Board: game.Board, BoardPlayer: game.BoardPlayer, ThisPlayer: playerNow, Player1: game.Player1, Player2: user, Winner: game.Winner, Round: game.Round, PlayerNow: game.PlayerNow, Err: ""}
			tictactoeRooms[roomid] = updateRoom
			tictactoeGames[roomid] = updateGame
			game = updateGame
		} else {
			game := TicTacToeBoxGame{Err: "Sorry! This room is full!"}
			b, err := json.Marshal(game)
			logError(err)
			w.Write(b)
			return
		}
	}

	updateGame := TicTacToeBoxGame{BoxCollection1: game.BoxCollection1, BoxCollection2: game.BoxCollection2, Board: game.Board, BoardPlayer: game.BoardPlayer, ThisPlayer: playerNow, Player1: game.Player1, Player2: game.Player2, Winner: game.Winner, Round: game.Round, PlayerNow: game.PlayerNow, Err: ""}
	b, err := json.Marshal(updateGame)
	logError(err)
	w.Write(b)
}

func tictactoeBoxGameHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	user, _ := getCookie(r)
	if user == "" {
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	var room TicTacToeBoxRoom
	roomid := strings.Join(r.URL.Query()["room"], "")
	isExist := isRoomExist(roomid, "TicTacToeBoxRoom")
	if !isExist {
		game := TicTacToeBoxGame{Err: "Sorry! The room is not existing!"}
		b, err := json.Marshal(game)
		logError(err)
		w.Write(b)
		return
	}

	room = tictactoeRooms[roomid]
	game := tictactoeGames[roomid]
	playerNow := getPlayerNow(user, room.player1, room.player2)
	if playerNow == "" {
		game := TicTacToeBoxGame{Err: "Sorry! This room is full!"}
		b, err := json.Marshal(game)
		logError(err)
		w.Write(b)
		return
	}

	if r.Method == http.MethodPost {
		reqBody, err := io.ReadAll(r.Body)
		logError(err)

		var getPost TicTacToeBoxRound
		err = json.Unmarshal(reqBody, &getPost)
		logError(err)

		// Update game
		var updateGame TicTacToeBoxGame
		if playerNow == "player1" {
			updateGame = TicTacToeBoxGame{BoxCollection1: getPost.BoxCollection1, BoxCollection2: game.BoxCollection2, Board: getPost.Board, BoardPlayer: getPost.BoardPlayer, ThisPlayer: "player1", Player1: game.Player1, Player2: game.Player2, Winner: checkTTTWinner(getPost.BoardShowPlayer), Round: game.Round, PlayerNow: "player2", Err: ""}
		} else {
			updateGame = TicTacToeBoxGame{BoxCollection1: game.BoxCollection1, BoxCollection2: getPost.BoxCollection2, Board: getPost.Board, BoardPlayer: getPost.BoardPlayer, ThisPlayer: "player2", Player1: game.Player1, Player2: game.Player2, Winner: checkTTTWinner(getPost.BoardShowPlayer), Round: game.Round + 1, PlayerNow: "player1", Err: ""}
		}
		tictactoeGames[roomid] = updateGame
	} else {
		// return array
		b, err := json.Marshal(game)
		logError(err)
		w.Write(b)
	}
}
