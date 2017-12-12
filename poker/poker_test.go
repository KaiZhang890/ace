package poker

import "testing"
import "fmt"

func Test_DeckPlay(t *testing.T) {
	d := Deck{card("5"), card("6"), card("7"), card("8"), card("9"), card("Queen")}
	fmt.Println(d.showIndex())
	ss := []string{"1", "3", "5"}
	d1, d2 := d.play(ss)
	fmt.Println(d1)
	fmt.Println(d2)
}

func Test_DeckPlay1(t *testing.T) {
	d1 := Deck{card("3")}
	d2 := Deck{card("2")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5")}
	d2 = Deck{card("Ace"), card("Ace")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5"), card("5"), card("3")}
	d2 = Deck{card("8"), card("8"), card("8"), card("3")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("6"), card("7"), card("8"), card("9")}
	d2 = Deck{card("8"), card("9"), card("10"), card("Jack"), card("Queen")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5"), card("6"), card("6"), card("7"), card("7")}
	d2 = Deck{card("Jack"), card("Jack"), card("Queen"), card("Queen"), card("King"), card("King")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5"), card("5"),
		card("6"), card("6"), card("6"),
		card("3"), card("3")}
	d2 = Deck{card("7"), card("7"), card("7"),
		card("8"), card("8"), card("8"),
		card("4"), card("5")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5"), card("5"), card("5")}
	d2 = Deck{card("2"), card("2"), card("2"), card("2")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5")}
	d2 = Deck{card("3"), card("3"), card("3"), card("3")}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5"), card("5"), card("5")}
	d2 = Deck{Card{"Black", "Joker"}, Card{"Colored", "Joker"}}
	checkCanPlay(d1, d2)

	d1 = Deck{card("5"), card("5"), card("5"),
		card("6"), card("6"), card("6")}
	d2 = Deck{Card{"Black", "Joker"}, Card{"Colored", "Joker"}}
	checkCanPlay(d1, d2)
}

func checkCanPlay(d1 Deck, d2 Deck) {
	if d2.CanPlay(d1) {
		fmt.Println(d2, "can beat", d1)
	} else {
		fmt.Println(d2, "can't beat", d1)
	}
}

func Test_DeckType(t *testing.T) {
	// 41
	// 13, 17, 19
	d := Deck{card("3")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3")}
	checkDeckType(&d)

	d = Deck{Card{"Black", "Joker"}, Card{"Colored", "Joker"}}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("4")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("4"), card("4")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"), card("6"), card("7")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"), card("4"), card("4")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("3"), card("3"), card("4"), card("4")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("5"), card("5"), card("4"), card("4")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"), card("5"), card("4")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"), card("6"), card("7"), card("8")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"), card("6"), card("7"), card("8"), card("9")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"), card("4"), card("4"), card("4"), card("4")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("5"), card("4"), card("4"), card("4"), card("5")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("4"), card("4"), card("5"), card("5"), card("6"), card("6")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("5"), card("4"), card("4"), card("4"), card("6")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"), card("6"), card("7"), card("8"), card("9"), card("10")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"), card("6"), card("7"), card("8"), card("9"), card("10"), card("Jack")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("4"), card("4"), card("4"), card("5"), card("5"), card("6"), card("6")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("4"), card("4"), card("5"), card("5"), card("6"), card("6"), card("7"), card("7")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"), card("6"), card("7"), card("8"), card("9"), card("10"), card("Jack"), card("Queen")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"), card("6"), card("7"), card("8"), card("9"), card("10"), card("Jack"), card("Queen"), card("King")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"), card("5")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"), card("4"),
		card("5"), card("5"), card("6"), card("6")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"),
		card("4"), card("4"),
		card("5"), card("5"),
		card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"), card("4"),
		card("5"), card("6"), card("7"), card("8")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"),
		card("6"), card("7"), card("8")}
	checkDeckType(&d)

	d = Deck{card("3"), card("4"), card("5"),
		card("6"), card("7"), card("8"),
		card("9"), card("10"), card("Jack"),
		card("Queen"), card("King"), card("Ace")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"),
		card("4"), card("4"),
		card("5"), card("5"),
		card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8"),
		card("9"), card("9")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"),
		card("6"), card("6"), card("6"),
		card("7"), card("7"), card("7")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"),
		card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"), card("5"),
		card("6"), card("6"), card("6"), card("6")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"),
		card("4"), card("4"),
		card("5"), card("5"),
		card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8"),
		card("9"), card("9"),
		card("10"), card("10")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"),
		card("6"), card("6"), card("6"),
		card("7"), card("8"), card("9"), card("10")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"), card("5"),
		card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"),
		card("4"), card("4"),
		card("5"), card("5"),
		card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8"),
		card("9"), card("9"),
		card("10"), card("10"),
		card("Jack"), card("Jack")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"), card("5"),
		card("6"), card("7"),
		card("8"), card("9"),
		card("10"), card("Jack")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"), card("3"),
		card("4"), card("4"), card("4"),
		card("5"), card("5"), card("5"),
		card("6"), card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8"),
		card("9"), card("9"),
		card("10"), card("10")}
	checkDeckType(&d)

	d = Deck{card("3"), card("3"),
		card("4"), card("4"),
		card("5"), card("5"),
		card("6"), card("6"),
		card("7"), card("7"),
		card("8"), card("8"),
		card("9"), card("9"),
		card("10"), card("10"),
		card("Jack"), card("Jack"),
		card("Queen"), card("Queen")}
	checkDeckType(&d)
}

func card(t string) Card {
	return Card{t, ""}
}

func checkDeckType(d *Deck) {
	dt, ok := d.deckType()
	if ok {
		fmt.Println(d, dt)
	} else {
		fmt.Println(d, "无效")
	}
}
