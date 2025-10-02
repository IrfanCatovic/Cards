package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// fmt.Println(card)
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// cards := newDeck()
	//Kreira novi deck of cards
	
	cards := newDeckFromFile("my_cards")
	hand, remainingCards := deal(cards, 5)
	//vraca 2 vrednosti, onim redom koji smo definisali u deck.go
	//prva vrednost je karte koje smo dobili, a druga koje su ostale u deck
	//oboje su type deck jer to deal vraca

	//deal(cards, deck) saljemo string karte i broj int u f-ju

	hand.print()
	remainingCards.print()

	fmt.Println(cards.toString()) //pretvara slice of string u jedan red stringova

	// cards.saveToFile("my_cards")


	cards.print()
	//ispisuje u consolu sve karte, list

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
}

// func newCard() string {
// 	return "Five of Diamonds"
// }