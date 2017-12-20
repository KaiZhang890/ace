package main

import (
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	// ConnectHost is a host
	ConnectHost = ""
	// ConnectPort is a port
	ConnectPort = "9000"
	// ConnectType is a type
	ConnectType = "tcp"
)

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
	c := NewClient(conn)
	defer func() {
		c.onClose()
	}()

OuterLoop:
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break OuterLoop
			} else {
				log.Fatal(err)
				continue
			}
		}

		userInput := strings.TrimSpace(string(buffer[:n]))
		c.onRead(userInput)
	}
}

var allUser []*User
var mutex = &sync.RWMutex{}

func AddUser(u *User) {
	mutex.Lock()
	log.Println(allUser)
	allUser = append(allUser, u)
	log.Println(allUser)
	mutex.Unlock()
}

func RemoveUser(u *User) {
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

func handleMatching() {
	for {
		var group []*User
		mutex.RLock()
		for _, user := range allUser {
			if user.state == stateMatching {
				group = append(group, user)
				if len(group) == 3 {
					break
				}
			}
		}
		mutex.RUnlock()

		if len(group) == 3 {
			NewRoom(group[0], group[1], group[2])
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
