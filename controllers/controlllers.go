package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-vote/dto"
	"user-vote/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	u := service.GetAllUsers(db)
	json.NewEncoder(w).Encode(u)
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	var newUser dto.User
	json.NewDecoder(r.Body).Decode(&newUser)
	service.CreateUser(newUser, db)
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	var newUser dto.User
	json.NewDecoder(r.Body).Decode(&newUser)
	service.UpdateUser(newUser, db)
	json.NewEncoder(w).Encode(newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	vars := mux.Vars(r)
	id := vars["id"]
	service.DeleteUser(id, db)
	//json.NewEncoder(w).Encode(newUser)
}

func Balance(w http.ResponseWriter, r *http.Request, client *ethclient.Client) {
	vars := mux.Vars(r)
	key := vars["key"] //id owner
	fmt.Fprint(w, "Home Page")
}

func Transfer(w http.ResponseWriter, r *http.Request, client *ethclient.Client) {
	//usar DTO transfer com address, value, idVenda,
	fmt.Fprint(w, "Home Page")
}

//criar método criar pedido
