package main

import (
	"hash/adler32"
	"time"
)

var (
	hashTable =  map[uint32] cacheEntry {}
)

type cacheEntry struct {
	Data []byte
	Expiration int64
}

func cacheCleaner() {
	dummy := cacheEntry{[]byte{}, 0}
	
	for {
		for key, value := range hashTable {
			if value.Expiration <= time.Seconds() {
				hashTable[key] = dummy, false
			}
		}
		
		time.Sleep(1000000000)
	}
}

func addToCache(hash uint32, value string) {
	hashTable[hash] = cacheEntry{[]byte(value), time.Seconds() + 30}
}

func inArray(hash uint32) bool {
	_, ok := hashTable[hash]
	return ok
}

func hash(toHash string) uint32 {
	return adler32.Checksum([]byte(toHash))
}
