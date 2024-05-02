package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"rsc.io/quote"
)

func main() {
	// Start the slow task
	go slowTask()

	// Do something else
	color.HiMagenta("Hello, World, here's a quote:")
	fmt.Println()
	color.Magenta(quote.Go())
}

func slowTask() {
	fmt.Print("Starting the slow task")
	time.Sleep(30 * time.Second)
	fmt.Print("Ending the slow task")
}

