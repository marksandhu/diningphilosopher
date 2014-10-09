package main

import (
	"fmt"

	"./classic"
	"./mutex"

	"./buffet"
)

func main() {
	fmt.Println("Starting classic")
	classic.Run()
	fmt.Println("Finished classic")

	fmt.Println("Starting mutex")
	mutex.Run()
	fmt.Println("Finished mutex")

	fmt.Println("Starting buffet")
	buffet.Run()
	fmt.Println("Finished buffet")
}
