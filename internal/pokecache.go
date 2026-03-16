package internal

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheMap  map[string]cacheEntry
	m         *sync.Mutex
	interval  time.Duration
	flagDebug bool
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration, flagDebug bool) Cache {
	var theMutex sync.Mutex
	theMap := make(map[string]cacheEntry)
	theCache := Cache{
		cacheMap:  theMap,
		m:         &theMutex,
		interval:  interval,
		flagDebug: flagDebug,
	}
	go theCache.ReapLoop(interval)
	return theCache
}

func (ca *Cache) Add(key string, val []byte) {
	if ca.flagDebug {
		fmt.Println("Add to cache")
	}
	locking(ca, "Add")
	defer unlocking(ca, "Add")
	ca.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (ca *Cache) Get(key string) ([]byte, bool) {
	if ca.flagDebug {
		fmt.Printf("There are %d entries in the cache\n", len(ca.cacheMap))
	}
	locking(ca, "Get")
	defer unlocking(ca, "Get")
	entry, found := ca.cacheMap[key]
	if found && ca.flagDebug {
		fmt.Println("Found in cache")
	}
	return entry.val, found
}

func (ca *Cache) ReapLoop(interval time.Duration) {
	timeChannel := time.Tick(interval)
	for range timeChannel {
		ca.reap(interval)
	}
}

func (ca *Cache) reap(interval time.Duration) {
	locking(ca, "reap")
	defer unlocking(ca, "reap")
	for key, value := range ca.cacheMap {
		age := time.Since(value.createdAt)
		if age > interval {
			if ca.flagDebug {
				fmt.Printf("deleting entry %s\n", key)
			}
			delete(ca.cacheMap, key)
		}
	}
}

func locking(ca *Cache, name string) {
	ca.m.Lock()
	if ca.flagDebug {
		fmt.Printf("%s locked...\n", name)
	}
}
func unlocking(ca *Cache, name string) {
	ca.m.Unlock()
	if ca.flagDebug {
		fmt.Printf("%s ...unlocked!\n", name)
	}
}
