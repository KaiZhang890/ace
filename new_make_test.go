package main

import "testing"
import "fmt"

type stu struct {
	name string
	age  int
}

func Test_New(t *testing.T) {
	p1 := new(stu)
	if p1 == nil {
		fmt.Println("p1 is nil")
	} else {
		fmt.Println(p1)
		p1.name = "Joy"
		p1.age = 20
		fmt.Println(p1)
	}
}

func Test_Make(t *testing.T) {
	m1 := make([]stu, 2, 10)
	fmt.Printf("lenth of %v is %d, cap is %d\n", m1, len(m1), cap(m1))

	v := newInt()
	fmt.Println(v)
}

func newInt() *int {
	var i int
	fmt.Printf("address of i is %p\n", &i)
	return &i //为何可以返回局部变量呢？
}
