package main

import (
	"ace/poker"
	"fmt"
	"log"
)

type Player struct {
	user *User
	deck poker.Deck
	bid  int
}

func (p Player) String() string {
	str := fmt.Sprintf("%s %s", p.user.name, p.deck)
	return str
}

func (p *Player) close() {
	log.Printf("%s close\n", p)
	p.user.close()
	p.deck = nil
}

func (p *Player) write(b []byte) (n int, err error) {
	return p.user.write(b)
}

func NewPlayer(u *User) *Player {
	p := Player{u, nil, 0}

	log.Printf("new %s\n", p)
	return &p
}
