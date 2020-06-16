package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// var tmpl = template.Must(template.ParseGlob("tmpl/*"))
var airplaneRooms map[string]AirplaneRoom
var airplaneGames map[string]AirplaneGame

// setupResponse() is to deal with CORS Cross Origin Resource Sharing
/* CORS does pre-flight requests sending an OPTIONS request to any URL,
so to handle a POST request you first need to handle an OPTIONS request to that same URL.
*/
func setupResponse(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Root page of all the game
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	b, err := ioutil.ReadFile("./frontend/dist/index.html")
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}

func setCookie(w http.ResponseWriter) string {
	value := strconv.Itoa(rand.Intn(1000000))
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:    "gameUser",
		Value:   value,
		Expires: expire,
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
	return value
}

func getCookie(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("gameUser")
	if err != nil || cookie == nil {
		fmt.Println("err", err, "cookie", cookie)
		return setCookie(w)
	}
	return cookie.Value
}

func setRoom() string {
	for index := 0; ; index++ {
		roomid := strconv.Itoa(rand.Intn(1000000))
		if _, ok := airplaneRooms[roomid]; !ok {
			return roomid
		}
	}
	return ""
}

// cleanRooms remove all the expire room from games list
func cleanRooms() {
	for roomid, room := range airplaneRooms {
		if time.Now().After(room.expire) {
			delete(airplaneRooms, roomid)
			delete(airplaneGames, roomid)
		}
	}
}

func airplaneInitHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	user := getCookie(w, r)
	roomid := setRoom()
	board := generateAirplane()
	fmt.Println("User name is ", user, "Room Id is ", roomid)

	room := AirplaneRoom{player1: user, player2: "", expire: time.Now().AddDate(0, 0, 1)}
	game := AirplaneGame{Board1: board, Board2: board, PlayerNow: "player1", Round: 1, Win: []string{}, ThisPlayer: "player1"}
	airplaneRooms[roomid] = room
	airplaneGames[roomid] = game

	fmt.Println("player 1 enter Game now is ", airplaneGames[roomid])
	cleanRooms()

	// Redirect to the game with roomid
	b, err := json.Marshal(roomid)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}

func airplaneGameHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	// Get room and cookie
	roomid := strings.Join(r.URL.Query()["room"], "")
	user := getCookie(w, r)
	fmt.Println("User name is ", user, "Room Id is ", roomid)

	// Check what is the user now and also if is a new user, update room and game
	room := airplaneRooms[roomid]
	game := airplaneGames[roomid]
	userNow := room.getUserNow(user)
	if userNow == "" {
		if room.player2 == "" {
			userNow = "player2"
			board := generateAirplane()
			newRoom := AirplaneRoom{player1: room.player1, player2: user, expire: room.expire}
			newGame := AirplaneGame{Board1: game.Board1, Board2: board, PlayerNow: game.PlayerNow, Round: game.Round, Win: game.Win, ThisPlayer: userNow}
			airplaneRooms[roomid] = newRoom
			airplaneGames[roomid] = newGame
			room = newRoom
			game = newGame
			fmt.Println("player 2 enter Game now is ", airplaneGames[roomid])

		} else {
			http.Redirect(w, r, "/notFound", http.StatusFound)
			return
		}
	}

	// Check if it is post or get
	if r.Method == http.MethodPost {
		//	if post, get new data, update it and then send the data back
		fmt.Println("Post data")
		newGame := game.getNewGame(r, userNow)
		fmt.Println("get new data", newGame)
		airplaneGames[roomid] = newGame
		fmt.Println("Game now is ", airplaneGames[roomid])
	} else {
		//	if get, then return data
		b, err := json.Marshal(game)
		if err != nil {
			println(err)
		}
		w.Write(b)
	}
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("frontend/dist/js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("frontend/dist/css/"))))

	airplaneGames = make(map[string]AirplaneGame, 0)
	airplaneRooms = make(map[string]AirplaneRoom, 0)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/FindAirplane/Game", airplaneInitHandler)
	http.HandleFunc("/FindAirplane/Game/room", airplaneGameHandler)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
