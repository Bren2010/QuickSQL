package main

import (
	"hash/adler32"
	"time"
)

// Start data declaration
var (
	hashTable =  map[uint32] cacheEntry {}
)

type cacheEntry struct {
	Data []byte
	Expiration int64
}
// End

// Start GoRoutines
func cacheCleaner() {
	// Convert the seconds specified in the config to nanoseconds.
	realPeriod := int64(cachePeriod * 1000000000)
	
	// A nil cache entry to has something to put in the delete statement.
	dummy := cacheEntry{[]byte{}, 0}
	
	for {
		for key, value := range hashTable { // foreach
			if value.Expiration <= time.Seconds() {
				// Remove value from hash table.
				hashTable[key] = dummy, false
			}
		}
		
		// Sleep for X seconds.
		time.Sleep(realPeriod)
	}
}
// End GoRoutines

// Start functions
// Inserts the cache entry into the in-memory hash table.
func addToCache(hash uint32, value []byte) {
	hashTable[hash] = cacheEntry{value, time.Seconds() + cacheAge}
}

// Checks if the given hash has a entry in the hash table.
func inArray(hash uint32) bool {
	_, ok := hashTable[hash]
	return ok
}

// Hashes the string with the current hash method.
func hash(toHash string) uint32 {
	return adler32.Checksum([]byte(toHash))
}
// End functions
