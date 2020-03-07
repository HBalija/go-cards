package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	ss := []string(d)
	return strings.Join(ss, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := string(bs)
	ss := strings.Split(s, ",")
	return deck(ss)
}

func (d deck) shuffle() {
	// create a random seed for generator depnding on current time
	source := rand.NewSource(time.Now().UnixNano())
	// use source as a basis for a new random number generator
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[newPosition], d[i] = d[i], d[newPosition]
	}
}
