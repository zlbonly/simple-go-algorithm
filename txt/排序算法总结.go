package txt

/*使用Golang实现了以下排序算法:
冒泡排序
选择排序
插入排序
快速排序
归并排序
堆排序

插入排序，冒泡排序，归并排序 是稳定排序。
所谓稳定排序是指 相同值的元素排序后，相对位置不发生改变
*/
/*
1、基本思想：两个数比较大小，较大的数下沉，较小的数冒起来
冒泡排序总的平均时间复杂度为：O(n2)，空间复杂度O1  稳定排序
*/
func Bubble(arr []int) {
	size := len(arr)
	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

/*
	2、思路：每次循环找出最小的数，跟数组第一个数交换顺序，接下来在剩余的数里重复以上逻辑
	时间复杂度（n*n ）,空间复杂度(O1) ，不稳定.
*/
func SelectSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			// 只要找到比要比较的值小的值，就更新min的位置，循环一圈就能找到最小的值的位置
			if arr[j] < arr[i] {
				min = j
			}
		}
		//交换最小值与这一次循环最左边值的位置
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}

/**
3、插入排序，类似扑克牌起牌，将未排序的数据插入到已排序的数据中
 时间复杂度（n*n ）,空间复杂度(O1) ，稳定
*/
func InsertSort(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				//如果要比较的数据小于左边的数据，则交换位置
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

/**
4、快速排序算法
快速排序算法通过多次比较和交换来实现排序，其排序流程如下：
(1)首先设定一个分界值，通过该分界值将数组分成左右两部分。
(2)将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。此时，左边部分中各元素都小于或等于分界值，而右边部分中各元素都大于或等于分界值
(3)然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理
(4)重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。当左、右两个部分各数据排序完成后，整个数组的排序也就完成了.

	1、选取中间元素 作为基准值
	2、左边都比中间元素小，
	3、右边都比中间元素 大
	然后递归
	时间复杂度 nlog（n）,空间复杂度log(n),不稳定排序
*/
func quickSort(arr []int, left int, right int) []int {
	if left < right {
		key := arr[(left+right)/2]
		i, j := left, right
		for {
			if arr[i] < key {
				i++
			}
			if arr[j] > key {
				j--
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		quickSort(arr, left, i-1)
		quickSort(arr, j+1, right)
	}
	return arr
}

/*
5、归并排序：
所谓归并，是指将两个有序数列合并为一个有序数列。
归并排序是利用归并的思想实现的排序方法。如上图。思路比较简单，就是对数组进行不断的分割，分割到只剩一个元素，然后，再两两合并起来。
时间复杂度 nlog（n）,空间复杂度nlog(n),稳定排序
*/
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	i := length / 2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])
	res := merge(left, right)
	return res
}

//合并数组
func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)
	//比较两个数组，谁小把元素值添加到结果集内
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
		} else {
			result = append(result, left[m])
			m++
		}
	}
	//如果有一个数组比完了，另一个数组还有元素的情况，则将剩余元素添加到结果集内
	result = append(result, right[n:]...)
	result = append(result, left[m:]...)
	return result
}

/**
6、堆排序：
	大顶堆：每个结点的值都大于或等于其左右孩子结点的值
	小顶堆：每个结点的值都小于或等于其左右孩子结点的值
	根据对的特性来形成公式就是，节点为i的话
	大顶堆: arr[i]>=arr[2i+1] && arr[i]>=arr[2i+2]
	小顶堆：arr[i]<=arr[2i+1] && arr[i]<=arr[2i+2]
*/
