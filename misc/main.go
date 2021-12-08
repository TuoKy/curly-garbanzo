package misc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

type session struct {
	Key string `json:"session_key"`
}

func makeClient() http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}
	client := http.Client{
		Jar: jar,
	}
	return client
}

func makeRequest(client http.Client, day int) string {
	// TODO: fix this. Other modules clling this may have to pass parameter or something
	jsonFile, err := os.Open("../misc/variables.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var ses session
	json.Unmarshal(byteValue, &ses)

	cookie := &http.Cookie{
		Name:   "session",
		Value:  ses.Key,
		MaxAge: 0,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day), nil)
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

	return string(body)
}

func GetRawInput(day int) string {
	return makeRequest(makeClient(), day)
}

// Returns input as an array with input string separeted by whitespaces
func GetInput(day int) []string {
	return strings.Fields(makeRequest(makeClient(), day))
}
