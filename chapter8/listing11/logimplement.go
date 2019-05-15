package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func init() {
	log.Println("file write")
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("file failed")
	}
	flag := log.Ldate | log.Ltime | log.Lshortfile
	Trace = log.New(ioutil.Discard, "TRACE：", flag)
	Warning = log.New(os.Stdout, "WARNING: ", flag)
	Info = log.New(os.Stdout, "INFO: ", flag)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flag)
}

func main() {
	Trace.Println("cbb trace")
	//println(Warning)
	Warning.Println("cbb warn")
	Info.Println("info ")
	Error.Println("error info")

}
