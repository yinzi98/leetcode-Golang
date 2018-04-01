给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
示例 1 :

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :

输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :

输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。





NYOJ上有道类似的贪心题 但当时没用结构体排序

package problem0402

func removeKdigits(num string, k int) string {
	// 返回值的长度
	digits := len(num) - k
	stack := make([]byte, len(num))
	top := 0

	for i := range num {
		// 在还能删除的前提下
		// 从上往下，删除 stack 中所有比 num[i] 大的数
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}
		stack[top] = num[i]
		top++
	}

	// 处理开头的　'0'
	i := 0
	for i < digits && stack[i] == '0' {
		i++
	}

	if i == digits {
		return "0"
	}
	return string(stack[i:digits])
}
