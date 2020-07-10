package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// var tmpl = template.Must(template.ParseGlob("tmpl/*"))
var airplaneRooms map[string]AirplaneRoom
var airplaneGames map[string]AirplaneGame
var airplaneTypes []AirplaneType

var rpsRooms map[string]RPSRoom
var rpsGames map[string]RPSGame

// ErrCookieNotSet is a error indicate that cookie is not set
var ErrCookieNotSet error = errors.New("Cookie didn't set")

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

	b, err := ioutil.ReadFile("frontend/dist/index.html")
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
		return "", ErrCookieNotSet
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

func checkAirplaneRoom(roomid string) bool {
	if _, ok := airplaneRooms[roomid]; ok {
		return true
	}
	return false
}

func checkRPSRoom(roomid string) bool {
	if _, ok := rpsRooms[roomid]; ok {
		return true
	}
	return false
}

// cleanRooms remove all the expire room from games list
func cleanRooms() {
	for roomid, room := range airplaneRooms {
		if time.Now().After(room.expire) {
			delete(airplaneRooms, roomid)
			delete(airplaneGames, roomid)
		}
	}

	for roomid, room := range rpsRooms {
		if time.Now().After(room.expire) {
			delete(rpsRooms, roomid)
			delete(rpsGames, roomid)
		}
	}
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("frontend/dist/js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("frontend/dist/css/"))))

	// Airplane
	airplaneRooms = make(map[string]AirplaneRoom, 0)
	airplaneGames = make(map[string]AirplaneGame, 0)
	airplaneTypes = make([]AirplaneType, 0)
	airplaneTypes = append(airplaneTypes, AirplaneType{AirplaneBody: 3, AirplaneWing: 5, AirplaneTail: 3})
	// Rock paper scissor
	rpsRooms = make(map[string]RPSRoom, 0)
	rpsGames = make(map[string]RPSGame, 0)

	// Airplane
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/FindAirplane/Game", airplaneInitHandler)
	http.HandleFunc("/api/FindAirplane/Game/room", airplaneGameHandler)
	// http.HandleFunc("/api/FindAirplane/restart", airplaneRestartHandler)
	// Rock paper scissor
	http.HandleFunc("/api/RockPaperScissor/Game", rpsInitHandler)
	http.HandleFunc("/api/RockPaperScissor/Game/room", rpsGameHandler)

	fmt.Println(http.ListenAndServe(":3000", nil))
}
