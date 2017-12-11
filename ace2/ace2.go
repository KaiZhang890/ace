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
	ConnectHost = "127.0.0.1"
	// ConnectPort is a port
	ConnectPort = "9000"
	// ConnectType is a type
	ConnectType = "tcp"
)

/*
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

*/
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

/*
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

func startGame(group []*user) {
	d1, d2, d3, _ := poker.Deal()

	str := fmt.Sprintf("You got:%s\n", d1)
	group[0].conn.Write([]byte(str))
	str = fmt.Sprintf("You got:%s\n", d2)
	group[1].conn.Write([]byte(str))
	str = fmt.Sprintf("You got:%s\n", d3)
	group[2].conn.Write([]byte(str))

}
*/

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
