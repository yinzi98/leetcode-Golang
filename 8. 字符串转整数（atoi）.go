实现 atoi，将字符串转为整数。

提示：仔细考虑所有输入情况。如果你想挑战自己，请不要看下面并自己考虑所有可能的输入情况。

说明：这题解释的比较模糊（即没有指定输入格式）。你得事先汇集所有的输入情况。

 

atoi的要求：

这个函数需要丢弃之前的空白字符，直到找到第一个非空白字符。之后从这个字符开始，选取一个可选的正号或负号后面跟随尽可能多的数字，并将其解释为数字的值。

字符串可以在形成整数的字符后包括多余的字符，将这些字符忽略，这些字符对于函数的行为没有影响。

如果字符串中的第一个非空白的字符不是有效的整数，或者没有这样的序列存在，字符串为空或者只包含空白字符则不进行转换。

如果不能执行有效的转换，则返回 0。如果正确的值超过的可表示的范围，则返回 INT_MAX（2147483647）或 INT_MIN（-2147483648）。

 水题水题水题
 
 package string_to_integer

// just consider 32-bit, fucking leetcode.com's test cases
const (
	MAX int32 = 1<<31 - 1
	MIN int32 = -1*MAX - 1
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	runes := []rune(str)
	var ret int32
	var neg bool
	i := 0
	for ; runes[i] == rune(' '); i++ {
	}
	if runes[i] == rune('-') || runes[i] == rune('+') {
		neg = runes[i] == rune('-')
		i++
	}

	for ; i < len(runes); i++ {
		if val, ok := toInt(runes[i]); !ok {
			break
		} else {
			if neg {
				if -ret < (MIN+val)/10 {
					return int(MIN)
				}
			} else {
				if ret > (MAX-val)/10 {
					return int(MAX)
				}
			}
			ret = 10*ret + val
		}
	}
	if neg {
		return int(-ret)
	}
	return int(ret)
}

func toInt(r rune) (int32, bool) {
	ret := int32(r - rune('0'))
	if ret >= 0 && ret <= 9 {
		return ret, true
	}
	return 0, false
}
