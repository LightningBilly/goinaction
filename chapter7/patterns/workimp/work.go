package workimp

import (
	"fmt"
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	queue chan Worker
	wg    sync.WaitGroup
}

func (p *Pool) Run(w Worker) {
	p.queue <- w
}

func (p *Pool) ShutDown() {
	close(p.queue)
	p.wg.Wait()
}

func New(poolSize int) *Pool {
	pool := &Pool{
		queue: make(chan Worker),
	}

	pool.wg.Add(poolSize)
	for i := 0; i < poolSize; i++ {
		go func(id int) {
			defer pool.wg.Done()
			for w := range pool.queue {
				fmt.Println("go ", id, "work")
				w.Task()
			}
		}(i)
	}
	return pool
}
