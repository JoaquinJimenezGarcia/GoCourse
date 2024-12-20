package main

import "fmt"

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()
	
	hand,_ := deal(cards, 5)

	fmt.Println(hand.toString())
	hand.saveToFile("./cardsgame/myDeck")

	cardsNew := newDeckFromFile("./cardsgame/otherDeck")
	cardsNew.print()
}