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

func(c *Cache)RemoveOldest(){
	ele := c.ll.back();
	if ele != nil {
		c.ll.remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache,kv.key)
		c.nbytes -= int64(len(kv.key)) +  int64(kv.Value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key,kv.Value)
		}
	}
}
//新增 修改
func(c *Cache) Add (key string,value Value){
	if ele ,ok := c.cache[key]; ok{
		c.ll.MoveToFront(ele)
		kv := ele.Vlaue.(*entry)
		c.nbytes  += int64(value.Len()) - int(kv.value.Len())
		kv.value = value
	}else{
		ele := c.ll.PushFront(&entry{key,value})
		c.cache[key] = value
		c.nbytes += int64(value.Len()) + int64(len(kv.key))
	}
	for c.nbytes != 0 && c.maxBytes < c.nbytes{
		c.RemoveOldest()
	}
}
