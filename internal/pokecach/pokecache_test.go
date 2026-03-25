package pokecach

import (
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	interval := 5 * time.Second
	casesToAddAndGet := map[string]struct {
		key string
		val []byte
	}{
		"simple_finding_1": {
			key: "https/my-test-data.com",
			val: []byte("SomeDataForTesting"),
		},
		"simple_finding_2": {
			key: "https/additional-test-data.uk",
			val: []byte("SomeOtherData, also for testing"),
		},
	}
	casesToOnlyGet := map[string]struct {
		key string
		val []byte
	}{
		"not_there_1": {
			key: "https/this_data_doesnt_exist.de",
			val: []byte("not used"),
		},
		"simple_2": {
			key: "this_neither",
			val: []byte("bla blub"),
		},
	}

	cache := NewCache(interval, true)
	for testName, testData := range casesToAddAndGet {
		cache.Add(testData.key, testData.val)
		actualValue, hasValue := cache.Get(testData.key)
		if !hasValue {
			t.Errorf("Test %s: Added key %s but not found with Get", testName, testData.key)
		}
		if string(actualValue) != string(testData.val) {
			t.Errorf("Test %s: Expected  %s but found %s", testName, string(actualValue), string(testData.val))
		}
	}

	for testName, testData := range casesToOnlyGet {
		_, hasValue := cache.Get(testData.key)
		if hasValue {
			t.Errorf("Test %s: Found key %s without adding it", testName, testData.key)
		}
	}
}

func TestReapLoop(t *testing.T) {
	interval := 5 * time.Millisecond
	waittime := 3 * interval

	casesBeforeDeletion := map[string]struct {
		key string
		val []byte
	}{
		"available_1": {
			key: "https/my-test-data.com",
			val: []byte("SomeDataForTesting"),
		},
		"available_2": {
			key: "https/additional-test-data.uk",
			val: []byte("SomeOtherData, also for testing"),
		},
	}

	cache := NewCache(interval, true)
	for testName, testData := range casesBeforeDeletion {
		cache.Add(testData.key, testData.val)
		_, hasValue := cache.Get(testData.key)
		if !hasValue {
			t.Errorf("Test %s: Expected to find key %s in cache", testName, testData.key)
		}
	}
	time.Sleep(waittime)
	for testName, testData := range casesBeforeDeletion {
		_, hasValue := cache.Get(testData.key)
		if hasValue {
			t.Errorf("Test %s: Key %s should have been deleted fromcache", "not_"+testName, testData.key)
		}
	}
}
