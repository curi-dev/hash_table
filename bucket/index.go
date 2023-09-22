package bucket

import (
	"fmt"
	"hashTable/constants"

	hashnode "hashTable/hashNode"
)

type Bucket struct {
	length int
	head   *hashnode.HashNode
	tail   *hashnode.HashNode
}

func (self *Bucket) Insert(newNode *hashnode.HashNode) bool {
	if self.head == nil {
		self.head = newNode
		self.tail = newNode
	} else {
		if self.length < constants.BUCKET_SIZE {
			self.tail.Next = newNode
			self.tail = newNode
		} else {
			return false // only scenario where item could not be inserted
		}
	}

	self.length++

	return true
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
