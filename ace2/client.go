package main

import (
	"log"
	"net"
	"unicode/utf8"
)

type Client struct {
	conn    net.Conn
	onRead  func(cmd string)
	onClose func()
}

func (c Client) String() string {
	str := c.conn.RemoteAddr().String()
	return str
}

func (c *Client) read(cmd string) {
	log.Printf("%s: read %s\n", c, cmd)

	if utf8.RuneCountInString(cmd) > 0 {
		u := NewUser(c, cmd)
		AddUser(u)
	} else {
		c.write([]byte("\ntype your name: "))
	}
}

func (c *Client) write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}

func (c *Client) close() {
	log.Printf("%s close\n", c)
	c.onRead = nil
	c.conn.Close()
}

func NewClient(conn net.Conn) *Client {
	c := Client{conn, nil, nil}
	c.onRead = c.read
	c.onClose = c.close
	c.write([]byte("type your name: "))

	log.Printf("new %s\n", c)
	return &c
}
