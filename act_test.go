package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"testing"
)

func Test_RuntimeCPU(t *testing.T) {
	// 设置最大的可同时使用的CPU核数和实际cpu核数一致
	//runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
}

type say interface {
	say123() string
}

type person struct {
	name string
}

func (p person) say123() string {
	return p.name + " say 123"
}

type bird struct {
	name string
}

func (b bird) say123() string {
	return b.name + " say 123"
}

func showSay123(s say) {
	fmt.Println(s)
	fmt.Println(s.say123())
}

func Test_SliceAppend(t *testing.T) {
	jay := person{name: "Jay"}
	showSay123(jay)
	pony := bird{name: "Pony"}
	showSay123(pony)
}

func Test_SliceUsage(t *testing.T) {
	var ss []string
	//ss := make([]string, 10)
	if ss == nil {
		fmt.Println("ss is nil")
	}
	ss = append(ss, "a", "b", "c", "d")
	fmt.Printf("len=%d cap=%d slice=%v\n", len(ss), cap(ss), ss)
	ss = append(ss[:2], ss[3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(ss), cap(ss), ss)
}

type Action struct {
	Action  string            `json:"a"`
	Message string            `json:"m"`
	Data    map[string]string `json:"d"`
}

func Test_ParseJson2(t *testing.T) {
	byt := []byte(`{"a":"ready","m":"准备好了","d":{"d1":"aaa","d2":"bbb"}}`)
	act := Action{}
	json.Unmarshal(byt, &act)
	fmt.Println(act.Message)
	fmt.Println(act.Data["d2"])
}

func Test_BytesSuffix(t *testing.T) {
	bytes := []byte("0Aa\r\n")
	len := len(bytes)
	if bytes[len-2] == 13 && bytes[len-1] == 10 {
		t.Logf("%v has \\r\\n suffix", bytes)
	} else {
		t.Errorf("%v not has \\r\\n suffix", bytes)
	}
}

func Test_ParseJson(t *testing.T) {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		t.Error(err)
		return
	}

	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
}
