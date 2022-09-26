package main

import (
	"fmt"

	"user-vote/routes"
)

func main() {

	fmt.Println("Startin User Vote Rest Server with Go")
	routes.HandleResquest()
}
