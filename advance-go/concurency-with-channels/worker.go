package main

import (
	"fmt"
	"time"
)

/*
	Worker pool pattern

	Number of jobs and workers.
	Workers are simulating some job.
	3 workers are doing 5 jobs from the pool.


*/
// reference to jobs and results channels
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker: ", id, "started job: ", j)
		time.Sleep(time.Second)
		fmt.Println("worker: ", id, "finished job: ", j)
		// send result to results channel
		result <- j * 2
	}

}
func main() {
	const numJobs = 5
	// Creating a new channel with make (chan val-type).
	// Send a value into a channel using the channel <- syntax.

	// difference betweeen buffered and unbuffered channels
	// buffered channel - make(chan int, numJobs)
	// unbuffered channel - make(chan int)

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start each worker in goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	time.Sleep(3 * time.Second)

	for j := 1; j <= numJobs; j++ {
		jobs <- j
		fmt.Println("Job sent: ", j)
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		// receive result from results channel
		result := <-results
		fmt.Println("Result received: ", result)
	}
}
