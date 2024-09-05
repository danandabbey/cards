package main

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}
