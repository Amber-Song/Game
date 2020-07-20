package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Card1       string
	Card2       string
	Card1Index  int
	Card2Index  int
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

func checkRoundWinner(card1 string, card2 string) int {
	if card1 == card2 {
		return 0
	} else {
		switch card1 {
		case "rock":
			switch card2 {
			case "paper":
				return 2
			case "scissor":
				return 1
			default:
				return 1
			}
		case "paper":
			switch card2 {
			case "rock":
				return 1
			case "scissor":
				return 2
			default:
				return 1
			}
		case "scissor":
			switch card2 {
			case "rock":
				return 2
			case "paper":
				return 1
			default:
				return 1
			}
		default:
			return 2
		}
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
	game := RPSGame{Collection1: collection, Collection2: collection, Round: 1, ThisPlayer: "", Player1: user, Player2: "", Card1: "", Card2: "", Card1Index: -1, Card2Index: -1}
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
			updateGame := RPSGame{Collection1: game.Collection1, Collection2: collection, Round: game.Round, ThisPlayer: userNow, Player1: game.Player1, Player2: user, Card1: game.Card1, Card2: game.Card2, Card1Index: game.Card1Index, Card2Index: game.Card2Index}
			rpsRooms[roomid] = updateRoom
			rpsGames[roomid] = updateGame
			game = updateGame
		} else {
			http.Redirect(w, r, "/api/RockPaperScissor/Game", http.StatusFound)
			return
		}
	}

	if r.Method == http.MethodPost {
		// Get the chosen card
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		var getPostCard RPSGame
		err = json.Unmarshal(reqBody, &getPostCard)
		if err != nil {
			fmt.Println(err)
		}

		// update game
		var updateGame RPSGame
		if userNow == "player1" {
			updateGame = RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: "player1", Player1: game.Player1, Player2: game.Player2, Card1: getPostCard.Card1, Card2: game.Card2, Card1Index: getPostCard.Card1Index, Card2Index: game.Card2Index}
		} else {
			updateGame = RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: "player2", Player1: game.Player1, Player2: game.Player2, Card1: game.Card1, Card2: getPostCard.Card2, Card1Index: game.Card1Index, Card2Index: getPostCard.Card2Index}
		}
		game = updateGame
		rpsGames[roomid] = updateGame

	} else {
		// return array
		updateGame := RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: userNow, Player1: game.Player1, Player2: game.Player2, Card1: game.Card1, Card2: game.Card2, Card1Index: game.Card1Index, Card2Index: game.Card2Index}
		b, err := json.Marshal(updateGame)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(b)
	}

	fmt.Println("Room: ", rpsRooms[roomid])
	fmt.Println("Game: ", rpsGames[roomid])
}

func rpsRoundHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	room := rpsRooms[roomid]
	userNow := room.getUserNow(user)
	game := rpsGames[roomid]
	var updateGame RPSGame

	if game.Card1 != "" {
		card1 := game.Card1
		card2 := game.Card2
		winner := checkRoundWinner(card1, card2)
		card1Index := game.Card1Index
		card2Index := game.Card2Index
		collection1 := game.Collection1
		collection2 := game.Collection2
		switch winner {
		case 1:
			collection1 = append(collection1[:card1Index], collection1[card1Index+1:]...)
			collection2 = append(collection2, card1)
		case 2:
			collection2 = append(collection2[:card2Index], collection2[card2Index+1:]...)
			collection1 = append(collection1, card2)
		default:
		}
		updateGame = RPSGame{Collection1: collection1, Collection2: collection2, Round: game.Round + 1, ThisPlayer: userNow, Player1: game.Player1, Player2: game.Player2, Card1: "", Card2: "", Card1Index: -1, Card2Index: -1}
		rpsGames[roomid] = updateGame
	} else {
		updateGame = RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: userNow, Player1: game.Player1, Player2: game.Player2, Card1: game.Card1, Card2: game.Card2, Card1Index: game.Card1Index, Card2Index: game.Card2Index}
	}

	fmt.Println("Round end!", updateGame)
	b, err := json.Marshal(updateGame)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}
