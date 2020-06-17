package main

import (
	"encoding/json"
	"errors"
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
var airplaneTypes []AirplaneType

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
		fmt.Println(err) // TODO root file broken
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

func getCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("gameUser")
	if err != nil {
		return "", err
	}
	if cookie == nil {
		return "", errors.New("Cookie didn't set")
	}
	return cookie.Value, nil
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

	boardLength, err := strconv.Atoi(strings.Join(r.URL.Query()["boardLength"], ""))
	if err != nil {
		boardLength = 8
	}
	airplaneType := 0

	user, err := getCookie(r)
	if user == "" {
		if err == errors.New("Cookie didn't set") {
			user = setCookie(w)
		} else {
			fmt.Println(err)
		}
	}
	roomid := setRoom()
	board := generateAirplane(boardLength, airplaneTypes[airplaneType].AirplaneBody, airplaneTypes[airplaneType].AirplaneWing, airplaneTypes[airplaneType].AirplaneTail)

	room := AirplaneRoom{player1: user, player2: "", expire: time.Now().AddDate(0, 0, 1), boardLength: boardLength, airplaneType: airplaneType}
	game := AirplaneGame{Board1: board, Board2: board, PlayerNow: "player1", Round: 1, Win: []string{}, ThisPlayer: ""}
	airplaneRooms[roomid] = room
	airplaneGames[roomid] = game

	cleanRooms()

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
	user, err := getCookie(r)
	if user == "" {
		if err == errors.New("Cookie didn't set") {
			user = setCookie(w)
		} else {
			fmt.Println(err)
		}
	}

	// Check what is the user now and also if is a new user, update room and game
	room := airplaneRooms[roomid]
	game := airplaneGames[roomid]
	userNow := room.getUserNow(user) // this should decide which is this play instead of this player in the struct
	if userNow == "" {
		if room.player2 == "" {
			userNow = "player2"
			board := generateAirplane(room.boardLength, airplaneTypes[room.airplaneType].AirplaneBody, airplaneTypes[room.airplaneType].AirplaneWing, airplaneTypes[room.airplaneType].AirplaneTail)
			updateRoom := AirplaneRoom{player1: room.player1, player2: user, expire: room.expire, boardLength: room.boardLength, airplaneType: room.airplaneType}
			updateGame := AirplaneGame{Board1: game.Board1, Board2: board, PlayerNow: game.PlayerNow, Round: game.Round, Win: game.Win, ThisPlayer: userNow}
			airplaneRooms[roomid] = updateRoom
			airplaneGames[roomid] = updateGame
			room = updateRoom
			game = updateGame
		} else {
			http.Redirect(w, r, "/FindAirplane/Game", http.StatusFound)
			return
		}
	}

	// Check if it is post or get
	if r.Method == http.MethodPost {
		//	if post, get new data, update it and then send the data back
		updateGame := game.getUpdateGameData(r, userNow, room.boardLength)
		airplaneGames[roomid] = updateGame
	} else {
		//	if get, then return data
		updateGame := AirplaneGame{Board1: game.Board1, Board2: game.Board2, PlayerNow: game.PlayerNow, Round: game.Round, Win: game.Win, ThisPlayer: userNow}
		b, err := json.Marshal(updateGame)
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
	airplaneTypes = make([]AirplaneType, 0)
	airplaneTypes = append(airplaneTypes, AirplaneType{AirplaneBody: 3, AirplaneWing: 5, AirplaneTail: 3})

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/FindAirplane/Game", airplaneInitHandler)
	http.HandleFunc("/FindAirplane/Game/room", airplaneGameHandler)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
