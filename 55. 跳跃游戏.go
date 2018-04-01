给定一个非负整数数组，您最初位于数组的第一个索引处。

数组中的每个元素表示您在该位置的最大跳跃长度。

确定是否能够到达最后一个索引。

示例：
A = [2,3,1,1,4]，返回 true。

A = [3,2,1,0,4]，返回 false。



贪心
关键能不能跨过数值为0的元素


package jump_game

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	step := nums[0]
	for i := 1; i < len(nums); i++ {
		if step > 0 {
			step--
			step = max(step, nums[i])
		} else {
			return false
		}
	}
	return true
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
