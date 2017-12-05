package poker

import (
	"net"
)

type State int

const (
	StateNaming State = iota
	StateIdle
	StateMatching
	StateStart
	StatePlaying
)

type User struct {
	name  string
	conn  net.Conn
	state State
}

type Player struct {
	user User
	deck Deck
}

type Room struct {
	p1 *Player
	p2 *Player
	p3 *Player
}
