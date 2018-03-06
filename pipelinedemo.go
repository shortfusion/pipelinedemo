package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Print("Web page returned ")
	fmt.Print(len(body))
	fmt.Println(" bytes")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
