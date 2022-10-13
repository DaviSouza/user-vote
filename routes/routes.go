package routes

import (
	"log"
	"net/http"
	"user-vote/controllers"
	"user-vote/database"
	"user-vote/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	db := database.ConnectDataBase()
	//producer := setupKafkaProducer() "user-vote/kafka"
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
	/*
		r.HandleFunc("/api/order/transfer", func(w http.ResponseWriter, r *http.Request) {
			controllers.Transfer(w, r, producer)
		}).Methods("Post")*/

	r.HandleFunc("/api/order/balance", controllers.Balance).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

/*
func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))
	return producer
}*/
