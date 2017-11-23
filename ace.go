package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type user struct {
	address string
	conn    net.Conn
}

type userOP struct {
	op   string
	user user
}

type group struct {
	groupID string
	users   []user
}

var globalUsers []user

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	service := ":9000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	c1 := make(chan userOP)
	go handleUserOP(c1)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		c1 <- userOP{op: "add", user: user{address: conn.RemoteAddr().String(), conn: conn}}
		go handleClient(conn, c1)
	}
}

func handleUserOP(c chan userOP) {
	for cu := range c {
		if cu.op == "add" {
			globalUsers = append(globalUsers, cu.user)
			str := cu.user.address + " added\n"
			sendToAll(str)
		} else if cu.op == "remove" {
			for i, gu := range globalUsers {
				if gu.address == cu.user.address {
					len := len(globalUsers)
					globalUsers[i] = globalUsers[len-1]
					//globalUsers[len-1] = nil
					globalUsers = globalUsers[:len-1]
					break
				}
			}

			str := cu.user.address + " removed\n"
			sendToAll(str)
		}
	}
}

func sendToAll(str string) {
	for _, gu := range globalUsers {
		gu.conn.Write([]byte(str))
	}
}

func handleClient(conn net.Conn, c chan userOP) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	req := make([]byte, 1024) // 1KB
	defer func() {
		c <- userOP{op: "remove", user: user{address: conn.RemoteAddr().String(), conn: conn}}
		conn.Close()
	}()

	for {
		readLen, err := conn.Read(req)
		if err != nil {
			log.Println(err)
			break
		}
		if readLen == 0 {
			break
		} else {
			// End with '\r\n'
			if req[readLen-2] == 13 && req[readLen-1] == 10 {
				var dat map[string]interface{}
				if err := json.Unmarshal(req[:readLen-2], &dat); err != nil {
					handleInvalidInstruction(conn, req[:readLen-2])
				} else {
					handleInstruction(conn, dat)
				}
			} else {
				handleInvalidInstruction(conn, req[:readLen])
			}

			req = make([]byte, 1024)
		}
	}
}

func handleInstruction(conn net.Conn, instruction map[string]interface{}) {
	v, ok := instruction["a"].(string)
	if ok && v == "a" {
		conn.Close()
	} else {
		log.Println(instruction)
	}
}

func handleInvalidInstruction(conn net.Conn, instruction []byte) {
	var res bytes.Buffer
	res.WriteString("Invalid instruction: ")
	res.Write(instruction)
	res.WriteString("\n")
	conn.Write(res.Bytes())
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
