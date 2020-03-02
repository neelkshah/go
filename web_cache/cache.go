package main

import (
	"bufio"
	"net/http"
	"strings"
	"time"
)

func Get(key string, cache *Cache) string {
	// get value after checking if key exists in cache
	return fetch(key, cache)
}

func CreateCache(timeout time.Duration) *Cache {
	// create a new cache
	cache := &Cache{timeout: timeout, hashmap: map[string]http.Response{}, fartherCache: nil}
	go refreshCache(cache)
	return cache
}

func DeleteCache(cache *Cache) {
	// delete a cache completely
	cache = nil
	return
}

func AddCacheLayer(cache *Cache, timeout time.Duration) *Cache {
	// add a new cache layer farther away
	cache.fartherCache = CreateCache(timeout)
	return cache
}

func DeleteCacheLayer(cache *Cache) {
	// delete one layer of a cache
	cache = cache.fartherCache
	return
}

func SetTimeout(cache *Cache, timeout time.Duration) *Cache {
	// set the timeout period for a cache
	cache.timeout = timeout
	return cache
}

func GetTimeout(cache *Cache) time.Duration {
	// set the timeout period for a cache
	return cache.timeout
}

func ReadInput(reader bufio.Reader) string{
	// read inout from the console
	requestUrl, _ := reader.ReadString('\n')
	requestUrl = strings.Split(requestUrl, "\n")[0]
	return requestUrl
}
