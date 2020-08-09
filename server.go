package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var airplaneRooms map[string]AirplaneRoom
var airplaneGames map[string]AirplaneGame

var rpsRooms map[string]RPSRoom
var rpsGames map[string]RPSGame

var tictactoeRooms map[string]TicTacToeBoxRoom
var tictactoeGames map[string]TicTacToeBoxGame

var CookieNotSet error = errors.New("Cookie didn't set")

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

func isRoomExist(roomid string, roomName string) bool {
	switch roomName {
	case "AirplaneRoom":
		if _, ok := airplaneRooms[roomid]; ok {
			return true
		}
		return false
	case "RPSRoom":
		if _, ok := rpsRooms[roomid]; ok {
			return true
		}
		return false
	case "TicTacToeBoxRoom":
		if _, ok := tictactoeRooms[roomid]; ok {
			return true
		}
		return false
	}

	return false
}

func generateRoomid(roomName string) string {
	for index := 0; ; index++ {
		roomid := getRandomInt(1000000)
		if !isRoomExist(roomid, roomName) {
			return roomid
		}
	}
}

func (room AirplaneRoom) removeExpiredRoom() {
	for roomid, room := range airplaneRooms {
		if time.Now().After(room.expire) {
			delete(airplaneRooms, roomid)
			delete(airplaneGames, roomid)
		}
	}
}
func (room RPSRoom) removeExpiredRoom() {
	for roomid, room := range rpsRooms {
		if time.Now().After(room.expire) {
			delete(rpsRooms, roomid)
			delete(rpsGames, roomid)
		}
	}
}
func (room TicTacToeBoxRoom) removeExpiredRoom() {
	for roomid, room := range tictactoeRooms {
		if time.Now().After(room.expire) {
			delete(tictactoeRooms, roomid)
			delete(tictactoeGames, roomid)
		}
	}
}

func getPlayerNow(user string, player1 string, player2 string) string {
	switch user {
	case player1:
		return "player1"
	case player2:
		return "player2"
	default:
		return ""
	}
}

func setCookie(w http.ResponseWriter) string {
	value := getRandomInt(1000000)
	expire := getExpireTime()
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
		return "", CookieNotSet
	}
	return cookie.Value, nil
}

// Return a random number in the string type
func getRandomInt(i int) string {
	return strconv.Itoa(rand.Intn(i))
}

// Send one day later as an expire day
func getExpireTime() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

func logError(err error) {
	if err != nil {
		log.Println("ERROR:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS on a Go Web Server
	setupResponse(w)
	if r.Method == "OPTIONS" {
		return
	}

	b, err := ioutil.ReadFile("frontend/dist/index.html")
	logError(err)
	w.Write(b)
}

func main() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("frontend/dist/js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("frontend/dist/css/"))))

	// Airplane
	airplaneRooms = make(map[string]AirplaneRoom, 0)
	airplaneGames = make(map[string]AirplaneGame, 0)
	// Rock paper scissor
	rpsRooms = make(map[string]RPSRoom, 0)
	rpsGames = make(map[string]RPSGame, 0)
	// TicTacToe
	tictactoeRooms = make(map[string]TicTacToeBoxRoom, 0)
	tictactoeGames = make(map[string]TicTacToeBoxGame, 0)

	// Index
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Game", indexHandler)
	// Airplane
	http.HandleFunc("/Game/api/FindAirplane/Game", airplaneInitHandler)
	http.HandleFunc("/Game/api/FindAirplane/Game/room", airplaneGameHandler)
	// http.HandleFunc("/api/FindAirplane/restart", airplaneRestartHandler)
	// Rock paper scissor
	http.HandleFunc("/Game/api/RockPaperScissor/Game", rpsInitHandler)
	http.HandleFunc("/Game/api/RockPaperScissor/Game/wait", rpsWaitForPlayer2Handler)
	http.HandleFunc("/Game/api/RockPaperScissor/Game/room", rpsGameHandler)
	http.HandleFunc("/Game/api/RockPaperScissor/Game/roundend", rpsRoundHandler)
	// TicTacToe
	http.HandleFunc("/Game/api/TicTacToeBox/Game", tictactoeBoxInitHandler)
	http.HandleFunc("/Game/api/TicTacToeBox/Game/wait", tictactoeBoxWaitHandler)
	http.HandleFunc("/Game/api/TicTacToeBox/Game/room", tictactoeBoxGameHandler)

	fmt.Println(http.ListenAndServe(":3000", nil))
}
