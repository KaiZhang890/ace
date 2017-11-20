package main

import "testing"
import "time"
import "log"

type client struct {
}

func Test_Select(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			log.Println(m1)
		case m2 := <-c2:
			log.Println(m2)
		}
	}
}
