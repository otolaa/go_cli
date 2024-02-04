package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

const serverPort = 80
const urlApi = "https://jsonplaceholder.typicode.com"
const postsApi = "/posts"

type post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func main() {
	// CLI color
	fmt.Println("Hi 世!? 🤖")
	fmt.Println(math.Pi)

	requestURL := fmt.Sprintf("%s%s", urlApi, postsApi)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	// fmt.Printf("client: response body: %s\n", resBody)

	var astros []post
	// var astros []interface{}
	jsonErr := json.Unmarshal(resBody, &astros)
	if jsonErr != nil {
		fmt.Printf("client: could not read response body: %s\n", jsonErr)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("%sid: %d%s", "\033[34m", astros[0].Id, "\033[0m"))
	fmt.Println(fmt.Sprintf("%suserId: %d%s", "\033[31m", astros[0].UserId, "\033[0m"))
	fmt.Println(fmt.Sprintf("%stitle: %s%s", "\033[33m", astros[0].Title, "\033[0m"))
	fmt.Println(fmt.Sprintf("%sbody: %s%s", "\033[35m", astros[0].Body, "\033[0m"))
}
