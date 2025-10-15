package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "procces  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "started job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := range 3 {
		go worker(w, jobs, results)
	}

	for j := range numJobs {
		jobs <- j
	}

	for range numJobs {
		<-results
	}

	close(jobs)
}
