package main

import (
    "io"
    "fmt"
    "os"
    "net/http"
)

func init() {
    if len(os.Args) != 2 {
        fmt.Println("usage ./example url")
        os.Exit(-1)
    }
}

func main() {
    url := os.Args[1]
    r, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }

    if r.StatusCode != 200 {
        fmt.Println("reee--", r.Status)
        return
    }

    io.Copy(os.Stdout, r.Body)
    if err := r.Body.Close(); err != nil {
        fmt.Println(err)
    }
}

