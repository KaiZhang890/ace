package main

import "testing"
import "fmt"

func Test_DeckType(t *testing.T) {
	d := Deck{Card{"3", "Heart"}}
	checkDeckType(&d)

	d = Deck{Card{"Ace", "Heart"}, Card{"Ace", "Diamond"}}
	checkDeckType(&d)

	d = Deck{Card{"Colored", "Joker"}, Card{"Black", "Joker"}}
	checkDeckType(&d)

	d = Deck{Card{"Queen", "Heart"}, Card{"Queen", "Diamond"}, Card{"Queen", "Spade"}}
	checkDeckType(&d)

	d = Deck{Card{"8", "Heart"}, Card{"8", "Diamond"}, Card{"8", "Spade"}, Card{"8", "Club"}}
	checkDeckType(&d)

	d = Deck{Card{"10", "Heart"}, Card{"2", "Diamond"}, Card{"10", "Spade"}, Card{"10", "Club"}}
	checkDeckType(&d)

	d = Deck{Card{"10", "Heart"}, Card{"10", "Diamond"}, Card{"King", "Spade"}, Card{"10", "Club"}, Card{"King", "Diamond"}}
	checkDeckType(&d)

	d = Deck{Card{"3", "Heart"}, Card{"4", "Heart"}, Card{"7", "Spade"}, Card{"6", "Club"}, Card{"5", "Spade"}}
	checkDeckType(&d)

	d = Deck{Card{"10", "Heart"}, Card{"10", "Diamond"}, Card{"10", "Spade"}, Card{"10", "Club"}, Card{"King", "Diamond"}}
	checkDeckType(&d)
}

func Test_DeckType2(t *testing.T) {
	d := Deck{{"Ace", "Heart"}, {"9", "Heart"}, {"10", "Diamond"}, {"Jack", "Spade"}, {"Queen", "Club"}, {"King", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"7", "Heart"}, {"7", "Spade"}, {"9", "Diamond"}, {"9", "Spade"}, {"8", "Club"}, {"8", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"Ace", "Heart"}, {"Ace", "Spade"}, {"Ace", "Diamond"}, {"Ace", "Club"}, {"8", "Club"}, {"3", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"4", "Spade"}, {"4", "Diamond"}, {"3", "Club"}, {"4", "Club"}, {"3", "Diamond"}}
	checkDeckType(&d)
}

func Test_DeckType3(t *testing.T) {
	d := Deck{{"4", "Heart"}, {"5", "Heart"}, {"9", "Diamond"}, {"10", "Spade"}, {"6", "Club"}, {"7", "Diamond"}, {"8", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"3", "Diamond"}, {"3", "Club"},
		{"4", "Heart"}, {"4", "Spade"}, {"4", "Diamond"}, {"4", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"4", "Heart"}, {"4", "Spade"}, {"5", "Diamond"}, {"5", "Club"}}
	checkDeckType(&d)

	d = Deck{{"9", "Heart"}, {"Jack", "Spade"}, {"6", "Diamond"}, {"8", "Club"},
		{"7", "Heart"}, {"10", "Spade"}, {"4", "Diamond"}, {"5", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"3", "Heart"}, {"4", "Spade"}, {"Ace", "Diamond"}, {"6", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"3", "Heart"}, {"4", "Spade"}, {"7", "Diamond"}, {"7", "Club"}}
	checkDeckType(&d)
}

func Test_DeckType4(t *testing.T) {
	d := Deck{{"3", "Heart"}, {"4", "Spade"}, {"7", "Diamond"}, {"6", "Club"},
		{"5", "Heart"}, {"8", "Spade"}, {"9", "Diamond"}, {"10", "Club"}, {"Jack", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"4", "Spade"}, {"7", "Diamond"}, {"6", "Club"}, {"5", "Heart"},
		{"8", "Spade"}, {"9", "Diamond"}, {"10", "Club"}, {"Jack", "Club"}, {"Queen", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"4", "Spade"}, {"7", "Diamond"}, {"6", "Club"}, {"5", "Heart"},
		{"8", "Spade"}, {"9", "Diamond"}, {"10", "Club"}, {"Jack", "Club"}, {"Queen", "Club"},
		{"King", "Spade"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"4", "Spade"}, {"7", "Diamond"}, {"6", "Club"}, {"5", "Heart"},
		{"8", "Spade"}, {"9", "Diamond"}, {"10", "Club"}, {"Jack", "Club"}, {"Queen", "Club"},
		{"King", "Spade"}, {"Ace", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"7", "Club"}, {"7", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"7", "Club"}, {"7", "Club"}, {"8", "Diamond"}, {"8", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"7", "Club"}, {"7", "Club"}, {"8", "Diamond"}, {"8", "Club"},
		{"9", "Diamond"}, {"9", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"7", "Club"}, {"7", "Club"}, {"8", "Diamond"}, {"8", "Club"},
		{"9", "Diamond"}, {"9", "Club"}, {"10", "Diamond"}, {"10", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"7", "Club"}, {"7", "Club"}, {"8", "Diamond"}, {"8", "Club"},
		{"9", "Diamond"}, {"9", "Club"}, {"10", "Diamond"}, {"10", "Club"},
		{"Jack", "Diamond"}, {"Jack", "Club"}}
	checkDeckType(&d)

	d = Deck{{"3", "Heart"}, {"3", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"7", "Club"}, {"7", "Club"}, {"8", "Diamond"}, {"8", "Club"},
		{"9", "Diamond"}, {"9", "Club"}, {"10", "Diamond"}, {"10", "Club"},
		{"Jack", "Diamond"}, {"Jack", "Club"}, {"Queen", "Diamond"}, {"Queen", "Club"}}
	checkDeckType(&d)
}

func Test_DeckType5(t *testing.T) {
	d := Deck{{"4", "Heart"}, {"4", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"5", "Diamond"}, {"5", "Club"},
		{"3", "Club"}, {"7", "Club"}, {"Jack", "Diamond"}, {"Queen", "Club"}}
	checkDeckType(&d)

	d = Deck{{"4", "Heart"}, {"4", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"5", "Diamond"}, {"5", "Club"},
		{"3", "Club"}, {"3", "Club"}, {"Jack", "Diamond"}, {"Jack", "Club"}}
	checkDeckType(&d)

	d = Deck{{"10", "Heart"}, {"10", "Spade"}, {"10", "Diamond"},
		{"Jack", "Club"}, {"Jack", "Heart"}, {"Jack", "Spade"},
		{"Queen", "Diamond"}, {"Queen", "Club"}, {"Queen", "Club"},
		{"3", "Club"}, {"3", "Diamond"},
		{"5", "Club"}, {"5", "Diamond"},
		{"7", "Club"}, {"7", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"10", "Heart"}, {"10", "Spade"}, {"10", "Diamond"},
		{"Jack", "Club"}, {"Jack", "Heart"}, {"Jack", "Spade"},
		{"Queen", "Diamond"}, {"Queen", "Club"}, {"Queen", "Club"},
		{"3", "Club"},
		{"5", "Diamond"},
		{"7", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"10", "Heart"}, {"10", "Spade"}, {"10", "Diamond"},
		{"Jack", "Club"}, {"Jack", "Heart"}, {"Jack", "Spade"},
		{"3", "Diamond"}, {"3", "Club"},
		{"4", "Diamond"}, {"4", "Club"}}
	checkDeckType(&d)

	d = Deck{{"10", "Heart"}, {"10", "Spade"}, {"10", "Diamond"},
		{"Jack", "Club"}, {"Jack", "Heart"}, {"Jack", "Spade"},
		{"3", "Diamond"},
		{"8", "Club"}}
	checkDeckType(&d)

}

func Test_DeckType6(t *testing.T) {
	d := Deck{{"4", "Heart"}, {"4", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"5", "Diamond"}, {"5", "Club"},
		{"6", "Heart"}, {"6", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"3", "Club"},
		{"7", "Club"},
		{"8", "Club"},
		{"9", "Club"},
		{"Jack", "Diamond"},
		{"Queen", "Club"}}
	checkDeckType(&d)

	d = Deck{{"4", "Heart"}, {"4", "Spade"}, {"4", "Diamond"}, {"4", "Club"},
		{"5", "Heart"}, {"5", "Spade"}, {"5", "Diamond"}, {"5", "Club"},
		{"6", "Heart"}, {"6", "Spade"}, {"6", "Diamond"}, {"6", "Club"},
		{"3", "Club"}, {"3", "Diamond"},
		{"8", "Club"}, {"8", "Spade"},
		{"Jack", "Diamond"}, {"Jack", "Club"}}
	checkDeckType(&d)
}

func Test_DeckType7(t *testing.T) {
	d := Deck{{"10", "Heart"}, {"10", "Spade"}, {"10", "Diamond"},
		{"Jack", "Club"}, {"Jack", "Heart"}, {"Jack", "Spade"},
		{"Queen", "Diamond"}, {"Queen", "Club"}, {"Queen", "Club"},
		{"King", "Diamond"}, {"King", "Club"}, {"King", "Club"},
		{"3", "Club"},
		{"5", "Diamond"},
		{"4", "Diamond"},
		{"7", "Diamond"}}
	checkDeckType(&d)

	d = Deck{{"10", "Heart"}, {"10", "Spade"}, {"10", "Diamond"},
		{"Jack", "Club"}, {"Jack", "Heart"}, {"Jack", "Spade"},
		{"Queen", "Diamond"}, {"Queen", "Club"}, {"Queen", "Club"},
		{"King", "Diamond"}, {"King", "Club"}, {"King", "Club"},
		{"3", "Club"}, {"3", "Diamond"},
		{"5", "Diamond"}, {"5", "Club"},
		{"4", "Diamond"}, {"4", "Club"},
		{"7", "Diamond"}, {"7", "Club"}}
	checkDeckType(&d)
}

func checkDeckType(d *Deck) {
	str, ok := d.deckType()
	if ok {
		fmt.Println(d, str)
	} else {
		fmt.Println(d, "无效")
	}
}
