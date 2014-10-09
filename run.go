package main

import (
	"fmt"

	"naina.localhost/learning/classic"

	"./mutex"
)

func main() {
	fmt.Println("Starting classic")
	classic.Run()
	fmt.Println("Finished classic")

	fmt.Println("Starting mutex")
	mutex.Run()
	fmt.Println("Finished mutex")
}
