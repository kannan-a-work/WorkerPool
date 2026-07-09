package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int)
	Results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(jobs, &wg, Results)
	}
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(Results)
	}()
	for res := range Results {
		fmt.Println("result is:", res)
	}

	fmt.Println("worker pool jobs finished")
}

func worker(ch <-chan int, wg *sync.WaitGroup, Results chan<- int) {
	defer wg.Done()
	for job := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("worker received job:", job)
		Results <- job * 2
	}
}
