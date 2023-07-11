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

	// если найден обновляем значение элемента и двигаем вперед
	if found {
		item.Value.(*mapItem).value = value

		c.queue.MoveToFront(item)
	} else {
		// если не найден новый элемент в начало очереди
		item = &ListItem{Value: &mapItem{key, value}}

		c.queue.PushFront(item.Value)
	}

	c.items[key] = c.queue.Front()

	// если не найден и превышена капасити удаляем последний элемент из очереди и соответствующий из словаря
	if !found && c.queue.Len() == c.capacity {
		delItem := c.queue.Back()
		delete(c.items, delItem.Value.(*mapItem).key)
		c.queue.Remove(delItem)
	}

	return found
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()

	defer c.Unlock()

	item, found := c.items[key]

	// элемент не найден
	if !found {
		return nil, found
	}

	// элемент найден - двигаем в начало очереди и возвращаем его значение
	c.queue.MoveToFront(item)

	return c.queue.Front().Value.(*mapItem).value, found
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
