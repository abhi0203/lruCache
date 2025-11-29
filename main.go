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
	hash  map[string]*Node
}

// Functions and Methods to operate the cache.

func NewCache(maxLength int) Cache {
	return Cache{
		queue: NewQueue(maxLength),
		hash:  make(map[string]*Node),
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

func (c Cache) Add(word string) {
	/*
		- Check the length of the cache. If the length is full then we have to first remove.
		- Else we have to add at the first.
		- Also, check if this is the very first element in the queue. In that case add directly.
		- Also update the hashmap data structure each time we add or remove the items.
	*/

	// Queue is empty. Create a new node.
	newNode := &Node{
		Left:  nil,
		Right: nil,
		Val:   word,
	}
	c.hash[word] = newNode
	// If the maxLength has not reached
	if c.queue.Length < c.queue.MaxLength {

		// If the item is first item in the queue
		if c.queue.Length == 0 {
			c.queue.Head = newNode
			c.queue.Length += 1
			c.queue.Tail = newNode
			return
		} else {
			newNode.Right = c.queue.Head
			c.queue.Head.Left = newNode
			c.queue.Head = newNode
			c.queue.Length += 1
		}
	} else {
		// Remove the last Item
		c.Remove()
		// Remove the item from hashmap
		delete(c.hash, word)
		// Add the item to the front
		newNode.Right = c.queue.Head
		c.queue.Head.Left = newNode
		c.queue.Head = newNode
		c.queue.Length += 1
	}

}

func (c Cache) Remove() {

}

func (c Cache) Adjust(word string) {

	if c.queue.Head == c.queue.Tail {
		return
	} else {
		// Create a temp node to point to the current node
		tempNode := c.hash[word]
		// Move the things around
		tempNode.Left.Right = tempNode.Right
		tempNode.Right = c.queue.Head
		tempNode.Left = nil
		c.queue.Head.Left = tempNode
		c.queue.Head = tempNode
	}

}

func (c Cache) Check(word string) {

	// Find if the item is present in the cache
	_, ok := c.hash[word]

	// if Okay call adjust else call add
	if ok {
		c.Adjust(word)
	} else {
		c.Add(word)
	}

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
