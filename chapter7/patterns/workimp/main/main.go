package main

import (
	"fmt"
	"github.com/goinaction/code/chapter7/patterns/workimp"
	"strconv"
	"sync"
)

type Job struct {
	Name string
}

func (j Job) Task() {
	fmt.Println(j.Name)
}

func main() {
	maxgo := 100

	var wg sync.WaitGroup

	wg.Add(maxgo)
	pool := workimp.New(3)

	for i := 0; i < 100; i++ {
		go func(id int) {
			defer wg.Done()
			s := strconv.FormatInt(int64(id), 10)
			job := Job{s}
			pool.Run(job)
		}(i)
	}
	wg.Wait()
	pool.ShutDown()
}
