package main

import (
	"fmt"
	"os"
	"user-vote/routes"
)

func main() {

	fmt.Println("Startin User Vote Rest Server with Go")
	fmt.Println(os.Getenv("KafkaBootstrapServers"))
	routes.HandleResquest()
}
