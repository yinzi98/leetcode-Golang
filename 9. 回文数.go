判断一个整数是否是回文数。不能使用辅助空间。

点击查看提示

一些提示:

负整数可以是回文数吗？（例如 -1）

如果你打算把整数转为字符串，请注意不允许使用辅助空间的限制。

你也可以考虑将数字颠倒。但是如果你已经解决了 “颠倒整数” 问题的话，就会注意到颠倒整数时可能会发生溢出。你怎么来解决这个问题呢？

本题有一种比较通用的解决方式。


不开新空间！！！！！！
所以暴力呗

package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	bits := 1
	tmp := x
	for {
		if tmp < 10 {
			break
		}
		tmp = tmp / 10
		bits++
	}
	i, j := 1, bits
	for i < j {
		left := (x / divisor(i)) % 10
		right := (x / divisor(j)) % 10
		if left != right {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func divisor(bits int) int {
	ret := 1
	for i := 0; i < bits-1; i++ {
		ret = ret * 10
	}
	return ret
}
