package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	nbytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int

}

fun New(maxBytes int64, OnEvicted func(string,Value) *Cache){
	return &Cache{
		maxBytes : maxBytes
		ll:		   list.New(),
		cache:	   make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value(*entry)
		return kv.value, true
	}
	return
}
