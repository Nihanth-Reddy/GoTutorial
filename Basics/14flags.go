package main

import (
	"flag"
	"fmt"
)

func main() {

	x := flag.String("x", "Nihanth", "enter some string")
	flag.Parse()
	fmt.Println("the value of x is: ", *x)
}
