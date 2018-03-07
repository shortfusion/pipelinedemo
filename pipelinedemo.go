package main

import (
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	website := "https://google.com/robots.txt"
	size := sizeSite(website)
	allows := countAllow(website)

	fmt.Print("Web page returned ")
	fmt.Print(size)
	fmt.Println(" bytes")

	fmt.Print("Robots.txt contains ")
	fmt.Print(allows)
	fmt.Println(" 'Allows'")

}

func sizeSite(testurl string) int {
	resp, err := http.Get(testurl)
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return len(body)
}

func countAllow(testurl string) int {
	resp, err := http.Get(testurl)
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	r := regexp.MustCompile("Allow")
	index := suffixarray.New(body)
	results := index.FindAllIndex(r, -1)
	return len(results)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
