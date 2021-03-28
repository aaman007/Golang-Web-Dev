package session

import (
	"../database"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"github.com/satori/go.uuid"
)

type UID string
const SLength int = 30

func getAlphabets(allCaps bool) string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	if allCaps {
		alphabets = strings.ToUpper(alphabets)
	}
	return alphabets
}

func getNumericsAndSpecialChars() string {
	return "0123456789$@!"
}

func uidGenerator(len int) UID {
	alphabets := getAlphabets(false) + getAlphabets(true) + getNumericsAndSpecialChars()
	res := ""
	for i:=0; i<len; i++ {
		res += string(alphabets[rand.Intn(65)])
	}
	return UID(res)
}

func NewV4() (UID, error) {
	return uidGenerator(25), nil
}

func (uid UID) String() string {
	return string(uid)
}

func getOrCreateSessionCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session")
	if err != nil {
		sid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",
			Value: sid.String(),
		}
	}
	cookie.MaxAge = 30
	http.SetCookie(w, cookie)
	return cookie
}

func CreateSessionCookie(w http.ResponseWriter, req *http.Request, username string) *http.Cookie {
	cookie := getOrCreateSessionCookie(w, req)
	database.DBSessions[cookie.Value] = database.Session{
		Username: username,
		LastActive: time.Now(),
	}
	return cookie
}

func GetUser(w http.ResponseWriter, req *http.Request) database.User {
	cookie := getOrCreateSessionCookie(w, req)

	// If the user already exits, get the user
	var user database.User
	if _session, ok := database.DBSessions[cookie.Value]; ok {
		_session.LastActive = time.Now()
		database.DBSessions[cookie.Value] = _session
		user = database.DBUsers[_session.Username]
	}
	return user
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	_session, ok := database.DBSessions[cookie.Value]
	if ok {
		_session.LastActive = time.Now()
		database.DBSessions[cookie.Value] = _session
	}

	_, ok = database.DBUsers[_session.Username]

	// refresh session
	cookie.MaxAge = SLength
	http.SetCookie(w, cookie)
	return ok
}

func CleanSessions() {
	fmt.Println("Cleaning Starting")
	ShowSessions()

	for key, val := range database.DBSessions {
		if time.Now().Sub(val.LastActive) > (time.Second * 30) {
			delete(database.DBSessions, key)
		}
	}

	database.DBCleanedSession = time.Now()
	fmt.Println("Cleaned Expired Sessions")
	ShowSessions()
}

func ShowSessions() {
	fmt.Println("----------START----------")
	for key, val := range database.DBSessions {
		fmt.Println(key, val.Username)
	}
	fmt.Println("-------END---------")
}