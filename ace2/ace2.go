package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"strings"
)

const (
	// ConnectHost is a host
	ConnectHost = "127.0.0.1"
	// ConnectPort is a port
	ConnectPort = "9000"
	// ConnectType is a type
	ConnectType = "tcp"
)

type user struct {
	name  string
	conn  net.Conn
	group []*user
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting...")
	listener, err := net.Listen(ConnectType, ConnectHost+":"+ConnectPort)
	e(err)
	defer listener.Close()

	log.Println("listening on " + ConnectHost + ":" + ConnectPort)

	c1 := make(chan *user)
	go handleReady(c1)
	for {
		connect, err := listener.Accept()
		e(err)
		if err != nil {
			continue
		}

		go handleConnect(connect, c1)
	}
}

func handleConnect(conn net.Conn, c chan *user) {
	user := user{name: "", conn: conn, group: nil}
	cmdLevel := 0
	conn.Write([]byte("type your name: "))

	defer func() {
		user.conn.Close()
		user.group = nil
	}()

OuterLoop:
	for {
		buffer := make([]byte, 100)
		len, err := conn.Read(buffer)
		if err == io.EOF {
			log.Println("user disconnected: ", user.name)
			break
		} else {
			e(err)
		}

		command := strings.TrimSpace(string(buffer[:len]))
		switch cmdLevel {
		case 0:
			user.name = command
			log.Println("client connected: ", user.name)

			var res bytes.Buffer
			res.WriteString("hello, " + user.name + "\n")
			res.WriteString("type command number(1-Ready 2-Exit):")
			conn.Write(res.Bytes())
			cmdLevel++
		case 1:
			if command == "1" {
				c <- &user
				cmdLevel++
			} else if command == "2" {
				break OuterLoop
			} else {
				conn.Write([]byte("type command number(1-Ready 2-Exit):"))
			}
		case 2:
			if user.group != nil {
				for _, u := range user.group {
					if u.name != user.name {
						u.conn.Write([]byte(user.name + " says: " + command + "\n"))
					}
				}
			} else if command == "2" {

			} else {
				conn.Write([]byte("Matching...(2-Cancel)\n"))
			}
		}
	}
}

var readyUsers []*user

func handleReady(c chan *user) {
	for cu := range c {

		isReady := false
		for _, ru := range readyUsers {
			if ru == cu {
				isReady = true
				break
			}
		}
		if !isReady {
			readyUsers = append(readyUsers, cu)
		}

		if len(readyUsers) < 3 {
			cu.conn.Write([]byte("Matching...(2-Cancel)\n"))
		} else {
			str := readyUsers[0].name + ", " + readyUsers[1].name + ", " + readyUsers[2].name + " in a group\n"
			group := []*user{readyUsers[0], readyUsers[1], readyUsers[2]}
			for i := 0; i < 3; i++ {
				readyUsers[i].conn.Write([]byte(str))
				readyUsers[i].group = group
			}
			readyUsers = readyUsers[3:]
		}
	}
}

func handleReadyCancel(c chan *user) {

}

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
