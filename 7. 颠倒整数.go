给定一个范围为 32 位 int 的整数，将其颠倒。

例 1:

输入: 123
输出:  321
 

例 2:

输入: -123
输出: -321
 

例 3:

输入: 120
输出: 21
 

注意:

假设我们的环境只能处理 32 位 int 范围内的整数。根据这个假设，如果颠倒后的结果超过这个范围，则返回 0。

//水题

package reverse_integer

// fucking leetcode's int32 overflow test cases
const (
	MAX int32 = 1<<31 - 1
	MIN       = -MAX - 1
)

func reverse(x int) int {
	var sign int32 = 1
	tmp := x
	if x < 0 {
		sign = -1
		tmp = -1 * x
	}
	list := []int32{}
	for tmp > 9 {
		list = append(list, int32(tmp%10))
		tmp = tmp / 10
	}
	list = append(list, int32(tmp))
	var ret int32
	for i := 0; i < len(list); i++ {
		if sign == 1 {
			if ret > (MAX-list[i])/10 {
				return 0
			}
		} else {
			if -ret < (MIN+list[i])/10 {
				return 0
			}
		}
		ret = 10*ret + list[i]
	}
	return int(ret * sign)
}
