package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"
)

type session struct {
	Key string `json:"session_key"`
}

var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	client = http.Client{
		Jar: jar,
	}
}

func main() {

	jsonFile, err := os.Open("variables.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var ses session
	json.Unmarshal(byteValue, &ses)

	cookie := &http.Cookie{
		Name:   "session",
		Value:  ses.Key,
		MaxAge: 0,
	}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/1/input", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error occured. Error is: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	//log.Printf(sb)
	arr := strings.Fields(sb)
	//log.Println(arr)

	var prev *int = nil
	var counter int = 0

	for i := 0; i < len(arr); i++ {
		cur, _ := strconv.Atoi(arr[i])
		if prev != nil && cur > *prev {
			counter++
		}
		prev = &cur
	}

	fmt.Println(counter)
}
