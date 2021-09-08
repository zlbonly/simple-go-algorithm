package hash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

// 定义函数类型Hash，采取依赖注入的方式，允许用于替换成自定义的Hash函数，也方便测试替换，默认使用crc32.ChecksumIEEE 算法
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int            // 虚拟节点倍数
	keys     []int          // hash环keys
	hashMap  map[int]string // 虚拟节点与真实节点映射表（键是虚拟节点的哈希值，值是真实节点的名称）
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

/**
Add 函数允许传入 0 或 多个真实节点的名称。
对每一个真实节点 key，对应创建 m.replicas 个虚拟节点，虚拟节点的名称是：strconv.Itoa(i) + key，即通过添加编号的方式区分不同虚拟节点。
使用 m.hash() 计算虚拟节点的哈希值，使用 append(m.keys, hash) 添加到环上。
在 hashMap 中增加虚拟节点和真实节点的映射关系。
最后一步，环上的哈希值排序。


*/
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

/**
选择节点就非常简单了，第一步，计算 key 的哈希值。
第二步，顺时针找到第一个匹配的虚拟节点的下标 idx，从 m.keys 中获取到对应的哈希值。如果 idx == len(m.keys)，说明应选择 m.keys[0]，因为 m.keys 是一个环状结构，所以用取余数的方式来处理这种情况。
第三步，通过 hashMap 映射得到真实的节点。


*/
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	fmt.Print(idx)
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

// Adds 8, 18, 28

/*2.2 数据倾斜问题
如果服务器的节点过少，容易引起 key 的倾斜。例如上面例子中的 peer2，peer4，peer6 分布在环的上半部分，下半部分是空的。那么映射到环下半部分的 key 都会被分配给 peer2，key 过度向 peer2 倾斜，缓存节点间负载不均。

为了解决这个问题，引入了虚拟节点的概念，一个真实节点对应多个虚拟节点。

假设 1 个真实节点对应 3 个虚拟节点，那么 peer1 对应的虚拟节点是 peer1-1、 peer1-2、 peer1-3（通常以添加编号的方式实现），其余节点也以相同的方式操作。

第一步，计算虚拟节点的 Hash 值，放置在环上。
第二步，计算 key 的 Hash 值，在环上顺时针寻找到应选取的虚拟节点，例如是 peer2-1，那么就对应真实节点 peer2。
虚拟节点扩充了节点的数量，解决了节点较少的情况下数据容易倾斜的问题。而且代价非常小，只需要增加一个字典(map)维护真实节点与虚拟节点的映射关系即可。
*/
/*func(key []byte) uint32 {
	i, _ := strconv.Atoi(string(key))
	return uint32(i)
}*/
