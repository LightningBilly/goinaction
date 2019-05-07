package main

import (
    "fmt"
    "time"
    "github.com/goinaction/code/chapter7/patterns/runnercode"
)

var TimeOut = 3 * time.Second

func job(id int) {
    fmt.Println("Job : ", id)
}

func main() {
	fmt.Println("start run")
    r := runner.New(TimeOut)
    r.Add(job, job)
    fmt.Println(r.Start())
}
