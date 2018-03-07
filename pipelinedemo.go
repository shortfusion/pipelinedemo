package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	website := "https://google.com"
	size := sizeSite(website)
	fmt.Print("Web page returned ")
	fmt.Print(size)
	fmt.Println(" bytes")
}

func sizeSite(testurl string) int {
	resp, err := http.Get(testurl)
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return len(body)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
