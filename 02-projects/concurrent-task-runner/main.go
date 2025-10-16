package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID    string
	value int
}

type Result struct {
	JobID       string
	OutputValue int
}

// worker to read form jobs and write to the results
func worker(id int, jobs <-chan Job, results chan<- Result) {
	fmt.Printf("Worker %d started.\n", id)

	for job := range jobs {
		output := job.value * job.value

		time.Sleep(1 * time.Millisecond) // simulate delay

		// write result to the results channel
		results <- Result{JobID: job.ID, OutputValue: output}

		fmt.Printf("Worker %s is finished job %d.\n", job.ID, id)
	}

	fmt.Printf("Worker %d stopped.\n", id)
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	// create buffered channels
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// track go routines
	var wg sync.WaitGroup
	// launching wokers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // icrement the worker for each worker launched

		// lauch worker go routine
		go func(workerId int) {
			defer wg.Done() // decrement the counter when the functions exits

			worker(workerId, jobs, results)
		}(w)
	}

	// send jobs to the cahannel
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: fmt.Sprintf("id-%d", i), value: i}
	}

	// close the job channel after all jobs are sent
	// this signal the 'worker' goroutines (the 'for range' loop) to stop
	close(jobs)

	// this blocks until the 'wg.done()' hasbeen called 'numWorkers' times
	wg.Wait()

	// this signal the main goroutine's 'for range' loop to stop
	close(results)

	fmt.Println("\n--- Results ---")

	// print all results
	for result := range results {
		fmt.Printf("Collected Result: %s = %d\n", result.JobID, result.OutputValue)
	}
}
