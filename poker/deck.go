package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func init() {
	fmt.Println("deck init()")
}

// Deck holds the cards in the deck to be shuffled
type Deck []Card

func newDeck() Deck {
	deck := Deck{}

	types := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Heart", "Diamond", "Spade", "Club"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}

	deck = append(deck, Card{Type: "Colored", Suit: "Joker"})
	deck = append(deck, Card{Type: "Black", Suit: "Joker"})

	return deck
}

func (d Deck) desc() string {
	var ret bytes.Buffer

	for _, c := range d {
		ret.WriteString(c.desc())
		ret.WriteString("\t")
	}

	return ret.String()
}

func (d Deck) shuffle() {
	for i := 1; i < len(d); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
}

func (d Deck) deckType() (string, bool) {
	switch len(d) {
	case 1:
		return "单张", true
	case 2:
		if d[0].Type == d[1].Type {
			return "一对", true
		} else if d[0].Suit == "Joker" && d[1].Suit == "Joker" {
			return "火箭", true
		}
	case 3:
		if (d[0].Type == d[1].Type) && (d[1].Type == d[2].Type) {
			return "三张", true
		}
	case 4:
		str, ok := handleDeckType4(d)
		if ok {
			return str, true
		}
	case 5:
		str, ok := handleDeckType5(d)
		if ok {
			return str, true
		}
	case 6:
		str, ok := handleDeckType6(d)
		if ok {
			return str, true
		}
	case 7:
		str, ok := handleDeckType7(d)
		if ok {
			return str, true
		}
	case 8:
		str, ok := handleDeckType8(d)
		if ok {
			return str, true
		}
	case 9:
		str, ok := handleDeckType9(d)
		if ok {
			return str, true
		}
	case 10:
		str, ok := handleDeckType10(d)
		if ok {
			return str, true
		}
	case 11:
		str, ok := handleDeckType11(d)
		if ok {
			return str, true
		}
	case 12:
		str, ok := handleDeckType12(d)
		if ok {
			return str, true
		}
	case 13:
		str, ok := handleDeckType13(d)
		if ok {
			return str, true
		}
	case 14:
		str, ok := handleDeckType14(d)
		if ok {
			return str, true
		}
	case 15:
		str, ok := handleDeckType15(d)
		if ok {
			return str, true
		}
	case 16:
		str, ok := handleDeckType16(d)
		if ok {
			return str, true
		}
	case 17:
		str, ok := handleDeckType17(d)
		if ok {
			return str, true
		}
	case 18:
		str, ok := handleDeckType18(d)
		if ok {
			return str, true
		}
	case 19:
		str, ok := handleDeckType19(d)
		if ok {
			return str, true
		}
	case 20:
		str, ok := handleDeckType20(d)
		if ok {
			return str, true
		}
	}

	return "", false
}

func handleDeckType4(d Deck) (string, bool) {
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	if len(m) == 1 {
		return "炸弹", true
	} else if len(m) == 2 {
		return "三带一", true
	}

	return "", false
}

func handleDeckType5(d Deck) (string, bool) {
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 2 {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		v1 := m[keys[0]]
		v2 := m[keys[1]]
		if (v1 == 3 && v2 == 2) || (v1 == 2 && v2 == 3) {
			return "三带一对", true
		}
	} else if len(m) == 5 {
		sort.Sort(d)
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 五张", true
		}
	}

	return "", false
}

func handleDeckType6(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 2 {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		v1 := m[keys[0]]
		v2 := m[keys[1]]
		if v1 == 4 || v2 == 4 {
			return "四带二", true
		} else if v1 == 3 && v2 == 3 {
			subStr := d[0].Type + "-" + d[3].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "飞机", true
			}
		}
	} else if len(m) == 3 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 三对", true
			}
		} else if v1 == 4 || v2 == 4 || v3 == 4 {
			return "四带二", true
		}
	} else if len(m) == 6 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 六张", true
		}
	}
	return "", false
}

func handleDeckType7(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 7 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" + d[6].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 七张", true
		}
	}
	return "", false
}

func handleDeckType8(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 2 {
		v1 := m[d[0].Type]
		v2 := m[d[7].Type]
		if v1 == 4 && v2 == 4 {
			subStr := d[0].Type + "-" + d[7].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "航天飞机", true
			}
		}
	} else if len(m) == 3 {
		var keys []string
		for _, c := range d {
			if m[c.Type] == 3 {
				keys = append(keys, c.Type)
			}
		}
		if len(keys) == 6 {
			subStr := keys[0] + "-" + keys[3]
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "飞机带小翼", true
			}
		}

	} else if len(m) == 4 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[5].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 四对", true
			}
		} else if v2 == 3 && v5 == 3 {
			subStr := d[2].Type + "-" + d[5].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "飞机带小翼", true
			}
		}
	} else if len(m) == 8 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" + d[6].Type + "-" + d[7].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 八张", true
		}
	}
	return "", false
}

func handleDeckType9(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 9 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" +
			d[6].Type + "-" + d[7].Type + "-" + d[8].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 九张", true
		}
	}

	return "", false
}

func handleDeckType10(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 4 {
		var keys []string
		for _, c := range d {
			if m[c.Type] == 3 {
				keys = append(keys, c.Type)
			}
		}
		if len(keys) == 6 {
			subStr := keys[0] + "-" + keys[3]
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "飞机带大翼", true
			}
		}
	} else if len(m) == 5 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[8].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 && v5 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type + "-" + d[8].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 五对", true
			}
		}
	} else if len(m) == 10 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" +
			d[6].Type + "-" + d[7].Type + "-" + d[8].Type + "-" + d[9].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 十张", true
		}
	}

	return "", false
}

func handleDeckType11(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 11 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" +
			d[6].Type + "-" + d[7].Type + "-" + d[8].Type + "-" + d[9].Type + "-" + d[10].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 十一张", true
		}
	}

	return "", false
}

func handleDeckType12(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 4 {
		v1 := m[d[2].Type]
		v2 := m[d[7].Type]
		if v1 == 4 && v2 == 4 {
			subStr := d[2].Type + "-" + d[7].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "航天飞机 带大翼", true
			}
		}
	} else if len(m) == 6 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[8].Type]
		v6 := m[d[10].Type]
		v7 := m[d[7].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 && v5 == 2 && v6 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type + "-" + d[8].Type + "-" + d[10].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 六对", true
			}
		} else if v2 == 4 && v7 == 4 {
			subStr := d[2].Type + "-" + d[7].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "航天飞机 带小翼", true
			}
		} else {
			var keys []string
			for _, c := range d {
				if m[c.Type] == 3 {
					keys = append(keys, c.Type)
				}
			}
			if len(keys) == 9 {
				subStr := keys[0] + "-" + keys[3] + "-" + keys[6]
				allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
				if strings.Contains(allStr, subStr) {
					return "飞机带小翼", true
				}
			}
		}
	} else if len(m) == 12 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" +
			d[6].Type + "-" + d[7].Type + "-" + d[8].Type + "-" + d[9].Type + "-" + d[10].Type + "-" + d[11].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return "单顺 十二张", true
		}
	}

	return "", false
}

func handleDeckType13(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	return "", false
}

func handleDeckType14(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 7 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[8].Type]
		v6 := m[d[10].Type]
		v7 := m[d[12].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 && v5 == 2 && v6 == 2 && v7 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type + "-" + d[8].Type + "-" + d[10].Type + "-" + d[12].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 七对", true
			}
		}
	}

	return "", false
}

func handleDeckType15(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 6 {
		var keys []string
		for _, c := range d {
			if m[c.Type] == 3 {
				keys = append(keys, c.Type)
			}
		}
		if len(keys) == 9 {
			subStr := keys[0] + "-" + keys[3] + "-" + keys[6]
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "飞机带大翼", true
			}
		}
	}

	return "", false
}

func handleDeckType16(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 8 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[8].Type]
		v6 := m[d[10].Type]
		v7 := m[d[12].Type]
		v8 := m[d[14].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 && v5 == 2 && v6 == 2 && v7 == 2 && v8 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type + "-" + d[8].Type + "-" +
				d[10].Type + "-" + d[12].Type + "-" + d[14].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 八对", true
			}
		} else {
			var keys []string
			for _, c := range d {
				if m[c.Type] == 3 {
					keys = append(keys, c.Type)
				}
			}
			if len(keys) == 12 {
				subStr := keys[0] + "-" + keys[3] + "-" + keys[6] + "-" + keys[9]
				allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
				if strings.Contains(allStr, subStr) {
					return "飞机带小翼", true
				}
			}

		}
	}

	return "", false
}

func handleDeckType17(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	return "", false
}

func handleDeckType18(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 6 {
		var keys []string
		for _, c := range d {
			if m[c.Type] == 4 {
				keys = append(keys, c.Type)
			}
		}
		if len(keys) == 12 {
			subStr := keys[0] + "-" + keys[4] + "-" + keys[8]
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "航天飞机带大翼", true
			}
		}
	} else if len(m) == 9 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[8].Type]
		v6 := m[d[10].Type]
		v7 := m[d[12].Type]
		v8 := m[d[14].Type]
		v9 := m[d[16].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 && v5 == 2 && v6 == 2 && v7 == 2 && v8 == 2 && v9 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type + "-" + d[8].Type + "-" +
				d[10].Type + "-" + d[12].Type + "-" + d[14].Type + "-" + d[16].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 九对", true
			}
		} else {
			var keys []string
			for _, c := range d {
				if m[c.Type] == 4 {
					keys = append(keys, c.Type)
				}
			}
			if len(keys) == 12 {
				subStr := keys[0] + "-" + keys[4] + "-" + keys[8]
				allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
				if strings.Contains(allStr, subStr) {
					return "航天飞机带小翼", true
				}
			}
		}
	}

	return "", false
}

func handleDeckType19(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	return "", false
}

func handleDeckType20(d Deck) (string, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 8 {
		var keys []string
		for _, c := range d {
			if m[c.Type] == 3 {
				keys = append(keys, c.Type)
			}
		}
		if len(keys) == 12 {
			subStr := keys[0] + "-" + keys[4] + "-" + keys[8]
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "飞机带大翼", true
			}
		}
	} else if len(m) == 10 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[8].Type]
		v6 := m[d[10].Type]
		v7 := m[d[12].Type]
		v8 := m[d[14].Type]
		v9 := m[d[16].Type]
		v10 := m[d[18].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 && v5 == 2 && v6 == 2 && v7 == 2 && v8 == 2 && v9 == 2 && v10 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type + "-" + d[8].Type + "-" +
				d[10].Type + "-" + d[12].Type + "-" + d[14].Type + "-" + d[16].Type + "-" + d[18].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return "双顺 十对", true
			}
		}
	}

	return "", false
}

func (d Deck) play(preDeck Deck) bool {
	if len(preDeck) > 0 {
		if len(preDeck) == 1 && len(d) == 1 {

		}
	} else {
		switch len(d) {
		case 1:
			return true
		case 2:
			if (d[0].Type == d[1].Type) || (d[0].Suit == "Joker" && d[1].Suit == "Joker") {
				return true
			}
		case 3:
			if (d[0].Type == d[1].Type) && (d[1].Type == d[2].Type) {
				return true
			}
		case 4:
			m := make(map[string]int)
			for _, c := range d {
				m[c.Type]++
			}

			for _, v := range m {
				if v >= 3 {
					return true
				}
			}
		case 5:

		}
	}

	return false
}

func (d Deck) Len() int {
	return len(d)
}

func (d Deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Deck) Less(i, j int) bool {
	if d[i].Type == "Colored" {
		return false

	} else if d[i].Type == "Black" {
		if d[j].Type == "Colored" {
			return true
		}
		return false

	} else if d[i].Type == "2" {
		switch d[j].Type {
		case "Colored", "Black":
			return true
		default:
			return false
		}

	} else if d[i].Type == "Ace" {
		switch d[j].Type {
		case "Colored", "Black", "2":
			return true
		default:
			return false
		}

	} else if d[i].Type == "King" {
		switch d[j].Type {
		case "Colored", "Black", "2", "Ace":
			return true
		default:
			return false
		}

	} else if d[i].Type == "Queen" {
		switch d[j].Type {
		case "Colored", "Black", "2", "Ace", "King":
			return true
		default:
			return false
		}

	} else if d[i].Type == "Jack" {
		switch d[j].Type {
		case "3", "4", "5", "6", "7", "8", "9", "10":
			return false
		default:
			return true
		}

	} else if d[i].Type == "10" {
		switch d[j].Type {
		case "3", "4", "5", "6", "7", "8", "9":
			return false
		default:
			return true
		}

	} else if d[i].Type == "9" {
		switch d[j].Type {
		case "3", "4", "5", "6", "7", "8":
			return false
		default:
			return true
		}

	} else if d[i].Type == "8" {
		switch d[j].Type {
		case "3", "4", "5", "6", "7":
			return false
		default:
			return true
		}

	} else if d[i].Type == "7" {
		switch d[j].Type {
		case "3", "4", "5", "6":
			return false
		default:
			return true
		}

	} else if d[i].Type == "6" {
		switch d[j].Type {
		case "3", "4", "5":
			return false
		default:
			return true
		}

	} else if d[i].Type == "5" {
		switch d[j].Type {
		case "3", "4":
			return false
		default:
			return true
		}

	} else if d[i].Type == "4" {
		switch d[j].Type {
		case "3":
			return false
		default:
			return true
		}

	} else if d[i].Type == "3" {
		return true
	}

	return true
}
