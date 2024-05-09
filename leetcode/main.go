package main

import (
	"math"
	"sort"
	"strings"
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

/*
*最小覆盖子串
 */
func minWindow(s string, t string) string {

	var lens = len(s)
	var lent = len(t)

	var left = 0
	var ansLeft, ansRight = -1, lens

	cnt1 := make([]int, 128)
	cnt2 := make([]int, 128)

	for i := 0; i < lent; i++ {
		cnt1[t[i]]++
	}
	for i := 0; i < lens; i++ {
		c := s[i]
		cnt2[c]++

		for isMatch(cnt1, cnt2) {
			if i-left < ansRight-ansLeft {
				ansRight = i
				ansLeft = left
			}

			cnt2[s[left]]--
			left++
		}
	}

	if ansLeft < 0 {
		return ""
	} else {
		return s[ansLeft : ansRight+1]
	}
}

func isMatch(cnt1 []int, cnt2 []int) bool {

	for i := 'A'; i <= 'Z'; i++ {
		if cnt1[i] > cnt2[i] {
			return false
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		if cnt1[i] > cnt2[i] {
			return false
		}
	}
	return true
}

/*
*有效的数独
 */
func isValidSudoku(board [][]byte) bool {

	col := [9][10]int{}
	row := [9][10]int{}
	box := [9][10]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			var num = int8(board[i][j] - '0')
			if col[i][num] != 0 {
				return false
			}
			if row[j][num] != 0 {
				return false
			}

			if box[j/3+(i/3)*3][num] != 0 {
				return false
			}
			col[i][num] = 1
			row[j][num] = 1
			box[j/3+(i/3)*3][num] = 1

		}
	}
	return true
}

/*
*生命游戏
 */
func gameOfLife(board [][]int) {

	row := len(board)
	col := len(board[0])
	arr := make([][]int, row)
	for i := 0; i < row; i++ {
		arr[i] = make([]int, col)
	}

	xrr := [8]int{1, 1, 1, -1, -1, -1, 0, 0}
	yrr := [8]int{1, 0, -1, 1, 0, -1, 1, -1}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cnt := 0
			for k := 0; k < 8; k++ {
				curx := i + xrr[k]
				cury := j + yrr[k]

				if curx < 0 || curx >= row || cury < 0 || cury >= col {
					continue
				}
				if board[curx][cury] == 1 {
					cnt++
				}
			}

			if cnt < 2 || cnt > 3 || (cnt == 2 && board[i][j] == 0) {
				arr[i][j] = 0
			} else {
				arr[i][j] = 1
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = arr[i][j]
		}
	}
}

func wordPattern(pattern string, s string) bool {

	str := strings.Split(s, " ")
	if len(pattern) != len(str) {
		return false
	}

	hashMap1 := map[byte]string{}
	hashMap2 := map[string]byte{}

	for i := 0; i < len(str); i++ {
		if _, ok := hashMap1[pattern[i]]; !ok {
			hashMap1[pattern[i]] = str[i]
		} else {
			if hashMap1[pattern[i]] != str[i] {
				return false
			}
		}

		if _, ok := hashMap2[str[i]]; !ok {
			hashMap2[str[i]] = pattern[i]
		} else {
			if hashMap2[str[i]] != pattern[i] {
				return false
			}
		}
	}
	return true
}

/*
*有效的字母异位词
 */
func isAnagram(s string, t string) bool {

	arr := [26]int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		arr[t[i]-'a']--
		if arr[t[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

/*
*字母异位词分组
 */
func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[string][]string)
	for _, str := range strs {
		t := []byte(str)
		sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })

		temp := string(t)
		hashMap[temp] = append(hashMap[temp], str)
	}

	ans := make([][]string, 0)

	for _, v := range hashMap {
		ans = append(ans, v)
	}
	return ans
}
