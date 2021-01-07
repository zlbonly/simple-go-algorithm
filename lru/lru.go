package lru

import "container/list"

/**

1、在这里我们直接使用 Go 语言标准库实现的双向链表list.List。字典的定义是map[string]*list.Element，键是字符串，值是双向链表中对应节点的指针；
2、maxBytes 是允许使用的最大内存，nbytes 是当前已使用的内存；
3、键值对 entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
4、为了通用性，我们允许值是实现了 Value 接口的任意类型，该接口只包含了一个方法 Len() int，用于返回值所占用的内存大小。
 */
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
}

type Entry struct {
	key   string
	value Value
}

type Value interface {
	// 	由于go么有泛型支持，使用interface{}来代替任意的类型。
	Len() int
}

func New(maxBytes int64) *Cache {
	return &Cache{
		maxBytes: maxBytes, // 允许使用最大内存，
		nbytes:   0,        // 当前使用内存
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

/**
1、如果键存在，则更新对应节点的值，并将该节点移到队尾。
2、不存在则是新增场景，首先队尾添加新节点 &entry{key, value}, 并字典中添加 key 和节点的映射关系。
更新 c.nbytes，如果超过了设定的最大值 c.maxBytes，则移除最少访问的节点。

 */
func (c *Cache) Add(key string, value Value) {
	if ele,ok :=c.cache[key];ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	}else {
		ele := c.ll.PushFront(&Entry{key:key,value:value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes <c.nbytes {
		// maxbytes 设置为0 ，代表不对内存大小设限，这里和groupcache 一致，所以不为0时，才会判断是否超过了限制
		c.RemoveOldest();
	}


}

/**
如果键对应的链表节点存在，则将对应节点移动到队尾，并返回查找到的值。
c.ll.MoveToFront(ele)，即将链表中的节点 ele 移动到队尾（双向链表作为队列，队首队尾是相对的，在这里约定 front 为队尾）
 */
func (c *Cache) Get(key string) (value Value, ok bool) {

	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		return kv.value, true
	}
	return
}

/*
	c.ll.Back() 取到队首节点，从链表中删除。
	delete(c.cache, kv.key)，从字典中 c.cache 删除该节点的映射关系。
	更新当前所用的内存 c.nbytes。
 */
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*Entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
