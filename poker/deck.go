package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("deck init()")
}

// Deck holds the cards in the deck to be shuffled
type Deck []Card

func (d Deck) showIndex() string {
	var ret bytes.Buffer
	ret.WriteString("[")

	for i, c := range d {
		if i == len(d)-1 {
			str := fmt.Sprintf("%s(%d)", c, i)
			ret.WriteString(str)
		} else {
			str := fmt.Sprintf("%s(%d) ", c, i)
			ret.WriteString(str)
		}
	}

	ret.WriteString("]")
	return ret.String()
}

type deckType struct {
	code  int
	desc  string
	value string
}

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

func (d Deck) play(ss []string) (Deck, Deck) {
	var d1 Deck
	var d2 Deck
	for i, c := range d {
		selected := false
		for _, s := range ss {
			j, ok := strconv.Atoi(s)
			if ok == nil && i == j {
				selected = true
			}
		}
		if selected {
			d1 = append(d1, c)
		} else {
			d2 = append(d2, c)
		}
	}
	return d1, d2
}

func (d Deck) canPlay(preDeck Deck) bool {
	dt1, ret := preDeck.deckType()
	if !ret {
		fmt.Println("preDeck is Invalid")
		return false
	}

	dt2, ret := d.deckType()
	if !ret {
		fmt.Println("deck is Invalid")
		return false
	}

	if dt1.code == dt2.code {
		str := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace-2-Black-Colored"
		i1 := strings.Index(str, dt2.value)
		i2 := strings.Index(str, dt1.value)
		if i1 > i2 {
			return true
		}
	} else {
		if dt2.code == 40 {
			return true
		} else if dt2.code == 39 {
			return true
		}
	}

	return false
}

func (d Deck) deckType() (deckType, bool) {
	switch len(d) {
	case 1:
		return deckType{1, "单张", d[0].Type}, true
	case 2:
		if d[0].Type == d[1].Type {
			return deckType{2, "一对", d[0].Type}, true
		} else if d[0].Suit == "Joker" && d[1].Suit == "Joker" {
			return deckType{40, "火箭", d[0].Type}, true
		}
	case 3:
		if (d[0].Type == d[1].Type) && (d[1].Type == d[2].Type) {
			return deckType{3, "三张", d[0].Type}, true
		}
	case 4:
		dt, ok := handleDeckType4(d)
		if ok {
			return dt, true
		}
	case 5:
		dt, ok := handleDeckType5(d)
		if ok {
			return dt, true
		}
	case 6:
		dt, ok := handleDeckType6(d)
		if ok {
			return dt, true
		}
	case 7:
		dt, ok := handleDeckType7(d)
		if ok {
			return dt, true
		}
	case 8:
		dt, ok := handleDeckType8(d)
		if ok {
			return dt, true
		}
	case 9:
		dt, ok := handleDeckType9(d)
		if ok {
			return dt, true
		}
	case 10:
		dt, ok := handleDeckType10(d)
		if ok {
			return dt, true
		}
	case 11:
		dt, ok := handleDeckType11(d)
		if ok {
			return dt, true
		}
	case 12:
		dt, ok := handleDeckType12(d)
		if ok {
			return dt, true
		}
	case 13:
		dt, ok := handleDeckType13(d)
		if ok {
			return dt, true
		}
	case 14:
		dt, ok := handleDeckType14(d)
		if ok {
			return dt, true
		}
	case 15:
		dt, ok := handleDeckType15(d)
		if ok {
			return dt, true
		}
	case 16:
		dt, ok := handleDeckType16(d)
		if ok {
			return dt, true
		}
	case 17:
		dt, ok := handleDeckType17(d)
		if ok {
			return dt, true
		}
	case 18:
		dt, ok := handleDeckType18(d)
		if ok {
			return dt, true
		}
	case 19:
		dt, ok := handleDeckType19(d)
		if ok {
			return dt, true
		}
	case 20:
		dt, ok := handleDeckType20(d)
		if ok {
			return dt, true
		}
	}

	return deckType{}, false
}

func handleDeckType4(d Deck) (deckType, bool) {
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	if len(m) == 1 {
		return deckType{39, "炸弹", d[0].Type}, true
	} else if len(m) == 2 {
		return deckType{4, "三带一", d[1].Type}, true
	}

	return deckType{}, false
}

func handleDeckType5(d Deck) (deckType, bool) {
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
			return deckType{5, "三带一对", d[2].Type}, true
		}
	} else if len(m) == 5 {
		sort.Sort(d)
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return deckType{6, "单顺 五张", d[0].Type}, true
		}
	}

	return deckType{}, false
}

func handleDeckType6(d Deck) (deckType, bool) {
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
			return deckType{7, "四带一对", d[3].Type}, true
		} else if v1 == 3 && v2 == 3 {
			subStr := d[0].Type + "-" + d[3].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return deckType{8, "飞机", d[0].Type}, true
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
				return deckType{9, "双顺 三对", d[0].Type}, true
			}
		} else if v1 == 4 || v2 == 4 || v3 == 4 {
			return deckType{10, "四带二张", d[3].Type}, true
		}
	} else if len(m) == 6 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return deckType{11, "单顺 六张", d[0].Type}, true
		}
	}
	return deckType{}, false
}

func handleDeckType7(d Deck) (deckType, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 7 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" + d[6].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return deckType{12, "单顺 七张", d[0].Type}, true
		}
	}
	return deckType{}, false
}

func handleDeckType8(d Deck) (deckType, bool) {
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
				return deckType{13, "航天飞机", d[0].Type}, true
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
				// 33344455 44555666
				return deckType{14, "飞机带小翼", keys[0]}, true
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
				return deckType{15, "双顺 四对", d[0].Type}, true
			}
		} else if v2 == 3 && v5 == 3 {
			subStr := d[2].Type + "-" + d[5].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				// 33344456 34555666 45556667
				return deckType{14, "飞机带小翼", d[2].Type}, true
			}
		}
	} else if len(m) == 8 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" + d[6].Type + "-" + d[7].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return deckType{16, "单顺 八张", d[0].Type}, true
		}
	}
	return deckType{}, false
}

func handleDeckType9(d Deck) (deckType, bool) {
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
			return deckType{17, "单顺 九张", d[0].Type}, true
		}
	}

	return deckType{}, false
}

func handleDeckType10(d Deck) (deckType, bool) {
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
				// 3344555666 4455566677 5556667788
				return deckType{18, "飞机带大翼", keys[0]}, true
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
				return deckType{19, "双顺 五对", d[0].Type}, true
			}
		}
	} else if len(m) == 10 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" +
			d[6].Type + "-" + d[7].Type + "-" + d[8].Type + "-" + d[9].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return deckType{20, "单顺 十张", d[0].Type}, true
		}
	}

	return deckType{}, false
}

func handleDeckType11(d Deck) (deckType, bool) {
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
			return deckType{21, "单顺 十一张", d[0].Type}, true
		}
	}

	return deckType{}, false
}

func handleDeckType12(d Deck) (deckType, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 3 {
		if m[d[0].Type] == 4 && m[d[5].Type] == 4 && m[d[11].Type] == 4 {
			// 333344445555
			return deckType{22, "航天飞机", d[0].Type}, true
		}
	} else if len(m) == 4 {
		var keys []string
		for _, c := range d {
			if m[c.Type] == 4 {
				keys = append(keys, c.Type)
			}
		}
		if len(keys) == 8 {
			subStr := keys[0] + "-" + keys[7]
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				//333344445566 334455556666 334444555566
				return deckType{23, "航天飞机 带大翼", keys[0]}, true
			}
		}
	} else if len(m) == 6 {
		v1 := m[d[0].Type]
		v2 := m[d[2].Type]
		v3 := m[d[4].Type]
		v4 := m[d[6].Type]
		v5 := m[d[8].Type]
		v6 := m[d[10].Type]
		if v1 == 2 && v2 == 2 && v3 == 2 && v4 == 2 && v5 == 2 && v6 == 2 {
			subStr := d[0].Type + "-" + d[2].Type + "-" + d[4].Type + "-" + d[6].Type + "-" + d[8].Type + "-" + d[10].Type
			allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
			if strings.Contains(allStr, subStr) {
				return deckType{24, "双顺 六对", d[0].Type}, true
			}
		} else if v4 == 4 {
			var keys []string
			for _, c := range d {
				if m[c.Type] == 4 {
					keys = append(keys, c.Type)
				}
			}
			if len(keys) == 8 {
				subStr := keys[0] + "-" + keys[7]
				allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
				if strings.Contains(allStr, subStr) {
					//333344445678 345677778888 344445555678
					return deckType{25, "航天飞机 带小翼", keys[0]}, true
				}
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
					// 333444555678
					return deckType{26, "飞机带小翼", d[0].Type}, true
				}
			}
		}
	} else if len(m) == 12 {
		subStr := d[0].Type + "-" + d[1].Type + "-" + d[2].Type + "-" + d[3].Type + "-" + d[4].Type + "-" + d[5].Type + "-" +
			d[6].Type + "-" + d[7].Type + "-" + d[8].Type + "-" + d[9].Type + "-" + d[10].Type + "-" + d[11].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			return deckType{27, "单顺 十二张", d[0].Type}, true
		}
	}

	return deckType{}, false
}

func handleDeckType13(d Deck) (deckType, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	return deckType{}, false
}

func handleDeckType14(d Deck) (deckType, bool) {
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
				return deckType{28, "双顺 七对", d[0].Type}, true
			}
		}
	}

	return deckType{}, false
}

func handleDeckType15(d Deck) (deckType, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 5 {
		subStr := d[0].Type + "-" + d[3].Type + "-" + d[6].Type + "-" + d[9].Type + "-" + d[12].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			// 333444555666777
			return deckType{29, "飞机", d[0].Type}, true
		}
	} else if len(m) == 6 {
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
				// 333444555667788 334455666777888
				return deckType{30, "飞机带大翼", keys[0]}, true
			}
		}
	}

	return deckType{}, false
}

func handleDeckType16(d Deck) (deckType, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}

	if len(m) == 4 {
		subStr := d[0].Type + "-" + d[4].Type + "-" + d[8].Type + "-" + d[12].Type
		allStr := "3-4-5-6-7-8-9-10-Jack-Queen-King-Ace"
		if strings.Contains(allStr, subStr) {
			// 3333444455556666
			return deckType{31, "航天飞机", d[0].Type}, true
		}
	} else if len(m) == 8 {
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
				return deckType{32, "双顺 八对", d[0].Type}, true
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
					// 3334445556667891
					return deckType{33, "飞机带小翼", keys[0]}, true
				}
			}

		}
	}

	return deckType{}, false
}

func handleDeckType17(d Deck) (deckType, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	return deckType{}, false
}

func handleDeckType18(d Deck) (deckType, bool) {
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
				// 333344445555667788
				return deckType{34, "航天飞机带大翼", keys[0]}, true
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
				return deckType{35, "双顺 九对", d[0].Type}, true
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
					// 3333444455556789ab
					return deckType{36, "航天飞机带小翼", keys[0]}, true
				}
			}
		}
	}

	return deckType{}, false
}

func handleDeckType19(d Deck) (deckType, bool) {
	sort.Sort(d)
	m := make(map[string]int)
	for _, c := range d {
		m[c.Type]++
	}
	return deckType{}, false
}

func handleDeckType20(d Deck) (deckType, bool) {
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
				// 333444555666778899aa
				return deckType{37, "飞机带大翼", keys[0]}, true
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
				return deckType{38, "双顺 十对", d[0].Type}, true
			}
		}
	}

	return deckType{}, false
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
