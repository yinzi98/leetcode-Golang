给一个只包含 '(' 和 ')' 的字符串，找出最长的有效（正确关闭）括号子串的长度。
对于 "(()"，最长有效括号子串为 "()" ，它的长度是 2。
另一个例子 ")()())"，最长有效括号子串为 "()()"，它的长度是 4。

1.记录每个符号的状态，(一律对应于0；)如果能够和前面的配上对，就记录为2，否则，记录为0
) ( ( ) ( ) ) ) ( ( ( ( ( ) ) ) ) (
形成记录
0 0 0 2 0 2 2 0 0 0 0 0 0 2 2 2 2 0
2.从左往右检查record，如果record[i]==2，record[j]==0，且record[j+1:i]中没有0，则record[i]=1,record[j]=1，那么，上面的record就变成了
0 1 1 1 1 1 1 0 0 1 1 1 1 1 1 1 1 0
3.统计record中，最多的连续为1的次数，就是结果。

package problem0032



func longestValidParentheses(s string) int {

	var left, max, temp int

	record := make([]int, len(s))



	// 统计Record

	for i, b := range s {

		if b == '(' {

			left++

		} else if left > 0 {

			left--

			record[i] = 2

		}

	}



	// 修改record

	for i := 0; i < len(record); i++ {

		if record[i] == 2 {

			j := i - 1

			for record[j] != 0 {

				j--

			}

			record[i], record[j] = 1, 1

		}

	}



	// 统计结果

	for _, r := range record {

		if r == 0 {

			temp = 0

			continue

		}



		temp++

		if temp > max {

			max = temp

		}

	}



	return max

}
