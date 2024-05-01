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

/*
*无重复字符的最长子串
 */
func lengthOfLongestSubstring(s string) int {

	len := len(s)
	left, right := 0, 0
	var arr = [128]int{}
	var ans = math.MinInt
	for right < len {

		for left < right && arr[s[right]] != 0 {
			arr[s[left]] = 0
			left++
		}

		ans = max(ans, right-left+1)
		arr[s[right]] = 1
		right++
	}
	if ans == math.MinInt {
		return 0
	} else {
		return ans
	}
}

/*
*串联所有单词的子串
 */
func findSubstring(s string, words []string) []int {

	len1 := len(s)
	len2 := len(words)
	len3 := len(words[0])

	var ans = []int{}
	var hashMap = make(map[string]int)

	for i := 0; i < len2; i++ {
		hashMap[words[i]]++
	}

	for i := 0; i < len3; i++ {

		var temp = make(map[string]int)
		for j := i; j+len3 <= len1; j += len3 {
			cur := s[j : j+len3]
			temp[cur]++

			if j >= i+len2*len3 {
				var index = j - len2*len3
				prev := s[index : index+len3]

				if temp[prev] == 1 {
					delete(temp, prev)
				} else {
					temp[prev]--
				}
				if !mapsEqual(temp, hashMap) {
					continue
				}
			}
			if temp[cur] != hashMap[cur] {
				continue
			}

			if mapsEqual(temp, hashMap) {
				ans = append(ans, j-(len2-1)*len3)
			}
		}
	}
	return ans
}

func mapsEqual(a, b map[string]int) bool {

	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if _, ok := b[k]; !ok || v != b[k] {
			return false
		}
	}
	return true
}
