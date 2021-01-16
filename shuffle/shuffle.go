package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
费雪耶兹 洗牌算法

主要思路为每次随机挑选一个值，放在数组末尾，然后在n-1个元素的数组中再随机挑选一个值，放在数组末尾，以此类推

*/

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle1(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))

		slice[a], slice[b] = slice[b], slice[a]
	}
}

func shuffle2(indexes []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}

func main() {
	var cnt1 = map[int]int{}

	for i := 0; i < 100000; i++ {
		var s1 = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle1(s1)

		cnt1[s1[0]]++
	}

	var cnt2 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var s1 = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle2(s1)
		cnt2[s1[0]]++
	}

	fmt.Println(cnt1, "\n", cnt2)
}
