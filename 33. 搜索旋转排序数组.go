假设按照升序排序的数组在预先未知的某个关键点上旋转。
（即 0 1 2 4 5 6 7 将变成 4 5 6 7 0 1 2）。
给你一个目标值来搜索，如果数组中存在这个数则返回它的索引，否则返回 -1。
你可以假设数组中不存在重复。

先假设old = [0, 1, 2, 4, 5, 6, 7]，利用二分查找法，很容易可以5的索引号，当old变换成了new = [4, 5, 6, 7, 0, 1, 2]以后，
同样可以使用二分查找法，因为old和new中的元素有明确的对应关系
old[i] == new[j]，只要i和j满足关系式
j=i+4
if j > len(old) {
    j -= len(old)
}
其中，4 = old中的最大值在new中的索引号 + 1
所以，如果我们手中只有new，我们可以假装自己还是在对old使用二分查找法，当需要获取old[i]的值进行比较判断的时候，使用new[j]的值替代即可。
差不多就是二分查找法的升级版

ackage problem0033



func search(nums []int, target int) int {

	var index, indexOfMax int

	length := len(nums)



	if length == 0 {

		return -1

	}



	// 获取最大值的索引号，以便进行索引号变换

	for indexOfMax+1 < length && nums[indexOfMax] < nums[indexOfMax+1] {

		indexOfMax++

	}



	low, high, median := 0, length-1, 0

	for low <= high {

		median = (low + high) / 2



		// 变换索引号

		index = median + indexOfMax + 1

		if index >= length {

			index -= length

		}

		// 假设nums是由升序切片old转换来的

		// 那么，old[median] == nums[index]



		// 传统二分查找法的比较判断

		// 原先需要old[median]的地方，使用nums[index]即可

		switch {

		case nums[index] > target:

			high = median - 1

		case nums[index] < target:

			low = median + 1

		default:

			return index

		}

	}



	return -1

}
