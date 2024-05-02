package main

import (
	"fmt"

	"github.com/fatih/color"
	"rsc.io/quote"
)

func main() {
	color.HiMagenta("Hello, World, here's a quote:")
	fmt.Println()
	color.Magenta(quote.Go())
}
