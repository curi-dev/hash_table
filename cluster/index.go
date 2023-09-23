package cluster

import (
	"fmt"
	"hash/fnv"

	"hashTable/bucket"
	"hashTable/constants"
)

type HashTable struct {
	cluster       []*bucket.Bucket
	max_buckets   int
	bucket_size   int
	size          int
	curr_quantity int
}

func New() *HashTable {
	initialSize := 11
	//structure := make([]*bucket.Bucket, initialSize)
	var structure []*bucket.Bucket

	return &HashTable{max_buckets: constants.MAX_BUCKETS, size: initialSize, cluster: structure, bucket_size: constants.BUCKET_SIZE}
}

func (self *HashTable) InsertItem(key string, value interface{}) {
	if float32(self.curr_quantity) >= float32(self.size*constants.BUCKET_SIZE)*constants.LOAD_FACTOR {
		self.grown()
	}

	hash := self.getHash(key)

	existentBucket := self.cluster[hash] // a container to a linked list
	fmt.Println("existentBucket: ", existentBucket)

	insertionErr := (*existentBucket).InsertOrReplace(key, value)

	if insertionErr != nil { // the only scenario where item is not inserted: bucket limit reached
		self.grown()
		return
	}

	self.curr_quantity++
}

func (self *HashTable) grown() {
	fmt.Println("grown hash table")
}

func (self *HashTable) RetrieveItem(key string) (bool, interface{}) {
	hash := self.getHash(key)
	bucket := self.cluster[hash]

	if bucket == nil {
		return false, nil
	}

	found := bucket.Find(key)

	if found != nil {
		value := found.GetValue()

		return true, value
	}

	return false, nil
}

func (self *HashTable) getHash(key string) uint32 {

	hash := fnv.New32a()
	hash.Write([]byte(key))
	hashValue := hash.Sum32()

	index := hashValue % uint32(self.size) // compression

	fmt.Printf("Key: %d\n ", index)

	return index
}

func (self *HashTable) Print() {

	for i, v := range self.cluster {
		if v != nil {
			v.Print(i)
		}
	}
}

// func (self *HashTable) Print() {
// for _, v := range self.cluster {
// print := fmt.Sprintf("%d: %s", v.key, v.value)
// fmt.Println(print)
// }
// }
