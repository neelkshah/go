package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
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
		body, err := ioutil.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%s\n", body)
		body, err = ioutil.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%s\n", body)
		//fmt.Println(response.Header["Etag"][0])
	}
}
