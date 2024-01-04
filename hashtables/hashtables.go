package hashtables

import (
	"fmt"
)

type HashTable struct {
	table map[string]interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{
		table: make(map[string]interface{}),
	}
}

func (ht *HashTable) Insert(key string, value interface{}) {
	ht.table[key] = value
}

func (ht *HashTable) Get(key string) interface{} {
	return ht.table[key]
}

func (ht *HashTable) PrintHashTable() {
	fmt.Println(ht.table)
}
