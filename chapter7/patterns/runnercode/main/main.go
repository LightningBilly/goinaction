package main

import (
	"fmt"
	"github.com/goinaction/code/chapter7/patterns/runnercode"
	"time"
)

var TimeOut = 3 * time.Second

func job(id int) {
	fmt.Println("Job : ", id)
	time.Sleep(time.Duration(id+1) * time.Second)
}

func main() {
	fmt.Println("start run")
	r := runner.New(TimeOut)
	r.Add(job, job)
	fmt.Println(r.Start())
}
