package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id 		string 	`json:"id"`
	Name 	string 	`json:"name"`
	Gender 	string 	`json:"gender"`
	Age 	int		`json:"age"`
}

func StoreUsers(mp map[string]User) {
	f, err := os.Create("db.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(mp)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadUsers() map[string]User {
	mp := map[string]User{}

	f, err := os.Open("db.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&mp)
	if err != nil {
		fmt.Println(err)
	}
	return mp
}
