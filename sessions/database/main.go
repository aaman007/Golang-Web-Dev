package database

import (
	"time"
)

var DBUsers = map[string]User{}
var DBSessions = map[string]Session{}
var DBCleanedSession time.Time