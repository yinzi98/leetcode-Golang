/**
给定一个字符串，找出不含有重复字符的 最长子串 的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列 而不是子串。
*/

package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	runeMap := make(map[rune]int)
	longest := 0
	preLength := 0
	for i, rn := range runes {
		var length int
		if val, ok := runeMap[rn]; !ok || val < i-preLength {
			length = preLength + 1
		} else {
			length = i - val
		}
		if length > longest {
			longest = length
		}
		preLength = length
		runeMap[rn] = i
	}
	return longest
}
