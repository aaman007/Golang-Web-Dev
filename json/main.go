package main

import (
	"encoding/json"
	"fmt"
)

type CPSite struct {
	Name string `json:"name"`
}

type User struct {
	Username 	string 		`json:"username"`
	Firstname 	string 		`json:"firstname"`
	Lastname 	string 		`json:"lastname"`
	Country 	string 		`json:"country"`
	CPSites		[]CPSite	`json:"cpSites"`
}

func main() {
	marshalling()
	unmarshalling()
	unmarshallingArray()
}

func marshalling() {
	fmt.Println("Marshalling")

	user := User{
		"aaman007",
		"Amanur",
		"Rahman",
		"Bangladesh",
		[]CPSite{{"Codeforces"}, {"CodeChef"}},
	}
	bs, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

func unmarshalling() {
	fmt.Println("Unmarshalling")

	jsonData := `{"username":"aaman007","firstname":"Amanur","lastname":"Rahman","country":"Bangladesh",
			"cpSites":[{"name":"Codeforces"},{"name":"CodeChef"}]}`

	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
	fmt.Println(user.Username)
	fmt.Println(user.CPSites)
	fmt.Println(user.CPSites[0].Name)
}

func unmarshallingArray() {
	fmt.Println("Unmarshalling Array")

	jsonData := `[{"username":"aaman007","firstname":"Amanur","lastname":"Rahman","country":"Bangladesh",
			"cpSites":[{"name":"Codeforces"},{"name":"CodeChef"}]}, {"username":"Decayed","firstname":"Rahman",
			"lastname":"Aaman","country":"Canada", "cpSites":[{"name":"LeetCode"},{"name":"TopCoder"}]}]`

	var users []User
	err := json.Unmarshal([]byte(jsonData), &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
	for _, user := range users {
		fmt.Println(user)
		fmt.Println(user.Username)
		fmt.Println(user.CPSites)
		fmt.Println(user.CPSites[0].Name)
	}
}
