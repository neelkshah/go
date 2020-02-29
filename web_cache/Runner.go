package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	timeout := time.Second * 500
	cache := CreateCache(timeout)
	for {
		requestUrl := ReadInput(*reader)
		if IsValidUrl(requestUrl) == false{
			fmt.Println("Query seems to be malformed. Please retry :(")
			continue
		}
		response := Get(requestUrl, cache)
		fmt.Println(response)
		//fmt.Println(response.Header["Etag"][0])
	}
}
