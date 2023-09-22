package cluster

import (
	"fmt"
	"hash/fnv"

	"hashTable/bucket"
	"hashTable/constants"

	hashnode "hashTable/hashNode"
)

type HashTable struct {
	cluster       []*bucket.Bucket
	max_buckets   int
	bucket_size   int
	size          int
	curr_quantity int
}

func New() *HashTable {
	//structure := [10]string{}
	//var structure []HashItem
	initialSize := 11
	structure := make([]*bucket.Bucket, initialSize)

	return &HashTable{max_buckets: constants.MAX_BUCKETS, size: initialSize, cluster: structure, bucket_size: constants.BUCKET_SIZE}
}

func (self *HashTable) InsertItem(key string, value interface{}) {
	item := hashnode.New(&hashnode.KeyString{Value: key}, value)

	hash := self.getHash(key)

	existentBucket := self.cluster[hash] // linked list

	fmt.Println("existentBucket: ", existentBucket)

	if existentBucket == nil {
		existentBucket = &bucket.Bucket{}
		existentBucket.Insert(item)

		self.cluster[hash] = existentBucket

		self.curr_quantity++

	} else {
		inserted := (*existentBucket).Insert(item)

		if !inserted {
			self.grown()
		} else {
			self.curr_quantity++
		}
	}

	if float32(self.curr_quantity) >= float32(self.size*constants.BUCKET_SIZE)*constants.LOAD_FACTOR {
		self.grown()
	}

	fmt.Println("self.curr_quantity: ", self.curr_quantity)

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
