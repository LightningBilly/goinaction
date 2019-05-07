// Sample program to show how to show you how to briefly work
// with the sql package.
package main

import (
	_ "github.com/goinaction/code/chapter2/cbbwork/matchers"
	"github.com/goinaction/code/chapter2/cbbwork/search"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	log.Println("start serviece")
	search.Run("Month")
}
