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

//var globalConns = make([]net.Conn)
var globalConns []net.Conn

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		log.Println(conn)
		globalConns = append(globalConns, conn)
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	req := make([]byte, 1024) // 1KB
	defer func() {
		var tgc []net.Conn
		for _, gc := range globalConns {
			if gc != conn {
				tgc = append(tgc, gc)
			}
		}
		globalConns = tgc
		log.Println(len(globalConns))
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
			var res bytes.Buffer
			// End with '\r\n'
			if req[readLen-2] == 13 && req[readLen-1] == 10 {
				var dat map[string]string
				if err := json.Unmarshal(req[:readLen-2], &dat); err != nil {
					res.WriteString("Invalid instruction: ")
					res.Write(req[:readLen-2])
					res.WriteString("\n")
				} else {
					res.WriteString("Instruction: ")
					res.Write(req[:readLen-2])
					res.WriteString("\n")
					if dat["a"] == "a" {
						conn.Write(res.Bytes())
						break
					}
				}
			} else {
				res.WriteString("Invalid instruction: ")
				res.Write(req[:readLen])
				res.WriteString("\n")
			}

			conn.Write(res.Bytes())
			req = make([]byte, 1024)
		}
	}
}

func closeConn(conn net.Conn) {

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
