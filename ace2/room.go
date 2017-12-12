package main

import (
	"ace/poker"
	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	p1        *Player
	p2        *Player
	p3        *Player
	rd        poker.Deck
	bidding   bool
	turn      int
	curDeck   poker.Deck
	curPlayer *Player
}

func (r Room) String() string {
	var res bytes.Buffer
	res.WriteString("[")
	if r.p1 != nil {
		res.WriteString(r.p1.user.name + " ")
	}
	if r.p2 != nil {
		res.WriteString(r.p2.user.name + " ")
	}
	if r.p3 != nil {
		res.WriteString(r.p3.user.name + " ")
	}
	res.Truncate(res.Len() - 1)
	res.WriteString("]")

	return res.String()
}

func (r Room) sendToAll(b []byte) {
	var err error
	if r.p1 != nil {
		_, err = r.p1.write(b)
		if err != nil {
			log.Println(err)
		}
	}

	if r.p2 != nil {
		_, err = r.p2.write(b)
		if err != nil {
			log.Println(err)
		}
	}

	if r.p3 != nil {
		_, err = r.p3.write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) u1Read(cmd string) {
	log.Printf("%s: u1Read %s\n", r, cmd)

	if r.turn == 1 {
		if r.bidding {
			r.p1.bid, _ = strconv.Atoi(cmd)
			if r.p1.bid == 3 {
				r.bidDone()
				return
			}
			r.p2.write([]byte("Bid(0-3):"))
		} else {
			ok := r.play(r.p1, cmd)
			if ok {
				if len(r.p1.deck) == 0 {
					str := fmt.Sprintf("%s Win!!!\n", r.p1.user.name)
					r.sendToAll([]byte(str))
					r.over()
				} else {
					str := fmt.Sprintf("You got %s\nInput the cards(3,3,3,4):", r.p2.deck)
					r.p2.write([]byte(str))
				}
			} else {
				return
			}
		}

		r.turn = 2
	}

}

func (r *Room) u1Close() {
	log.Printf("%s u1Close\n", r)

	r.p1.close()
	r.p1 = nil
}

func (r *Room) u2Read(cmd string) {
	log.Printf("%s: u2Read %s\n", r, cmd)

	if r.turn == 2 {
		if r.bidding {
			r.p2.bid, _ = strconv.Atoi(cmd)
			if r.p2.bid == 3 {
				r.bidDone()
				return
			}
			r.p3.write([]byte("Bid(0-3):"))
		} else {
			ok := r.play(r.p2, cmd)
			if ok {
				if len(r.p2.deck) == 0 {
					str := fmt.Sprintf("%s Win!!!\n", r.p2.user.name)
					r.sendToAll([]byte(str))
					r.over()
				} else {
					str := fmt.Sprintf("You got %s\nInput the cards(3,3,3,4):", r.p3.deck)
					r.p3.write([]byte(str))
				}
			} else {
				return
			}
		}

		r.turn = 3
	}
}

func (r *Room) u2Close() {
	log.Printf("%s u2Close\n", r)

	r.p2.close()
	r.p2 = nil
}

func (r *Room) u3Read(cmd string) {
	log.Printf("%s: u3Read %s\n", r, cmd)

	if r.turn == 3 {
		if r.bidding {
			r.p3.bid, _ = strconv.Atoi(cmd)
			r.bidDone()
			return
		}

		ok := r.play(r.p3, cmd)
		if ok {
			if len(r.p3.deck) == 0 {
				str := fmt.Sprintf("%s Win!!!\n", r.p3.user.name)
				r.sendToAll([]byte(str))
				r.over()
			} else {
				str := fmt.Sprintf("You got %s\nInput the cards(3,3,3,4):", r.p1.deck)
				r.p1.write([]byte(str))
			}
		} else {
			return
		}
		r.turn = 1
	}
}

func (r *Room) u3Close() {
	log.Printf("%s u3Close\n", r)

	r.p3.close()
	r.p3 = nil
}

func (r *Room) bidDone() {
	var boss *Player
	if r.p3.bid > r.p1.bid && r.p3.bid > r.p2.bid {
		boss = r.p3
		r.turn = 3
	} else if r.p2.bid > r.p1.bid {
		boss = r.p2
		r.turn = 2
	} else {
		boss = r.p1
		r.turn = 1
	}
	boss.deck = append(boss.deck, r.rd...)
	sort.Sort(boss.deck)
	r.bidding = false

	str := fmt.Sprintf("%s is boss\n", boss.user.name)
	r.sendToAll([]byte(str))

	str = fmt.Sprintf("You got %s\nInput the cards(3,3,3,4):", boss.deck)
	boss.write([]byte(str))
}

func (r *Room) play(p *Player, cmd string) bool {
	ss := strings.Split(cmd, ",")
	sd, rd := p.deck.Play2(ss)
	if len(sd) == 0 {
		if p != r.curPlayer {
			str := fmt.Sprintf("%s pass\n", p.user.name)
			r.sendToAll([]byte(str))
			return true
		}
	} else {
		if len(r.curDeck) == 0 || r.curPlayer == p {
			_, ok := sd.DeckType()
			if ok {
				r.curDeck = sd
				r.curPlayer = p
				p.deck = rd
				str := fmt.Sprintf("%s play:%s\n", p.user.name, sd)
				r.sendToAll([]byte(str))
				return true
			}
		} else {
			ok := sd.CanPlay(r.curDeck)
			if ok {
				r.curDeck = sd
				r.curPlayer = p
				p.deck = rd
				str := fmt.Sprintf("%s play:%s\n", p.user.name, sd)
				r.sendToAll([]byte(str))
				return true
			}
		}
	}

	return false
}

func (r *Room) over() {
	r.sendToAll([]byte("Game Over!\n"))

	r.p1.user.state = stateIdle
	r.p1.user.client.onRead = r.p1.user.read
	r.p1.user.client.onClose = r.p1.user.close

	r.p2.user.state = stateIdle
	r.p2.user.client.onRead = r.p2.user.read
	r.p2.user.client.onClose = r.p2.user.close

	r.p3.user.state = stateIdle
	r.p3.user.client.onRead = r.p3.user.read
	r.p3.user.client.onClose = r.p3.user.close

	r.p1 = nil
	r.p2 = nil
	r.p3 = nil
	r.rd = nil
	r.curDeck = nil
	r.curPlayer = nil
}

func NewRoom(u1 *User, u2 *User, u3 *User) *Room {
	p1 := NewPlayer(u1)
	p2 := NewPlayer(u2)
	p3 := NewPlayer(u3)
	r := Room{p1: p1, p2: p2, p3: p3}

	r.p1.user.state = stateStart
	r.p1.user.client.onRead = r.u1Read
	r.p1.user.client.onClose = r.u1Close

	r.p2.user.state = stateStart
	r.p2.user.client.onRead = r.u2Read
	r.p2.user.client.onClose = r.u2Close

	r.p3.user.state = stateStart
	r.p3.user.client.onRead = r.u3Read
	r.p3.user.client.onClose = r.u3Close

	str := fmt.Sprintf("%s in a room\n", r)
	r.sendToAll([]byte(str))

	r.p1.deck, r.p2.deck, r.p3.deck, r.rd = poker.Deal()
	str = fmt.Sprintf("You got %s\n", r.p1.deck)
	r.p1.write([]byte(str))
	str = fmt.Sprintf("You got %s\n", r.p2.deck)
	r.p2.write([]byte(str))
	str = fmt.Sprintf("You got %s\n", r.p3.deck)
	r.p3.write([]byte(str))

	r.bidding = true
	r.turn = 1
	r.p1.write([]byte("Bid(0-3):"))

	log.Printf("new %s\n", r)
	return &r
}
