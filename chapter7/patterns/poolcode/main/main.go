package main

import (
    "sync"
	"fmt"
	"github.com/goinaction/code/chapter7/patterns/poolcode"
	"io"
)

const (
    PoolSize = 2
    GoroteineNum = 25
)

type Connection struct {
	ID int
}

func (con *Connection) Close() error {
	fmt.Println("close con : ", con.ID)
	con.ID = 0
	return nil
}

func (con *Connection) Work(job int) {
	fmt.Println(con.ID, "do job : ", job)
}

var InitId int

func createConnection() (io.Closer, error) {
	InitId++
	return &Connection{
		ID: InitId,
	}, nil
}

func main() {
	fmt.Println("start")
	p, _ := pool.New(createConnection, PoolSize)
    wg := sync.WaitGroup{}
    wg.Add(GoroteineNum)
    for i:=0;i<GoroteineNum;i++ {
        go func(job int) {
            defer wg.Done() 
            con, err := p.Aquire()
            if err != nil {
                fmt.Println("job : ", i, "err : ", err)
                return
            }

            con.(*Connection).Work(job)
            p.Release(con)
        }(i)
    }

    wg.Wait()
    p.Close()
    fmt.Println("finish job")
}
