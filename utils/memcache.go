package utils

import (
	"log"
	memcache "github.com/bradfitz/gomemcache/memcache"
)

var MemcacheClient *memcache.Client

func NewMemcacheClient(c *Config) {
	MemcacheClient = memcache.New(c.Memcache)
	log.Printf("Memcache connected")
}