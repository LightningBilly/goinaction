package runner

import (
    "os"
    "os/signal"
    "time"
    "errors"
    "fmt"
)
type Runner struct {
    interrupt chan os.Signal //中断信号
    complete chan error
    timeout <-chan time.Time //超时信号
    tasks []func(int)
}

var ErrInterrupt = errors.New("recieved interrupt")
var ErrTimeout = errors.New("recieved timeout")

func New(d time.Duration) *Runner {
    fmt.Println(d)
    return &Runner{
        interrupt : make(chan os.Signal, 1),
        complete: make(chan error),
        timeout : time.After(d),
    }
}

func (r *Runner) Add(funcs ...func(int)) {
    r.tasks = append(r.tasks, funcs...)
}

func (r *Runner) run() error {
    for i, task :=range r.tasks {
        if r.getInterrupt() {
            return ErrInterrupt
        }
        task(i)
    }

    return nil
}

func (r *Runner) Start() error {
    signal.Notify(r.interrupt, os.Interrupt)
    go func() {
        r.complete <- r.run()
    } ()

    select {
        case err:=<-r.complete:
            return err
        case <-r.timeout:
            return ErrTimeout
    }
}

func (r *Runner) getInterrupt() bool {
    select {
        case <-r.interrupt:
            signal.Stop(r.interrupt) //程序退出停止发送信号
            return true
        default:
            return false
    }
}
