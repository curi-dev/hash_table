package bucket

import (
	"errors"
	"fmt"
	"hashTable/constants"

	hashnode "hashTable/hashNode"
)

type Bucket struct {
	length int
	InUse  bool
	head   *hashnode.HashNode
	tail   *hashnode.HashNode
}

func (self *Bucket) InsertOrReplace(key string, value interface{}) error {
	if !self.InUse {
		newNode := hashnode.New(&hashnode.KeyString{Value: key}, value)

		self.head = newNode
		self.tail = newNode
	} else {

		existentNode := self.Find(key)
		if existentNode != nil {
			existentNode.UpdateValue(value)
		} else {
			if self.length < constants.BUCKET_SIZE {
				newNode := hashnode.New(&hashnode.KeyString{Value: key}, value)

				self.tail.Next = newNode
				self.tail = newNode
			} else {
				return errors.New("No memory avaiable in the bucket")
			}
		}
	}

	self.length++

	self.InUse = true
	return nil
}

func (self *Bucket) Find(key string) *hashnode.HashNode {

	curr_node := self.head
	for {
		keyValue := curr_node.GetKey()
		if keyValue.Value == key {
			return curr_node
		}

		curr_node = curr_node.Next

		if curr_node == nil {
			return nil
		}
	}
}

func (self *Bucket) Print(bucketNumber int) {

	currNode := self.head
	nodeNumber := 0
	for {
		if currNode == nil {
			break
		}

		fmt.Println("Bucket ", bucketNumber, ": ", nodeNumber, "=> ", currNode)

		currNode = currNode.Next

		nodeNumber++
	}
}
