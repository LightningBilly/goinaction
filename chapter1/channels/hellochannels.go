package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	wg.Done()
}

type n interface {
	notify()
}

type user struct{}

func (u user) notify() {
	fmt.Print("f")
}

// main is the entry point for the program.
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()

	arr := [5]int{0: 1, 3: 5}
	fmt.Println(arr)

	var x n
	x = &user{}
	x.notify()
	time.Now()
	sort.IsSorted()
}
