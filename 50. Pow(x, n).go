实现 pow(x, n)。

示例 1:
输入: 2.00000, 10
输出: 1024.00000
示例 2:
输入: 2.10000, 3
输出: 9.26100






指数是整数，所以，可以利用乘法计算幂
用o(N)的算法会浪费很多时间，o(lgN)的算法要快的多。





package problem0050

func myPow(x float64, n int) float64 {

	if n < 0 {

		return 1.0 / pow(x, -n)

	}



	return pow(x, n)

}



func pow(x float64, n int) float64 {

	if x == 0 {

		return 0

	}



	if n == 0 {

		return 1

	}



	res := pow(x, n>>1)

	if n&1 == 0 {

		return res * res

	}



	return res * res * x

}
