package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Seed our randomness with the current time
func init() {
	fmt.Println("poker init()")
	rand.Seed(time.Now().UnixNano())
}

func main() {
	index := 0
	var d1, d2, d3, d4 Deck
	reader := bufio.NewReader(os.Stdin)
	var b1, b2, b3 int
	var turn int
L:
	for {
		switch index {
		case 0:
			d1, d2, d3, d4 = deal()
			fmt.Println("User1:", d1)
			fmt.Println("User2:", d2)
			fmt.Println("User3:", d3)
			fmt.Println("Remain:", d4)
			index++
		case 1:
			fmt.Print("User1 bid: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			b1, _ = strconv.Atoi(text)
			index++
		case 2:
			fmt.Print("User2 bid: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			b2, _ = strconv.Atoi(text)
			index++
		case 3:
			fmt.Print("User3 bid: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			b3, _ = strconv.Atoi(text)
			index++
		case 4:
			if b2 > b1 && b2 > b3 {
				turn = 1
			} else if b3 > b1 && b3 > b2 {
				turn = 2
			} else {
				turn = 0
			}

			if turn == 0 {
				d1 = append(d1, d4...)
				sort.Sort(d1)
				fmt.Println("User1 is the Boss")
			} else if turn == 1 {
				d2 = append(d2, d4...)
				sort.Sort(d2)
				fmt.Println("User2 is the Boss")
			} else if turn == 2 {
				d3 = append(d3, d4...)
				sort.Sort(d3)
				fmt.Println("User3 is the Boss")
			}

			index++
		case 5:
			var curDeck Deck
			curUser := -1
			isWin := false
			for !isWin {
				fmt.Println()
				ti := turn % 3
				var ud *Deck
				if ti == 0 {
					fmt.Println("User1's turn to play")
					fmt.Println("User1:", d1.showIndex())
					ud = &d1
				} else if ti == 1 {
					fmt.Println("User2's turn to play")
					fmt.Println("User2:", d2.showIndex())
					ud = &d2
				} else if ti == 2 {
					fmt.Println("User3's turn to play")
					fmt.Println("User3:", d3.showIndex())
					ud = &d3
				}

				fmt.Println("Input the cards(3,3,3,4):")
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text)
				ss := strings.Split(text, ",")
				sd, rd := ud.play2(ss)
				if len(sd) == 0 {
					fmt.Println("Pass")
					turn++
					continue
				}

				if len(curDeck) == 0 || curUser == ti {
					_, ok := sd.deckType()
					if ok {
						curDeck = sd
						curUser = ti
						*ud = rd
						fmt.Println("Play:", sd)
						if len(rd) == 0 {
							fmt.Println("Win!!!!")
							isWin = true
						}
					}
				} else {
					ok := sd.canPlay(curDeck)
					if ok {
						curDeck = sd
						curUser = ti
						*ud = rd
						fmt.Println("Play:", sd)
						if len(rd) == 0 {
							fmt.Println("Win!!!!")
							isWin = true
						}
					}
				}
				turn++
			}
			index++
		case 6:
			fmt.Println("User1:", d1)
			fmt.Println("User2:", d2)
			fmt.Println("User3:", d3)
			fmt.Println("Continue(C) or End:")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if text == "C" {
				index = 0
			} else {
				break L
			}
		default:
			fmt.Print("TBD")
			reader.ReadString('\n')
		}
	}
}

func deal() (Deck, Deck, Deck, Deck) {
	deck := newDeck()
	deck.shuffle()

	d1 := make(Deck, 17)
	copy(d1, deck[:17])
	sort.Sort(d1)

	d2 := make(Deck, 17)
	copy(d2, deck[17:34])
	sort.Sort(d2)

	d3 := make(Deck, 17)
	copy(d3, deck[34:51])
	sort.Sort(d3)

	d4 := make(Deck, 3)
	copy(d4, deck[51:54])

	return d1, d2, d3, d4
}
