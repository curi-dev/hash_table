package main

import (
	"fmt"
	"hashTable/cluster"
)

func main() {

	fmt.Println("Hash table!")

	hashTable := cluster.New()

	hashTable.InsertItem("banana", "B")
	hashTable.InsertItem("apple", "A")
	hashTable.InsertItem("pineapple", "D")
	hashTable.InsertItem("orange", "D")
	hashTable.InsertItem("cherry", "D")
	hashTable.InsertItem("peach", "D")
	hashTable.InsertItem("plum", "D")
	hashTable.InsertItem("lemon", "D")
	hashTable.InsertItem("kiwi", "D")
	hashTable.InsertItem("papaya", "D")
	hashTable.InsertItem("fig", "D")

	hashTable.Print()

	_, value := hashTable.RetrieveItem("banana")

	fmt.Println("value found: ", value)

}
