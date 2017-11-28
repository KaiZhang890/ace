package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Seed our randomness with the current time
func init() {
	fmt.Println("poker init()")
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// deck := New()
	// Debug(deck)
	// Shuffle(deck)
	// Debug(deck)
	// Deal(deck, 2)
	fmt.Println("main()")
	// deck := newDeck()
	// deck.shuffle()

	// // deal 发牌
	// d1 := deck[:17]
	// sort.Sort(d1)
	// fmt.Println(d1.desc())
	// d2 := deck[17:34]
	// sort.Sort(d2)
	// fmt.Println(d2.desc())
	// d3 := deck[34:51]
	// sort.Sort(d3)
	// fmt.Println(d3.desc())
	// d4 := deck[51:54]
	// fmt.Println(d4.desc())

	// d5 := append(d3, d4...)
	// sort.Sort(d5)
	// fmt.Println(d5.desc())

	d1 := Deck{Card{Type: "3", Suit: "Heart"}, Card{Type: "4", Suit: "Spade"}, Card{Type: "3", Suit: "Club"}, Card{Type: "3", Suit: "Spade"}}
	b := d1.play(Deck{})
	fmt.Println(b)
}
