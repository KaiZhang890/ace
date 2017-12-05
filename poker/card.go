package poker

import (
	"bytes"
	"fmt"
)

func init() {
	fmt.Println("card init()")
}

// Card holds the card suits and types in the deck
type Card struct {
	Type string
	Suit string
}

func (c Card) String() string {
	//return fmt.Sprintf("%s(%s)", c.Type, c.Suit)
	return c.Type
}

func (c Card) desc() string {
	var ret bytes.Buffer

	switch c.Suit {
	case "Heart":
		ret.WriteString("♥️")
	case "Diamond":
		ret.WriteString("♦️")
	case "Spade":
		ret.WriteString("♠️")
	case "Club":
		ret.WriteString("♣️")
	}

	switch c.Type {
	case "Ace":
		ret.WriteString("A")
	case "2":
		ret.WriteString("2")
	case "3":
		ret.WriteString("3")
	case "4":
		ret.WriteString("4")
	case "5":
		ret.WriteString("5")
	case "6":
		ret.WriteString("6")
	case "7":
		ret.WriteString("7")
	case "8":
		ret.WriteString("8")
	case "9":
		ret.WriteString("9")
	case "10":
		ret.WriteString("10")
	case "Jack":
		ret.WriteString("J")
	case "Queen":
		ret.WriteString("Q")
	case "King":
		ret.WriteString("K")
	}

	if c.Suit == "Joker" && c.Type == "Colored" {
		ret.WriteString("C-Joker")
	} else if c.Suit == "Joker" && c.Type == "Black" {
		ret.WriteString("B-Joker")
	}

	return ret.String()
}
