//An attempt to use the Chandy/Misra solution.
package mutex

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)
import "log"

var nPhilosphers, nBites = 5, 3 //Number for philosphers, Number of bites each philospher takes
var logger *log.Logger

type Fork struct {
	id int
	sync.Mutex
}

func NewFork(id int) *Fork {
	fork := Fork{
		id: id,
	}

	return &fork
}

type Philospher int

func (p Philospher) String() string {
	return fmt.Sprintf("Philospher %d", p)
}

func Dine(phil Philospher,
	l_fork, r_fork *Fork,
	preferRight bool,
	wg *sync.WaitGroup) {

	defer wg.Done()
	logger.Println(phil, "is seated between forks", l_fork.id, r_fork.id)
	//Number of times the philospher eats
	for i := 0; i < nBites; i++ {
		logger.Println(phil, "is hungry")
		if preferRight {
			GetFork(phil, "right", r_fork)
			GetFork(phil, "left", l_fork)
		} else {
			GetFork(phil, "left", l_fork)
			GetFork(phil, "right", r_fork)
		}
		logger.Println(phil, "is eating")
		time.Sleep(time.Duration(rand.Int63n(1e9)))

		ReturnFork(phil, "left", l_fork)
		ReturnFork(phil, "right", r_fork)

		logger.Println(phil, "is thinking")
		time.Sleep(time.Duration(rand.Int63n(1e9)))
	}

	logger.Println(phil, "finished eating")

}

func GetFork(phil Philospher, phil_hand string, fork *Fork) {
	logger.Println(phil, "reaches", phil_hand, "for fork", fork.id)
	fork.Lock()
	logger.Println(phil, "has fork", fork.id)
}

func ReturnFork(phil Philospher, phil_hand string, fork *Fork) {
	logger.Println(phil, "returns fork", fork.id, "to his", phil_hand)
	fork.Unlock()
	logger.Println(phil, "puts down fork", fork.id)
}

func Run() {
	logger = log.New(os.Stdout, "mutex.go", log.LstdFlags)
	var wg sync.WaitGroup

	fork_0 := NewFork(nPhilosphers)
	left := fork_0

	for i := 1; i < nPhilosphers; i++ {
		wg.Add(1)
		right := NewFork(i)
		go Dine(Philospher(i), left, right, false, &wg)
		left = right
	}
	wg.Add(1)
	go Dine(Philospher(nPhilosphers), left, fork_0, true, &wg)

	wg.Wait()
}
