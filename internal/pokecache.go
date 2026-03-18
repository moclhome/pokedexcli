package internal

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	cacheMap  map[string]cacheEntry
	m         *sync.Mutex
	interval  time.Duration
	flagDebug bool

	/*LogAdding   func()
	LogGetting  func(cacheLength int, found bool)
	LogDeleting func(key string)*/
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
	locking(ca, "Add")
	defer unlocking(ca, "Add")
	ca.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	if ca.flagDebug {
		log.Println("Add to cache")
	}
}

func (ca *Cache) Get(key string) ([]byte, bool) {
	locking(ca, "Get")
	defer unlocking(ca, "Get")
	entry, found := ca.cacheMap[key]
	if ca.flagDebug {
		log.Printf("There are %d entries in the cache\n", len(ca.cacheMap))
		if found {
			log.Println("Found in cache")
		}
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
			delete(ca.cacheMap, key)
			if ca.flagDebug {
				log.Printf("Deleting %s from cache", key)
			}
		}
	}
}

func locking(ca *Cache, name string) {
	ca.m.Lock()
	if ca.flagDebug {
		log.Printf("%s locked...\n", name)
	}
}
func unlocking(ca *Cache, name string) {
	ca.m.Unlock()
	if ca.flagDebug {
		log.Printf("%s ...unlocked!\n", name)
	}
}
