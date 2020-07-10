package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// RPSRoom store the data which doesn't change a lot
type RPSRoom struct {
	player1 string
	player2 string
	expire  time.Time
}

// RPSGame store the information which change a lot
type RPSGame struct {
	Collection1 []string
	Collection2 []string
	Round       int
	ThisPlayer  string // thisplayer is not variable that two players should synchronise
	Player1     string
	Player2     string
}

// RPSTurn is used to pass between server and client
type RPSTurn struct {
	Card1 string
	Card2 string
}

func returnRPS(i int) string {
	switch i {
	case 0:
		return "rock"
	case 1:
		return "paper"
	case 2:
		return "scissor"
	}
	return ""
}

func generateCardCollection() []string {
	var collection []string
	for i := 0; i < 10; i++ {
		r := rand.Intn(3)
		collection = append(collection, returnRPS(r))
	}
	return collection
}

func (room RPSRoom) getUserNow(user string) string {
	fmt.Println("user: ", user, "room player1: ", room.player1, "player2: ", room.player2)
	switch user {
	case room.player1:
		return "player1"
	case room.player2:
		return "player2"
	default:
		return ""
	}
}

func rpsInitHandler(w http.ResponseWriter, r *http.Request) {
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
	roomid := setRoom()

	collection := generateCardCollection()
	room := RPSRoom{player1: user, player2: "", expire: time.Now().AddDate(0, 0, 1)}
	game := RPSGame{Collection1: collection, Collection2: collection, Round: 1, ThisPlayer: "", Player1: user, Player2: ""}
	rpsRooms[roomid] = room
	rpsGames[roomid] = game

	fmt.Println("Room: ", rpsRooms[roomid])
	fmt.Println("Game: ", rpsGames[roomid])

	b, err := json.Marshal(roomid)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}

func rpsGameHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	// Get room and cookie
	roomid := strings.Join(r.URL.Query()["room"], "")
	check := checkRPSRoom(roomid)
	if !check {
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}
	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
	}

	room := rpsRooms[roomid]
	game := rpsGames[roomid]
	userNow := room.getUserNow(user)
	fmt.Println("Game--- roomid: ", roomid, "; user: ", user, "; usernow: ", userNow)

	if userNow == "" {
		fmt.Println("player2 enter!")
		if room.player2 == "" {
			userNow = "player2"
			collection := generateCardCollection()
			updateRoom := RPSRoom{player1: room.player1, player2: user, expire: room.expire}
			updateGame := RPSGame{Collection1: game.Collection1, Collection2: collection, Round: game.Round, ThisPlayer: userNow, Player1: game.Player1, Player2: user}
			rpsRooms[roomid] = updateRoom
			rpsGames[roomid] = updateGame
			game = updateGame
		} else {
			return
		}
	} else {
		fmt.Println("two user has already came in!")
		updateGame := RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: userNow, Player1: game.Player1, Player2: game.Player2}
		game = updateGame
	}

	fmt.Println("Room: ", rpsRooms[roomid])
	fmt.Println("Game: ", rpsGames[roomid])

	b, err := json.Marshal(game)
	if err != nil {
		println(err)
	}
	w.Write(b)
}
