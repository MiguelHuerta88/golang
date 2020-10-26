package main

func main() {
	t := triangle{
		height: 50.56,
		base:   10.00,
	}

	s := square{
		sideLength: 36.00,
	}

	printArea(s)
	printArea(t)
}
