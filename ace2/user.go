package main

import (
	"fmt"
	"log"
	"strings"
)

type userState int

const (
	stateIdle userState = iota
	stateMatching
	stateStart
	statePlaying
)

type User struct {
	client *Client
	name   string
	state  userState
}

func (u User) String() string {
	str := fmt.Sprintf("%s[%s][%d]", u.name, u.client, u.state)
	return str
}

func (u *User) read(cmd string) {
	log.Printf("%s: read %s\n", u, cmd)

	if u.state == stateIdle {
		if strings.ToUpper(cmd) == "Y" {
			u.state = stateMatching
		} else {
			u.askReady()
		}
	}
}

func (u *User) close() {
	log.Printf("%s close\n", u)
	u.client.close()
}

func (u *User) askReady() {
	str := fmt.Sprintf("%s, ready?[Y/n]", u.name)
	u.client.conn.Write([]byte(str))
}

func NewUser(c *Client, name string) *User {
	u := User{c, name, stateIdle}
	u.client.onRead = u.read
	u.client.onClose = u.close

	u.askReady()
	log.Printf("new %s\n", u)
	return &u
}
