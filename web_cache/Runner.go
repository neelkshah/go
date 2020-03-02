package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	timeout := time.Second * 10
	cache := CreateCache(timeout)
	AddCacheLayer(cache, time.Second * 20)
	for {
		requestUrl := ReadInput(*reader)
		if isValidUrl(requestUrl) == false{
			fmt.Println("Query seems to be malformed. Please retry :(")
			continue
		}
		response := Get(requestUrl, cache)
		fmt.Println(response)
	}
}
