package main

import (
	"fmt"
)

// Important data structures for the cache

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head      *Node
	Tail      *Node
	Length    int
	MaxLength int
}

type Cache struct {
	queue Queue
	hash  Hash
}

type Hash map[string]*Node

// Functions and Methods to operate the cache.

func NewCache(maxLength int) Cache {
	return Cache{
		queue: NewQueue(maxLength),
		hash:  Hash{},
	}
}

func NewQueue(maxLength int) Queue {
	queue := Queue{
		Head:      nil,
		Tail:      nil,
		Length:    0,
		MaxLength: maxLength,
	}
	return queue
}

func (c Cache) Add(item string) {
	// Check if the item is already present in the Cache
	_, ok := Hash[item]

}

func (c Cache) Remove() {

}

func (c Cache) Adjust() {

}

func (c Cache) Check(word string) {

}

func (c Cache) Display() {

}

func main() {
	fmt.Println("STARTING THE CACHE")
	cache := NewCache(5)
	for _, word := range []string{"Abhiram", "Sara", "Neha", "Alka", "Srivatsa", "Saatvika", "Vedanshi", "Vidhya Ji"} {
		cache.Check(word)
		cache.Display()
	}
}

/*
Steps to the Algorithm
- I create an empty cache element.
- Empty cache has empty queue and empty hashmap.
- Below are the functionalities
- ADD
	- We first check the hashmap if the item is present in the cache.
		- If it is present, we get the value and call ADJUST to bring item to the first.
	- If the item is not present in the hashmap, we need to add it.
		- If the lenght of the queue has not reached, then we will add the item to the head.
			- Add the item in the hashmap as well.
		- Else if the length of the queue has reached, then we REMOVE the last item and ADD the new item at the head.

- REMOVE
- Adjust



*/
