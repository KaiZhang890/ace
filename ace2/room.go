package main

import (
	"bytes"
	"fmt"
	"log"
)

type Room struct {
	u1 *User
	u2 *User
	u3 *User
}

func (r Room) String() string {
	var res bytes.Buffer
	res.WriteString("[")
	if r.u1 != nil {
		res.WriteString(r.u1.name + " ")
	}
	if r.u2 != nil {
		res.WriteString(r.u2.name + " ")
	}
	if r.u3 != nil {
		res.WriteString(r.u3.name + " ")
	}
	res.Truncate(res.Len() - 1)
	res.WriteString("]")

	return res.String()
}

func (r Room) sendToAll(b []byte) {
	var err error
	if r.u1 != nil {
		_, err = r.u1.client.conn.Write(b)
		if err != nil {
			log.Println(err)
		}
	}

	if r.u2 != nil {
		_, err = r.u2.client.conn.Write(b)
		if err != nil {
			log.Println(err)
		}
	}

	if r.u3 != nil {
		_, err = r.u3.client.conn.Write(b)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) u1Read(cmd string) {
	log.Printf("%s: u1Read %s\n", r, cmd)

	str := fmt.Sprintf("%s: %s\n", r.u1.name, cmd)
	r.sendToAll([]byte(str))
}

func (r *Room) u1Close() {
	log.Printf("%s u1Close\n", r)

	r.u1.close()
	r.u1 = nil
}

func (r *Room) u2Read(cmd string) {
	log.Printf("%s: u2Read %s\n", r, cmd)

	str := fmt.Sprintf("%s: %s\n", r.u2.name, cmd)
	r.sendToAll([]byte(str))
}

func (r *Room) u2Close() {
	log.Printf("%s u2Close\n", r)

	r.u2.close()
	r.u2 = nil
}

func (r *Room) u3Read(cmd string) {
	log.Printf("%s: u3Read %s\n", r, cmd)

	str := fmt.Sprintf("%s: %s\n", r.u3.name, cmd)
	r.sendToAll([]byte(str))
}

func (r *Room) u3Close() {
	log.Printf("%s u3Close\n", r)

	r.u3.close()
	r.u3 = nil
}

func NewRoom(u1 *User, u2 *User, u3 *User) *Room {
	r := Room{u1, u2, u3}
	r.u1.state = stateStart
	r.u1.client.onRead = r.u1Read
	r.u1.client.onClose = r.u1Close

	r.u2.state = stateStart
	r.u2.client.onRead = r.u2Read
	r.u2.client.onClose = r.u2Close

	r.u3.state = stateStart
	r.u3.client.onRead = r.u3Read
	r.u3.client.onClose = r.u3Close

	str := fmt.Sprintf("%s in a room\n", r)
	r.sendToAll([]byte(str))

	log.Printf("new %s\n", r)
	return &r
}
