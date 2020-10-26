package main

import "fmt"

func main() {
	eb := englishBot{}
	printGreeting(eb)

	sb := spanishBot{}
	printGreeting(sb)
}

// interface can call getGreeting on struct types
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

// since we are not using eb we can remove it from declaration
func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
