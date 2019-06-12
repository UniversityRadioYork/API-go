package utils

import (
	"log"
	memcache "github.com/bradfitz/gomemcache/memcache"
)

var MemcacheClient *memcache.Client

func NewMemcacheClient(c *Config) {
	MemcacheClient = memcache.New(c.Memcache)
	err := MemcacheClient.Set(&memcache.Item{Key: "test", Value: []byte("test")})
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("Memcache connected")
	}
	MemcacheClient.Delete("test")
}