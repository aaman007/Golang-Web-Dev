package database

import (
	"net/http"
	"time"
)

type User struct {
	Username string
	Firstname string
	Lastname string
	Password []byte
	Role string
}

type Session struct {
	Username string
	LastActive time.Time
}

type Path struct {
	Route string
	HandlerFunction http.HandlerFunc
}
