package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
// have to run e.g go run main.go deck.go
type deck []string

// function that will return two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// went rougue and tried to create a similar function using receiver func but it didnt work.
// got this for the output
/*
	0 Spades of Aces
	1 Spades of Two
	2 Spades of Three
	3 Spades of Four
	4 Spades of Five
	0 Spades of Aces
	1 Spades of Two
	2 Spades of Three
	3 Spades of Four
	4 Spades of Five
*/
// doesnt work becuase its using reference to changes instance im assuming
/*func (d deck) dealUpdate(handSize int) (deck, deck) {
	return d[:handSize], d[:handSize]
}*/

// no receiver to this function
func newDeck() deck {
	cards := deck{}

	// creating the deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// replace variable with _ when you want to tell go theres a variable but your not going to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver functions. receivers usually assign reference value 1-2 char abbreviation. Thats why we used d
// usually a regular function will be structured like so
//	func add(i int, y int)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	// []string(d) will convert deck to a string[]. since thats what a deck essentially is based on our definition
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1. log the error and return newDeck()
		// option 2. log the error and entirely quit program
		fmt.Println("Error", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {
	/*source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		randomNum := r.Intn(len(d) - 1)
		d[i], d[randomNum] = d[randomNum], d[i] // fancy one liner swap
	}*/

	//Another way to do the same thing. code on top seems outdated
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		randomNum := rand.Intn(len(d) - 1)
		d[i], d[randomNum] = d[randomNum], d[i]
	}
}
