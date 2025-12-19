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
	queue *Queue
	hash  map[string]*Node
}

// Functions and Methods to operate the cache.

func NewCache(maxLength int) Cache {
	fmt.Println("NewCache: Creating a new Cache")
	return Cache{
		queue: NewQueue(maxLength),
		hash:  make(map[string]*Node),
	}
}

func NewQueue(maxLength int) *Queue {
	fmt.Println("NewQueue: Creating a new queue")
	queue := Queue{
		Head:      nil,
		Tail:      nil,
		Length:    0,
		MaxLength: maxLength,
	}
	return &queue
}

func (c Cache) Add(word string) {
	/*
		- Check the length of the cache. If the length is full then we have to first remove.
		- Else we have to add at the first.
		- Also, check if this is the very first element in the queue. In that case add directly.
		- Also update the hashmap data structure each time we add or remove the items.
	*/
	fmt.Println("Add: Word is not present and we need to add the word ", word)
	fmt.Println("Add: Queue Object is ", c.queue)
	newNode := &Node{
		Left:  nil,
		Right: nil,
		Val:   word,
	}
	c.hash[word] = newNode
	// If the maxLength has not reached
	if c.queue.Length < c.queue.MaxLength {
		fmt.Println("Add: The max length of the queue has not reached.")
		// If the item is first item in the queue
		if c.queue.Length == 0 {
			fmt.Println("Add: The queue length is 0")
			c.queue.Head = newNode
			c.queue.Length += 1
			c.queue.Tail = newNode
			fmt.Println("Add: Queue object is ", c.queue)
		} else {
			fmt.Println("The queue length >0 but not reached max")
			newNode.Right = c.queue.Head
			c.queue.Head.Left = newNode
			c.queue.Head = newNode
			c.queue.Length += 1
			fmt.Println("Add: The value in the queue is ", c.queue.Head.Val)
			fmt.Println("Add: Length of the queue is ", c.queue.Length)
			fmt.Println("Add: Max lenght of the queue is ", c.queue.MaxLength)
		}
	} else {
		// Remove the last Item
		c.Remove()
		// Add the item to the front
		newNode.Right = c.queue.Head
		c.queue.Head.Left = newNode
		c.queue.Head = newNode
		c.queue.Length += 1
	}

}

func (c Cache) Remove() {
	fmt.Println("Remove: The queue is full and we need to remove the last item.")
	if c.queue.Head == c.queue.Tail {
		// Only one item left in the queue
		fmt.Println("Remove: Only 1 item in the Queue")
		// Remove the item from hashmap
		delete(c.hash, c.queue.Tail.Val)
		c.queue.Head = nil
		c.queue.Tail = nil
		c.queue.Length -= 1
		fmt.Println("Remove: New Hash Map is ")
		fmt.Println(c.hash)
	} else {
		fmt.Println("Remove: More than 1 items in the queue")
		current := c.queue.Tail
		// Remove the item from hashmap
		if _, exists := c.hash[c.queue.Tail.Val]; exists {
			delete(c.hash, c.queue.Tail.Val)
			fmt.Println("Remove: New Hash Map is ")
			fmt.Println(c.hash)
		} else {
			fmt.Println("Remove: Word is not present in the hashmap ", c.queue.Tail.Val)
		}
		c.queue.Tail = current.Left
		c.queue.Tail.Right = nil
		current = nil
		c.queue.Length -= 1

	}

}

func (c Cache) Adjust(word string) {
	fmt.Println("Adjust: Word is present and now we need to adjust the cache ", word)
	tempNode := c.hash[word]
	if c.queue.Head == tempNode {
		fmt.Println("Adjust: Current node is already at the head of the queue. Nothing to be done.")
		tempNode = nil
		return
	} else if c.queue.Tail == tempNode {
		fmt.Println("Adjust: Current Node is tail node and this needs to be adjusted.")
		// Remove tail node safely
		tempNode.Left.Right = nil
		c.queue.Tail = tempNode.Left
		tempNode.Left = nil
		// Add the tempnode to the head
		tempNode.Right = c.queue.Head
		c.queue.Head.Left = tempNode
		c.queue.Head = tempNode
		return
	} else {
		fmt.Println("Adjust: Current node is some node in the middle and needs to be adjusted.")
		// Remove the node and adjust right and left of previous and last node
		tempNode.Right.Left = tempNode.Left
		tempNode.Left.Right = tempNode.Right
		// Move the temp node to the right.
		tempNode.Left = nil
		tempNode.Right = c.queue.Head
		c.queue.Head.Left = tempNode
		c.queue.Head = tempNode
		return
	}

}

func (c Cache) Check(word string) {

	fmt.Println("Check: Checking if the word is present in the cache ", word)
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
	fmt.Println("Time to display the items from the cache.")
	fmt.Println(c.hash)
	current := c.queue.Head
	if current != nil {
		fmt.Println("Current value is ", current.Val)
		for current.Right != nil {
			fmt.Println(current.Val)
			current = current.Right
		}
		fmt.Println(current.Val)
	} else {
		fmt.Println("Current is nil")
	}

}

func main() {
	fmt.Println("STARTING THE CACHE")
	cacheLen := 3
	if cacheLen < 2 {
		panic("Cache length must be greater than 1")
	}
	cache := NewCache(cacheLen)
	// for _, word := range []string{"Abhiram", "Sara", "Neha", "Alka", "Srivatsa", "Saatvika", "Vedanshi", "Vidhya Ji"}
	for _, word := range []string{"Abhiram", "Sara", "Neha", "Abhiram", "Alka", "Srivatsa", "Alka", "Saatvika", "Vedanshi", "Vidhya Ji"} {
		fmt.Println("**********************")
		fmt.Println("main: Time for cache access.")
		cache.Check(word)
		cache.Display()
		fmt.Println("??????????????????????")
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
