package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	fmt.Fprintln(&b, "World!")

	fmt.Println(b.String())
	b.WriteTo(os.Stdout)
}
