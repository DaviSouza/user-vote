package routes

import (
	"log"
	"net/http"
	"user-vote/controllers"
	"user-vote/database"
	"user-vote/middleware"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	client := etherResquest()
	db := database.ConnectDataBase()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllUsers(w, r, db)
	}).Methods("Get")
	//r.HandleFunc("/api/user/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(w, r, db)
	}).Methods("Post")
	r.HandleFunc("/api/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r, db)
	}).Methods("Delete")
	r.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r, db)
	}).Methods("Put")
	r.HandleFunc("/api/order/balance", func(w http.ResponseWriter, r *http.Request) {
		controllers.Balance(w, r, client)
	})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

func etherResquest() *ethclient.Client {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	return client
}
