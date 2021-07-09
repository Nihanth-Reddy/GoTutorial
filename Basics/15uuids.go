package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	val := uuid.New()
	fmt.Printf("%v :: %T\n", val, val)
	val2 := uuid.NewString()
	fmt.Printf("%v :: %T\n", val2, val2)
}
