package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Tutorial")
	var rings = []string{"gold", "diamond", "platinum", "silver"}

	fmt.Printf("The cap of rings is: %d, and ring at 3rd index is: %v\n", cap(rings), rings[3])

	rings = append(rings, "silver")
	fmt.Printf("After adding %v at the %d index the cap of rings is %d\n", rings[3], len(rings)-1, cap(rings))

	things := make([]string, 4, 4)
	fmt.Println(len(things))
	things[0] = "Computer"
	things[1] = "Keyboard"
	things[2] = "mouse"
	things[3] = "monitor"
	fmt.Println(len(things))
	for _, value := range things {
		fmt.Println(value)
	}
	fmt.Printf("Appending the %v at the end of the things\n", "Speakers")
	things = append(things, "speakers")
	fmt.Println(len(things), cap(things))

	var intsSlice = []int{5, 6, 3, 7, 9, 11, 2}
	isSorted := sort.IntsAreSorted(intsSlice)

	if isSorted {
		fmt.Println("Given integers slice is sorted and shortest number is: ", intsSlice[0])
	} else {
		fmt.Println("Given integers slice is not sorted, sorting now...")
		sort.Ints(intsSlice)
		fmt.Println("Sorted the Slice and shortest number is: ", intsSlice[0])
	}
}
