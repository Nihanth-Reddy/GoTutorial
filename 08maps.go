package main

import "fmt"

func main() {
	fmt.Println("____Maps Tutorial_____")

	var elements = map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Berelium",
		"B":  "Boron",
	}

	fmt.Println("Length of elements is: ", len(elements))
	elements["C"] = "Carbon"
	elements["O"] = "Oxygen"
	elements["N"] = "Nitrogen"
	fmt.Println(elements)

	students := make(map[string]map[string]string)
	students["Nihanth"] = map[string]string{
		"Age":    "25",
		"Gender": "Male",
		"Course": "GoLang",
	}
	students["Reddy"] = map[string]string{
		"Age":    "25",
		"Gender": "Male",
		"Course": "GoLang",
	}
	students["Kottam Nihanth Reddy"] = map[string]string{
		"Age":    "25",
		"Gender": "Male",
		"Course": "To Lang",
	}

	// fmt.Println("Students map is: ", students)
	fmt.Println("Name", "Age", "Gender", "Course")
	for k, v := range students {
		fmt.Println(k)
		for i, j := range v {
			fmt.Printf("	%v: %v\n", i, j)
		}
	}
	delete(students, "Kottam Nihanth Reddy")
	fmt.Println(students)
}
