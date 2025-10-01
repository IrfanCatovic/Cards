package main

import "fmt"

//Create a new type of 'deck
// //which is a slice of strings


type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamons", "Hearts","Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues {
			cards = append(cards, value + "of" + suit)
		}
		//ovde _ je upisano umesto indeksa "i" "j" ali posto ih nigde ne ispisujemo
		//samo upisemo ovo da postoje, ali ne trebaju nam
	}

	return cards
}

func (d deck) print () {
	for i, card:= range d{
		fmt.Println(i,card)
	}
}