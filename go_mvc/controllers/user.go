package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/aaman007/Golang-Web-Dev/go_mvc/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"net/http"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(mp map[string]models.User) *UserController {
	return &UserController{mp}
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	userId := p.ByName("id")
	user := uc.session[userId]

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJson)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}
	json.NewDecoder(req.Body).Decode(&user)

	// Generating UUID
	sId, _ := uuid.NewV4()
	user.Id = sId.String()

	// Storing Data to Storage
	uc.session[user.Id] = user
	models.StoreUsers(uc.session)

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userJson)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	userId := p.ByName("id")

	// Delete and Update Storage
	delete(uc.session, userId)
	models.StoreUsers(uc.session)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user ", userId, "\n")
}