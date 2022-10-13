package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"user-vote/dto"
	"user-vote/kafka"

	"user-vote/service"

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
}

func Balance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	balance := service.Balance(key)
	json.NewEncoder(w).Encode(balance)
}

func Transfer(w http.ResponseWriter, r *http.Request, producer kafka.KafkaProducer) {
	var pay dto.Pay

	json.NewDecoder(r.Body).Decode(&pay)
	//payment = dto.Payment{KeySender: pay.KeySender, Recipient: pay.Recipient, ValuePay: pay.ValuePay}
	isPay := true //service.Transfer(payment)
	order := new(dto.Order)
	if isPay {
		order.IdGame = pay.IdGame
		order.IdUser = pay.IdUser
	} else {
		order.Erro = "Error in transaction"
	}
	transactionJson, _ := json.Marshal(order)
	producer.Publish(string(transactionJson), os.Getenv("KafkaTransactionsTopic"))
	json.NewEncoder(w).Encode(order)
}

//criar método criar pedido
