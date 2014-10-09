//N customer go to a buffet and eat from a pot. The last customer to find the pot
//empty requests the chef to cook more food.
package buffet

import (
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	SERVINGS_PER_POT    int = 5
	SERVINGS_PER_PERSON int = 4
	PEOPLE              int = 3
)

var logger *log.Logger

type Pot int

func Run() {
	var wg sync.WaitGroup
	logger = log.New(os.Stdout, "buffet", log.LstdFlags)
	chRequest := make(chan Pot)
	chPot := make(chan Pot, 1)

	go Cook(chRequest)
	chPot <- Pot(0)

	for i := 0; i < PEOPLE; i++ {
		wg.Add(1)
		//A placeholder for blocking
		time.Sleep(time.Duration(rand.Int63n(2e8)))
		go Customer(i, chPot, chRequest, &wg)
	}
	wg.Wait()
}

func Customer(id int, pot, request chan Pot, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < SERVINGS_PER_PERSON; i++ {
		logger.Println("Customer", id, "is hungry")
		qty := <-pot
		if qty <= 0 {
			logger.Println("Customer", id, "find the buffet empty")
			//Signal the cook
			request <- qty
			//Block and wait for the cook to finish cooking
			qty = <-request
		}
		logger.Println("Customer", id, "is eating")
		pot <- qty - 1
		time.Sleep(time.Duration(rand.Int63n(1e8)))
	}
	logger.Println("Customer", id, "has finished eating")
}

func Cook(pot chan Pot) {
	for {
		logger.Println("Cook is sleeping")
		//Block till a request for more arrives
		<-pot
		logger.Println("The cook grumbles,", "\"I am cooking.\"")
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		logger.Println("The cook says,", "\"Here you go.\"")
		pot <- Pot(SERVINGS_PER_POT)
	}
}
