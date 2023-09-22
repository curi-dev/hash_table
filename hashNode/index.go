package hashnode

type Key struct {
	Value *int
}

type KeyString struct {
	Value string
}

type HashNode struct {
	//key   Key
	key   *KeyString
	value interface{}
	Next  *HashNode
}

func New(key *KeyString, value interface{}) *HashNode {
	return &HashNode{key: key, value: value}
}

func (self *HashNode) GetValue() interface{} {
	return self.value
}

func (self *HashNode) GetKey() *KeyString {
	return self.key
}

// func (self *HashNode) Next() *HashNode {
// 	return self.next
// }
