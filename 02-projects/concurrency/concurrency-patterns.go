package main

import (
	"fmt"
//	"time"
)

// recieve a done value from main and instant return
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work")
		}
	}
}

// convert nums to slice channal
func slieceToChannel(nums []int) (<-chan int) {
	sliceChannel := make(chan int)
	go func(){
		for _, num := range nums {
			sliceChannel <- num
		}
		close(sliceChannel)
	}()

	return sliceChannel
}

func sq(in <-chan int) (<-chan int){
	result := make(chan int)
	go func(){
		for i := range in {
			result <- i * i // squre the value
		}

		close(result)
	}()

	return result
}

func main() {

	// 1- For  Select Pattern
	/*
	charChannal := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
		// put each character on the channal
		case charChannal <- char:

		}
	}

	func() {
		for {
			select {
			default:
				fmt.Println("Doing Work")
			}
		}
	}()

	close(charChannal)

	for result := range charChannal {
		fmt.Println("Result: ", result)
	}
	time.Sleep(time.Second * 10)
	*/



	// 2- Done Channal
	/*
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)
	close(done) // paass signal to done channal
	*/


	// 3- Pipeline (we are using unbuffered channal inside 'slieceToChannel' and 'sq', so communication b/w them is synchrounus)
	// input
	nums := []int{1, 2, 3, 4, 5}
	// stage 1: convert array to slice channal
	dataChannel := slieceToChannel(nums)
	// stage 2: square each item
	finalChannal := sq(dataChannel)
	// stage 3: print each line
	for n := range finalChannal {
		fmt.Println("Value: ", n)
	}



}
