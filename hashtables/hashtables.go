package hashtables

const hashTableSize = 100

type HashTable struct {
	keys   [hashTableSize]int
	values [hashTableSize]interface{}
	size   int
}

func hashFunction(key int) int {
	return key % hashTableSize
}

func NewHashTable() *HashTable {
	return &HashTable{
		keys:   [hashTableSize]int{},
		values: [hashTableSize]interface{}{},
		size:   0,
	}
}

func (ht *HashTable) Insert(key int, value interface{}) {
	index := hashFunction(key)

	for ht.keys[index] != 0 {
		index = (index + 1) % hashTableSize
	}

	ht.keys[index] = key
	ht.values[index] = value
	ht.size++
}

func (ht *HashTable) Get(key int) interface{} {
	index := hashFunction(key)

	for ht.keys[index] != key {
		index = (index + 1) % hashTableSize
	}

	return ht.values[index]
}
