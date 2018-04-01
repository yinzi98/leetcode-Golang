/**
给定一个字符串，找出不含有重复字符的 最长子串 的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列 而不是子串。
*/



利用s[left:i+1]来表示s[:i+1]中的包含s[i]的最长子字符串。 location[s[i]]是字符s[i]在s[:i+1]中倒数第二次出现的序列号。 
当left < location[s[i]]的时候，说明字符s[i]出现了两次。需要设置 left = location[s[i]] + 1, 保证字符s[i]只出现一次。

大神思路：利用Location保存字符上次出现的序列号，避免了查询工作。location和Two Sum中的m是一样的作用。
// m 负责保存map[整数]整数的序列号
	m := make(map[int]int, len(nums))



package problem0003

func lengthOfLongestSubstring(s string) int {

	// location[s[i]] == j 表示：

	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]

	// location[s[i]] == -1 表示： s[i] 在s中第一次出现

	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符

	for i := range location {

		location[i] = -1 // 先设置所有的字符都没有见过

	}



	maxLen, left := 0, 0



	for i := 0; i < len(s); i++ {

		// 说明s[i]已经在s[left:i+1]中重复了

		// 并且s[i]上次出现的位置在location[s[i]]

		if location[s[i]] >= left {

			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分

		} else if i+1-left > maxLen {

			// fmt.Println(s[left:i+1])

			maxLen = i + 1 - left

		}

		location[s[i]] = i

	}



	return maxLen

}
