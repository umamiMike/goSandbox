package main

import "sync"
import "fmt"

type inMemoryHashTable struct {
	m   map[string][]byte
	lck sync.RWMutex
}

func NewInMemoryHashTable() *inMemoryHashTable {
	return &inMemoryHashTable{m: make(map[string][]byte)}
}

func (i *inMemoryHashTable) Get(key string) ([]byte, error) {
	i.lck.RLock()
	defer i.lck.RUnlock()
	val, ok := i.m[key]
	if !ok {
		return nil, fmt.Errorf("the error is value %v", ok)
	}
	return val, nil
}

func (i *inMemoryHashTable) Set(key string, val []byte) error {
	i.lck.Lock()
	defer i.lck.Unlock()
	i.m[key] = val
	return nil

}

func main() {
	db := NewInMemoryHashTable()
	db.Set("smart", []byte("lucky"))

	val, _ := db.Get("smart")
	fmt.Println(string(val))

	errorval, _ := db.Get("snart")
	fmt.Println(string(errorval))
}
