package main

import "fmt"

func main() {
	// maps are reference type. so when passed to functions they can directly modify data
	/* ways to declare a map */
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
	//fmt.Println(colors)

	/*// creating maps using var syntax
	var colors1 map[string]string
	fmt.Println(colors1)

	// creating map using make command
	colors2 := make(map[string]string)
	colors2["blue"] = "#dfds343"
	fmt.Println(colors2)

	// how to delete keys from maps
	delete(colors2, "blue")
	fmt.Println(colors2)*/
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
