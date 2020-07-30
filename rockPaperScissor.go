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

// RPSRoom store the room data which doesn't change a lot
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
	UpdateGame  bool
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
	for i := 0; i < 3; i++ {
		r := rand.Intn(3)
		collection = append(collection, returnRPS(r))
	}
	return collection
}

func getRPSRoundWinner(card1 string, card2 string) int {
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

	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
	}

	var room RPSRoom
	room.removeExpiredRoom()
	roomid := room.generateRoomid()

	collection := generateCardCollection()
	room = RPSRoom{player1: user, player2: "", expire: getExpireTime()}
	game := RPSGame{Collection1: collection, Collection2: collection, Round: 1, ThisPlayer: "player1", Player1: user, Player2: "", Card1: "", Card2: "", Card1Index: -1, Card2Index: -1, UpdateGame: false}
	rpsRooms[roomid] = room
	rpsGames[roomid] = game

	b, err := json.Marshal(roomid)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}

// This function return game data during the time player1 keep query and player2 haven't enter the room
func rpsWaitForPlayer2Handler(w http.ResponseWriter, r *http.Request) {
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
	var room RPSRoom
	roomid := strings.Join(r.URL.Query()["room"], "")
	isExist := room.isRoomExist(roomid)
	if !isExist {
		return
	}

	room = rpsRooms[roomid]
	game := rpsGames[roomid]
	playerNow := room.getPlayerNow(user)

	// Check if it is player2
	if playerNow == "" {
		if room.player2 == "" {
			playerNow = "player2"
			collection := generateCardCollection()
			updateRoom := RPSRoom{player1: room.player1, player2: user, expire: room.expire}
			updateGame := RPSGame{Collection1: game.Collection1, Collection2: collection, Round: game.Round, ThisPlayer: playerNow, Player1: game.Player1, Player2: user, Card1: game.Card1, Card2: game.Card2, Card1Index: game.Card1Index, Card2Index: game.Card2Index, UpdateGame: game.UpdateGame}
			rpsRooms[roomid] = updateRoom
			rpsGames[roomid] = updateGame
			game = updateGame
		} else {
			http.Redirect(w, r, "/Game/api/RockPaperScissor/Game", http.StatusFound)
			return
		}
	}

	updateGame := RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: playerNow, Player1: game.Player1, Player2: game.Player2, Card1: game.Card1, Card2: game.Card2, Card1Index: game.Card1Index, Card2Index: game.Card2Index, UpdateGame: game.UpdateGame}
	// Return data
	b, err := json.Marshal(updateGame)
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
	user, _ := getCookie(r)
	if user == "" {
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	var room RPSRoom
	roomid := strings.Join(r.URL.Query()["room"], "")
	isExist := room.isRoomExist(roomid)
	if !isExist {
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	room = rpsRooms[roomid]
	game := rpsGames[roomid]
	playerNow := room.getPlayerNow(user)

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
		if playerNow == "player1" {
			updateGame = RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: "player1", Player1: game.Player1, Player2: game.Player2, Card1: getPostCard.Card1, Card2: game.Card2, Card1Index: getPostCard.Card1Index, Card2Index: game.Card2Index, UpdateGame: false}
		} else {
			updateGame = RPSGame{Collection1: game.Collection1, Collection2: game.Collection2, Round: game.Round, ThisPlayer: "player2", Player1: game.Player1, Player2: game.Player2, Card1: game.Card1, Card2: getPostCard.Card2, Card1Index: game.Card1Index, Card2Index: getPostCard.Card2Index, UpdateGame: false}
		}
		rpsGames[roomid] = updateGame

	} else {
		// return array
		b, err := json.Marshal(game)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(b)
	}
}

func rpsRoundHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	// Get room and cookie
	user, _ := getCookie(r)
	if user == "" {
		user = setCookie(w)
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	var room RPSRoom
	roomid := strings.Join(r.URL.Query()["room"], "")
	isExist := room.isRoomExist(roomid)
	if !isExist {
		http.Redirect(w, r, "/NotFound", http.StatusFound)
		return
	}

	room = rpsRooms[roomid]
	game := rpsGames[roomid]
	playerNow := room.getPlayerNow(user)

	if !game.UpdateGame {
		card1 := game.Card1
		card2 := game.Card2
		winner := getRPSRoundWinner(card1, card2)
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
		updateGame := RPSGame{Collection1: collection1, Collection2: collection2, Round: game.Round + 1, ThisPlayer: playerNow, Player1: game.Player1, Player2: game.Player2, Card1: "", Card2: "", Card1Index: -1, Card2Index: -1, UpdateGame: true}
		rpsGames[roomid] = updateGame
		game = updateGame
	}

	b, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}
