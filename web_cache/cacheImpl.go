package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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


func isValidUrl(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil && u.Scheme != "" && u.Host != ""
}


func fetchFromSource(requestUrl string) http.Response {
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err)
	}
	return *res
}


func fetch(requestUrl string, cache *Cache) string {
	var tempCache = cache
	var delCache *Cache
	start := time.Now().UnixNano()
	var found bool
	var response http.Response
	var depth = 1
	for {
		if cachedValue, ok := tempCache.hashmap[requestUrl]; ok == true {
			fmt.Println("Found in cache!")
			found = true
			response = cachedValue
			delCache = tempCache
			break
		} else {
			if tempCache.fartherCache != nil {
				tempCache = tempCache.fartherCache
				depth += 1
				continue
			} else {
				break
			}
		}
	}
	if found == true {
		delete(delCache.hashmap, requestUrl)
		body, err := ioutil.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		addToCache(requestUrl, response, cache)
		end := time.Now().UnixNano()
		execTime := end - start
		fmt.Println("Fetched in ", execTime, "ns from depth ", depth)
		return string(body)
	}
	fmt.Println("Not found in cache!")
	response = fetchFromSource(requestUrl)
	body, err := ioutil.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	cache.hashmap[requestUrl] = response
	addToCache(requestUrl, response, cache)
	end := time.Now().UnixNano()
	execTime := end - start
	fmt.Println("Fetched in ", execTime, "ns")
	return string(body)
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
	return t.Sub(time.Now())
}


func refreshEntry(urlString string, ttl time.Duration, cache *Cache) {
	for {
		if response, ok := cache.hashmap[urlString]; ok {
			time.Sleep(ttl)
			fmt.Println("Refreshing entry")
			req, _ := http.NewRequest("GET", urlString, nil)
			req.Header.Add("If-None-Match", response.Header["Etag"][0])
			httpClient := http.Client{}
			response, _ := httpClient.Do(req)
			if response.StatusCode == http.StatusNotModified{
				continue
			} else{
				response := fetchFromSource(urlString)
				cache.hashmap[urlString] = response
			}
		} else{
			break
		}
	}
}


func addToCache(urlString string, response http.Response, cache *Cache) {
	ttl := getDuration(response.Header)
	cache.hashmap[urlString] = response
	go refreshEntry(urlString, ttl, cache)
}


func refreshCache(cache *Cache) {
	for {
		time.Sleep(cache.timeout)
		if len(cache.hashmap) > 0 {
			fmt.Printf("Refreshing cache!\n%d keys affected", len(cache.hashmap))
			hashmap := cache.hashmap
			cache.hashmap = map[string]http.Response{}
			if cache.fartherCache == nil {
				return
			} else {
				for key, value := range hashmap {
					addToCache(key, value, cache.fartherCache)
				}
				hashmap = nil
			}
		}
	}
}
