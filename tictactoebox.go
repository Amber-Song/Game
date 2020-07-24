package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

type TicTacToeBoxRound struct {
	BoxCollection1  []string
	BoxCollection2  []string
	Board           [][][]string
	BoardPlayer     [][][]string
	BoardShowPlayer [][]string
}

func generateBox() []string {
	return []string{"small", "small", "medium", "medium", "large", "large"}
}

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

func (room TicTacToeBoxRoom) getUserNow(user string) string {
	switch user {
	case room.player1:
		return "player1"
	case room.player2:
		return "player2"
	default:
		return ""
	}
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

	cleanRooms()

	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
	}
	roomid := setTictactoeRoom()

	room := TicTacToeBoxRoom{player1: user, player2: "", expire: time.Now().AddDate(0, 0, 1)}
	game := TicTacToeBoxGame{BoxCollection1: generateBox(), BoxCollection2: generateBox(), Board: generateBoard(), BoardPlayer: generateBoard(), ThisPlayer: "player1", Player1: user, Player2: "", Winner: "", Round: 1, PlayerNow: "player1"}
	tictactoeRooms[roomid] = room
	tictactoeGames[roomid] = game

	b, err := json.Marshal(roomid)
	if err != nil {
		fmt.Println(err)
	}
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
	roomid := strings.Join(r.URL.Query()["room"], "")
	check := checkTictactoeRoom(roomid)
	if !check {
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	room := tictactoeRooms[roomid]
	game := tictactoeGames[roomid]
	userNow := room.getUserNow(user)

	if userNow == "" {
		if room.player2 == "" {
			userNow = "player2"
			updateRoom := TicTacToeBoxRoom{player1: room.player1, player2: user, expire: room.expire}
			updateGame := TicTacToeBoxGame{BoxCollection1: game.BoxCollection1, BoxCollection2: game.BoxCollection2, Board: game.Board, BoardPlayer: game.BoardPlayer, ThisPlayer: userNow, Player1: game.Player1, Player2: user, Winner: game.Winner, Round: game.Round, PlayerNow: game.PlayerNow}
			tictactoeRooms[roomid] = updateRoom
			tictactoeGames[roomid] = updateGame
			game = updateGame
		} else {
			http.Redirect(w, r, "/NotFound", http.StatusFound)
			return
		}
	}

	updateGame := TicTacToeBoxGame{BoxCollection1: game.BoxCollection1, BoxCollection2: game.BoxCollection2, Board: game.Board, BoardPlayer: game.BoardPlayer, ThisPlayer: userNow, Player1: game.Player1, Player2: game.Player2, Winner: game.Winner, Round: game.Round, PlayerNow: game.PlayerNow}
	b, err := json.Marshal(updateGame)
	if err != nil {
		fmt.Println(err)
	}
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
	roomid := strings.Join(r.URL.Query()["room"], "")
	check := checkTictactoeRoom(roomid)
	if !check {
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	room := tictactoeRooms[roomid]
	game := tictactoeGames[roomid]
	userNow := room.getUserNow(user)

	if r.Method == http.MethodPost {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		var getPost TicTacToeBoxRound
		err = json.Unmarshal(reqBody, &getPost)
		if err != nil {
			fmt.Println(err)
		}

		// Update game
		var updateGame TicTacToeBoxGame
		if userNow == "player1" {
			updateGame = TicTacToeBoxGame{BoxCollection1: getPost.BoxCollection1, BoxCollection2: game.BoxCollection2, Board: getPost.Board, BoardPlayer: getPost.BoardPlayer, ThisPlayer: "player1", Player1: game.Player1, Player2: game.Player2, Winner: checkTTTWinner(getPost.BoardShowPlayer), Round: game.Round, PlayerNow: "player2"}
		} else {
			updateGame = TicTacToeBoxGame{BoxCollection1: game.BoxCollection1, BoxCollection2: getPost.BoxCollection2, Board: getPost.Board, BoardPlayer: getPost.BoardPlayer, ThisPlayer: "player2", Player1: game.Player1, Player2: game.Player2, Winner: checkTTTWinner(getPost.BoardShowPlayer), Round: game.Round + 1, PlayerNow: "player1"}
		}
		tictactoeGames[roomid] = updateGame
	} else {
		// return array
		b, err := json.Marshal(game)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(b)
	}
}
