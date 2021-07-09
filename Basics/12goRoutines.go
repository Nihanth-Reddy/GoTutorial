package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func printNumbers(prefix string) {
	for i := range prefix {
		fmt.Println(prefix, ":", i)
		sleepDuration := time.Duration(rand.Intn(500))
		time.Sleep(time.Millisecond * sleepDuration)
	}
}

func maxInt(s []int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	maxNumber := s[0]
	for _, val := range s {
		if val > maxNumber {
			maxNumber = val
		}
	}
	fmt.Println(time.Now().Format(time.StampMicro), "Maximum integer in given slice is ", maxNumber)
}

func minInt(s []int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	minNumber := s[0]
	for _, val := range s {
		if val < minNumber {
			minNumber = val
		}
	}
	fmt.Println(time.Now().Format(time.StampMicro), "Minimum integer in given slice is ", minNumber)
}

func average(s []int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	averageNumber := 0
	for _, val := range s {
		averageNumber += val
	}
	fmt.Println(time.Now().Format(time.StampMicro), "Average integer in given slice is ", averageNumber/len(s))
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var waitGroup sync.WaitGroup
	fmt.Println(time.Now().Format(time.StampMicro), "Start of max int")
	waitGroup.Add(1)
	go maxInt(s, &waitGroup)
	fmt.Println(time.Now().Format(time.StampMicro), "Start of average int")
	waitGroup.Add(1)
	go average(s, &waitGroup)
	fmt.Println(time.Now().Format(time.StampMicro), "Start of min int")
	waitGroup.Add(1)
	go minInt(s, &waitGroup)
	// for i := 0; i < 10; i++ {
	// 	go printNumbers("Number")
	// }
	waitGroup.Wait()
	// var input string
	// fmt.Scanf("%s", &input)
}
