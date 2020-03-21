package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	timeout := time.Second * 29
	cache := CreateCache(timeout)
	AddCacheLayer(cache, time.Second * 43)
	for {
		fmt.Println("> ")
		requestUrl := ReadInput(*reader)
		if isValidUrl(requestUrl) == false {
			fmt.Println("Query seems to be malformed. Please retry :(")
			continue
		}
		_ = Get(requestUrl, cache)
		//fmt.Println(response)
	}
}
