package main

import (
	"fmt"
	"sync"
)

func sendData(ch chan<- int, data int, wg *sync.WaitGroup) {
	defer wg.Done()

	// send data to channal
	ch <- data // blocks until recieved

	fmt.Printf("%v has been send!\n", data)
}

func recieveData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// recieve data to channal
	data := <-ch

	fmt.Printf("Recieved: %v\n", data)
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go sendData(ch, 9, &wg)
	go recieveData(ch, &wg)
	wg.Wait()

	fmt.Println("Both send and recieved are done")



	wg.Add(2)
	go sendData(ch, 69, &wg)
	go recieveData(ch, &wg)
	wg.Wait()

	fmt.Println("Both send and recieved are done")



	wg.Add(2)
	go sendData(ch, 100, &wg)
	go recieveData(ch, &wg)
	wg.Wait()

	fmt.Println("Both send and recieved are done")
	
}
