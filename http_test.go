package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_GenerateCombinations(t *testing.T) {
	for combination := range GenerateCombinations("0123456789", 4) {
		fmt.Println(combination)
	}
	fmt.Println("Done!")
}

func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)

		AddLetter(c, "", alphabet, length)
	}(c)
	return c
}

func AddLetter(c chan string, combo string, alphabet string, length int) {
	if length <= 0 {
		c <- combo
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		//c <- newCombo
		AddLetter(c, newCombo, alphabet, length-1)
	}
}

func Test_Request(t *testing.T) {
	restCode := sendRequest("123456")
	log.Println(restCode)
}

func sendRequest(password string) string {
	var jsonStr bytes.Buffer
	jsonStr.WriteString(`{"appId":"","mobile":"","pwd":"`)
	jsonStr.WriteString(password)
	jsonStr.WriteString(`","type":""}`)

	url := "http://test.com/u/login"
	req, err := http.NewRequest("POST", url, &jsonStr)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		log.Println(err.Error())
		return ""
	}

	return dat["retCode"].(string)
}
