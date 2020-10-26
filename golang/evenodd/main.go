package main

import "fmt"

func main() {
	n := generateSlice(10)

	// now we display the type of number
	print(n)
}

func generateSlice(x int) []int {
	s := []int{}
	for i := 0; i < x; i++ {
		s = append(s, i)
	}
	return s
}

func print(s []int) {
	for i, value := range s {
		var t string
		if value%2 == 0 {
			t = "even"
		} else {
			t = "odd"
		}

		fmt.Println(i, " is "+t)
	}
}
