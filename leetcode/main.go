package main

import (
	"math"
	"sort"
)

func main() {

}

/*
*11. 盛最多水的容器
 */
func maxArea(height []int) int {
	var ans = 0
	left := 0
	right := len(height) - 1

	for left < right {
		if height[left] < height[right] {
			ans = max(ans, height[left]*(right-left))
			left++
		} else {
			ans = max(ans, height[right]*(right-left))
			right--
		}
	}
	return ans
}

/*
*三数之和
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		target := -nums[i]

		if target < 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			for left < right {
				if nums[left]+nums[right] == target {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					} // 跳过重复数字
					for right--; right > left && nums[right] == nums[right+1]; right-- {
					} // 跳过重复数字

				} else if nums[left]+nums[right] > target {
					right--
				} else {
					left++
				}
			}
		}

	}
	return res
}

/*
*长度最小的子数组
 */
func minSubArrayLen(target int, nums []int) int {
	var ans = math.MaxInt
	var sum = 0
	var left = 0
	var right = 0
	var size = len(nums)

	for right < size {
		sum += nums[right]

		for left < right && sum-nums[left] >= target {
			sum -= nums[left]
			left++
		}

		if sum >= target {
			ans = min(ans, right-left+1)
		}
		right++
	}
	if ans == math.MaxInt {
		return 0
	} else {
		return ans
	}
}
