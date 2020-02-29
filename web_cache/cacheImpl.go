package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	timeLayout = "Mon, 02-Jan-06 15:04:05 MST"
)


type Cache struct{
	hashmap map[string]http.Response
	timeout time.Duration
	fartherCache *Cache
}


func IsValidUrl(str string) bool {
	u, err := url.Parse(str)
	fmt.Println(err)
	return err == nil && u.Scheme != "" && u.Host != ""
}


func getDuration(headers http.Header) time.Duration{
	params := strings.Split(headers["Set-Cookie"][0],";")
	var expDate string
	for i:=0; i<len(params); i++ {
		if strings.Contains(params[i], "expires"){
			expDate = strings.Split(params[i],"=")[1]
		}
	}
	t, _ := time.Parse(timeLayout, expDate)
	fmt.Println(t.Sub(time.Now()))
	return t.Sub(time.Now())
}


//func addToCache(urlString string, response http.Response, cache Cache) {
//	duration := getDuration(response.Header)
//	cache.hashmap[urlString] = response
//	//go refreshEntry(urlString, duration, cache)
//}
//
//
//func refreshEntry(urlString string, duration time.Duration, cache Cache) {
//	for {
//		if response, ok := cache.hashmap[urlString]; ok {
//			time.Sleep(duration)
//			req, _ := http.NewRequest("GET", urlString, nil)
//			req.Header.Add("If-None-Match", response.Header["Etag"][0])
//			response, _ := http.Client{}.Do(req)
//			if response.StatusCode == http.StatusNotModified{
//
//			}
//		} else{
//			break
//		}
//	}
//}


func fetch(requestUrl string, cache *Cache) http.Response {
	if cachedValue, ok := cache.hashmap[requestUrl]; ok == true {
		fmt.Println("Found in cache!")
		return cachedValue
	} else {
		fmt.Println("Not found in cache!")
		res := fetchFromSource(requestUrl)
		cache.hashmap[requestUrl] = res
		fmt.Println(res.Status)
		return res
	}
}


func fetchFromSource(requestUrl string) http.Response {
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err)
	}
	return *res
}
