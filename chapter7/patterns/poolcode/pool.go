package pool

import (
    "sync"
    "io"
    "errors"
    "fmt"
)
type Pool struct {
    m sync.Mutex
    resource chan io.Closer
    factory func() (io.Closer, error)
    closed bool
}

var ErrClosed = errors.New("pool has been closed")

func New(f func()(io.Closer, error), size uint) (*Pool, error) {
    fmt.Println("new pool")
    if size<=0 {
        return nil, errors.New("pool size is too small")
    }

    return &Pool {
        resource: make(chan io.Closer, size),
        factory: f,
    }, nil
}


func (p *Pool) Aquire() (io.Closer, error) {
    p.m.Lock()
    defer p.m.Unlock()

    if p.closed {
        return nil, ErrClosed
    }

    select {
        case obj, ok := <-p.resource:
            if !ok {
                return nil, ErrClosed
            }
            return obj, nil
        default:
            fmt.Println("new obj")
            return p.factory()
    }
}

func (p *Pool)Close() {
    p.m.Lock()
    defer p.m.Unlock()

    if p.closed {
        return
    }
    
    p.closed=true
    close(p.resource)// 关闭通道，不再产生新的数据

    for r :=range p.resource {
        r.Close()
    }
}


func (p *Pool) Release(r io.Closer) {
    p.m.Lock()
    defer p.m.Unlock()

    if p.closed {
        r.Close()
        return
    }

    select {
        case p.resource<-r:
            fmt.Println("in queue")
        default:
            fmt.Println("queue is full")
            r.Close()
    }
}
