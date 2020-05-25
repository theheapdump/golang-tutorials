package utils

import (
	"fmt"
)

type MemCache struct {
	Cache map[string]string
}

func (cache MemCache) AddToCache(key string, value string) {
	fmt.Println("Add To Cache : ", key, "[", value, "]")
	cache.Cache[key] = value
	fmt.Println("Contents of the cache ", cache.Cache)
}

func (cache MemCache) Get(key string) string {
	fmt.Println("Get From Cache : ", key)
	return cache.Cache[key]
}

func (cache MemCache) Remove(key string) {
	fmt.Println("Remove From Cache : ", key)
	delete(cache.Cache, key)
	fmt.Println("Contents of the cache ", cache.Cache)
}

func (cache MemCache) KeyExists(key string) bool {
	fmt.Println("Check if key ", key, " existis in Cache")
	_, exists := cache.Cache[key]
	return exists
}
