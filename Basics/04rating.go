package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name, rating string
	var numRating float64

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	reader = bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your rating: ")
	rating, _ = reader.ReadString('\n')
	for {
		rating = strings.TrimSpace(rating)
		numRating, _ = strconv.ParseFloat(rating, 64)
		if 1 <= numRating && numRating <= 5 {
			break
		} else {
			reader = bufio.NewReader(os.Stdin)
			fmt.Printf("Please enter a valid rating between 1 and 5: ")
			rating, _ = reader.ReadString('\n')
		}
	}

	fmt.Printf("%v Thanks for rating our service %v stars %v", time.Now().Format(time.Stamp), numRating, name)
}
