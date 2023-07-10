package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}
type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type mapItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()

	defer c.Unlock()

	item, found := c.items[key]

	if found {
		item.Value.(*mapItem).value = value
		c.queue.PushFront(item)
		return found
	}

	item = &ListItem{Value: &mapItem{key, value}}

	if c.queue.Len() == c.capacity {
		delItem := c.queue.Back()
		delete(c.items, delItem.Value.(*mapItem).key)
		c.queue.Remove(delItem)
	}

	c.items[key] = item
	c.queue.PushFront(item)

	return found
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()

	defer c.Unlock()

	item, found := c.items[key]

	if !found {
		return nil, found
	}

	c.queue.PushFront(item)

	return item.Value.(*mapItem).value, found
}

func (c *lruCache) Clear() {
	c.Lock()

	defer c.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
