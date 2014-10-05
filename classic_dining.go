package main

import "runtime"
import "sync"
import "fmt"

const TOTAL int = 5

func Dine(id int, l_fork, r_fork chan int) {
	//Get the left fork
	left := <-l_fork
	right := <-r_fork

	fmt.Println(id, "is eating with forks", left, "and", right)
	runtime.Gosched()
	fmt.Println(id, "is finished eating")

	//Give back the forks
	l_fork <- left
	r_fork <- right
	runtime.Gosched()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	forks := make([]chan int, TOTAL)
	for i := 0; i < TOTAL; i++ {
		wg.Add(1)
		go func(k int, wg *sync.WaitGroup) {
			defer wg.Done()

			fork := make(chan int, 1)
			forks[k] = fork
			fork <- k
		}(i, &wg)
	}
	wg.Wait()
	var wg2 sync.WaitGroup
	for i := 0; i < TOTAL; i++ {
		wg2.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()

			Dine(id, forks[id], forks[(id+1)%TOTAL])
		}(i, &wg2)
	}
	wg2.Wait()
	for i := 0; i < TOTAL; i++ {
		close(forks[i])
		<-forks[i]
	}
}
