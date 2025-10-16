package main

import (
	"fmt"
	"sync"
)

type Count struct {
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
}

func (c *Count) increment() {
	c.mu.Lock()
	c.count++
	count := c.count
	c.mu.Unlock()
	fmt.Println("Icrementing...: ", count)
	c.wg.Done()
}

func (c *Count) decrement() {
	c.mu.Lock()
	c.count--
	count := c.count
	c.mu.Unlock()
	fmt.Println("Decrementing...: ", count)
	c.wg.Done()
}

func main() {

	counter := &Count{}

	fmt.Println("Initial Count: ", counter.count)

	// total 28 go-routines
	counter.wg.Add(28)
	// 10 increment
	go func() {
		for i := 0; i < 10; i++ {
			counter.increment()
		}
	}()
	// 3 decrement
	go func() {
		for i := 0; i < 3; i++ {
			counter.decrement()
		}
	}()
	// 5 increment
	go func() {
		for i := 0; i < 5; i++ {
			counter.increment()
		}
	}()
	// 10 decrement
	go func() {
		for i := 0; i < 10; i++ {
			counter.decrement()
		}
	}()
	// wait for go routine to finish
	counter.wg.Wait()

	fmt.Println("Final Count: ", counter.count)
}
