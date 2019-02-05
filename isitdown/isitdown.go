package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]

	for i := 0; i < len(urls); i++ {
		fmt.Println("Connecting to website", urls[i], "returned", isitdown(urls[i]))

	}
}

func isitdown(url string) bool {
	rsp, err := http.Get(url)
	if err != nil {
		return false
	}
	rsp.Body.Close()
	return true
}
