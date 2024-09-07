package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []card

type card struct {
	suit  string
	value string
}

func newDeck() deck {
	cards := deck{}
	s := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	v := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range s {
		for _, value := range v {
			cards = append(cards, card{suit, value})
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card.value+" of "+card.suit)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
*/

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(source)

	for i := range d {
		n := r.Intn(len(d) - 1)

		d[i], d[n] = d[n], d[i]
	}
}
