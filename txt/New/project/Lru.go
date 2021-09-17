package txt

import "container/list"

/*
LRU(最近最少使用) 算法的设计原则是：
1、如果一个数据在最近一段时间没有被访问到，那么在将来它被访问的可能性也很小。
也就是说，当限定的空间已存满数据时，应当把最久没有被访问到的数据淘汰。

2、基于哈希表和双向链表的LRU算法实现
如果要自己实现一个LRU算法，可以用哈希表加双向链表实现：

设计思路是，使用哈希表存储 key，值为链表中的节点，节点中存储值，双向链表来记录节点的顺序，头部为最近访问节点。

3、LRU算法中有两种基本操作：

get(key)：查询key对应的节点，如果key存在，将节点移动至链表头部。
set(key, value)： 设置key对应的节点的值。如果key不存在，则新建节点，置于链表开头。
如果链表长度超标，则将处于尾部的最后一个节点去掉。如果节点存在，更新节点的值，同时将节点置于链表头部。

leetcode ：
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；
如果关键字不存在，则插入该组「关键字-值」。
当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

*/

type entry struct {
	key   int
	value int
}
type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	ll    *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*list.Element{}, list.New()}
}

func (c *LRUCache) Get(key int) int {
	e := c.cache[key]
	if e == nil {
		return -1
	}
	c.ll.MoveToFront(e)
	return e.Value.(entry).value
}

func (c *LRUCache) Put(key int, value int) {

	if e := c.cache[key]; e != nil {
		e.Value = entry{key, value}
		c.ll.MoveToFront(e) // 刷新缓存使用时间
		return
	}
	c.cache[key] = c.ll.PushFront(entry{key, value})
	if len(c.cache) > c.cap {
		delete(c.cache, c.ll.Remove(c.ll.Back()).(entry).key)
	}
}
