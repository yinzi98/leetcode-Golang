/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 长度最长为1000。

示例:

输入: "babad"

输出: "bab"

注意: "aba"也是有效答案
 

示例:

输入: "cbbd"

输出: "bb"
*/


/**
思路
1:用Map保存字符和最后一次出现的下标
2:并记录出现最大length的起始和结束的下标，用于截取字符串
3:根据当前字符在是否出现以及上次出现的位置决定是拼接子字符串还是重新起一个子字符串
时间复杂度，空间复杂度都取决于字符串长度  且都是O（n）
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
