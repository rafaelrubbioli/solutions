package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	Capacity int
	Cache    map[int]*list.Element
	LRU      *list.List
}

type Item struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	l := list.List{}
	return LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*list.Element),
		LRU:      l.Init(),
	}
}

func (c *LRUCache) Get(key int) int {
	item, ok := c.Cache[key]
	if !ok {
		return -1
	}

	c.LRU.MoveToFront(item)
	return item.Value.(Item).Value
}

func (c *LRUCache) Put(key int, value int) {
	if e, ok := c.Cache[key]; ok {
		e.Value = Item{Key: key, Value: value}
		c.LRU.MoveToFront(e)
		return
	}

	if len(c.Cache) >= c.Capacity {
		delete(c.Cache, c.LRU.Back().Value.(Item).Key)
		c.LRU.Remove(c.LRU.Back())
	}

	e := c.LRU.PushFront(Item{Key: key, Value: value})
	c.Cache[key] = e
}

func main() {
	fmt.Println("Case 1:")
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	fmt.Println("")

	fmt.Println("Case 2:")
	cache = Constructor(2)
	fmt.Println(cache.Get(2))
	cache.Put(2, 6)
	fmt.Println(cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	fmt.Println("")

	fmt.Println("Case 1:")
	cache = Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
