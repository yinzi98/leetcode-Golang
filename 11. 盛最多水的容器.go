给定 n 个正整数 a1，a2，...，an，其中每个点的坐标用（i, ai）表示。 
画 n 条直线，使得线 i 的两个端点处于（i，ai）和（i，0）处。请找出其中的两条直线，使得他们与 X 轴形成的容器能够装最多的水。

真·水题

package container_with_most_water

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max, left, right := 0, 0, len(height)-1
	for left < right {
		area := minHeight(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			for left < len(height)-1 {
				left++
				if height[left] > height[left-1] {
					break
				}
			}
		} else {
			for right > 0 {
				right--
				if height[right] > height[right+1] {
					break
				}
			}
		}
	}
	return max
}

func minHeight(x, y int) int {
	if x < y {
		return x
	}
	return y
}
