package main

import (
	"ace/poker"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

const (
	// ConnectHost is a host
	ConnectHost = "127.0.0.1"
	// ConnectPort is a port
	ConnectPort = "9000"
	// ConnectType is a type
	ConnectType = "tcp"
)

type userState int

const (
	stateNaming userState = iota
	stateIdle
	stateMatching
	stateStart
	statePlaying
)

type user struct {
	name  string
	conn  net.Conn
	state userState
	group []*user
}

func (u *user) handleInput(userInput string) {
	switch u.state {
	case stateNaming:
		if utf8.RuneCountInString(userInput) > 0 {
			u.name = userInput
			u.state = stateIdle

			str := fmt.Sprintf("%s, are you ready?[Y/n]", u.name)
			u.conn.Write([]byte(str))
		} else {
			u.conn.Write([]byte("\ntype your name: "))
		}
	case stateIdle:
		if strings.ToUpper(userInput) == "Y" {
			u.state = stateMatching
			u.conn.Write([]byte("Matching..."))
		} else {
			str := fmt.Sprintf("%s, are you ready?[Y/n]", u.name)
			u.conn.Write([]byte(str))
		}
	case stateMatching:
		u.conn.Write([]byte("Matching..."))
	case stateStart:
		if len(userInput) > 0 {
			for _, user := range u.group {
				str := fmt.Sprintf("%s say: %s\n", u.name, userInput)
				user.conn.Write([]byte(str))
			}
		}
	}
}

var allUser []*user
var mutex = &sync.RWMutex{}

func addUser(u *user) {
	mutex.Lock()
	allUser = append(allUser, u)
	mutex.Unlock()
}

func removeUser(u *user) {
	mutex.Lock()
	for i, u2 := range allUser {
		if u == u2 {
			copy(allUser[i:], allUser[i+1:])
			allUser[len(allUser)-1] = nil
			allUser = allUser[:len(allUser)-1]
			break
		}
	}
	mutex.Unlock()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting...")
	listener, err := net.Listen(ConnectType, ConnectHost+":"+ConnectPort)
	e(err)
	defer listener.Close()
	log.Println("listening on " + ConnectHost + ":" + ConnectPort)

	go handleMatching()
	// handle accept
	for {
		connect, err := listener.Accept()
		e(err)
		if err != nil {
			continue
		}

		go handleConnect(connect)
	}
}

func handleConnect(conn net.Conn) {
	currentUser := user{"", conn, stateNaming, nil}

	addUser(&currentUser)
	conn.Write([]byte("type your name: "))

	defer func() {
		str := fmt.Sprintf("%s left\n", currentUser.name)
		group := currentUser.group
		if group != nil {
			for _, user := range group {
				if user != &currentUser {
					user.conn.Write([]byte(str))
				}
			}

			currentUser.group = nil
		}
		conn.Close()
	}()

OuterLoop:
	for {
		buffer := make([]byte, 100)
		n, err := conn.Read(buffer)
		if err == io.EOF {
			removeUser(&currentUser)
			break OuterLoop
		} else {
			e(err)
		}

		userInput := strings.TrimSpace(string(buffer[:n]))
		currentUser.handleInput(userInput)
	}
}

func handleMatching() {
	for {
		var group []*user
		for _, user := range allUser {
			if user.state == stateMatching {
				group = append(group, user)
				if len(group) == 3 {
					break
				}
			}
		}

		if len(group) == 3 {
			str := fmt.Sprintf("\n%s, %s, %s in a group\n", group[0].name, group[1].name, group[2].name)
			for _, user := range group {
				user.group = group
				user.state = stateStart
				user.conn.Write([]byte(str))
			}
			startGame(group)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func startGame(group []*user) {
	d1, d2, d3, _ := poker.Deal()

	str := fmt.Sprintf("You got:%s\n", d1)
	group[0].conn.Write([]byte(str))
	str = fmt.Sprintf("You got:%s\n", d2)
	group[1].conn.Write([]byte(str))
	str = fmt.Sprintf("You got:%s\n", d3)
	group[2].conn.Write([]byte(str))

}
