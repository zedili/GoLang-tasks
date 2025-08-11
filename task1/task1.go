package main

import (
	"fmt"
	"sort"
	"strings"
)

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
// 方法1
func singleNumber1(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		v, exist := numMap[num]
		if exist {
			numMap[num] = v + 1
		} else {
			numMap[num] = 1
		}
	}
	for k, v := range numMap {
		// fmt.Println(k, v)
		if v == 1 {
			return k
		}
	}
	return 0
}

// 方法2
func singleNumber2(nums []int) int {
	result := 0
	for _, num := range nums {
		fmt.Println(result)
		result ^= num
	}
	return result
}

// 是否回文数
func isPalindrome(x int) bool {

	if x == 0 {
		return true
	}

	if x < 0 {
		return false
	}

	if x%10 == 0 {
		return false
	}

	revseredX := 0
	for x > revseredX {
		// x     revseredX
		// 123   3
		// 12    3 *10 + 2 = 32
		// 121   1
		// 12    1 * 10 + 2  = 12
		// 1     12 * 10 + 1 = 121
		// 1221

		revseredX = revseredX*10 + x%10
		x /= 10

	}
	return x == revseredX || x == revseredX/10
}

// 有效的括号
// 考察：字符串处理、栈的使用
// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	stack := []rune{} // rune 数组实现一个栈
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 入栈
			stack = append(stack, char)
		case ')', '}', ']':
			// 如果空栈，遇到闭括号，说明没有开括号，无效
			if len(stack) == 0 {
				return false
			}
			//  从栈顶去取出开括号，取出当前闭括号对应的开括号，判断是否相同
			if stack[len(stack)-1] != pairs[char] {
				return false
			}
			// 括号匹配成功，弹出栈顶符号
			stack = stack[:len(stack)-1]
		}
	}
	// 所有括号匹配成功，栈为空
	return len(stack) == 0
}

// 最长公共前缀
// 查找字符串数组中的最长公共前缀
// 如果不存在公共前缀，返回空字符串
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 初始化前缀为第一个字符串
	prefix := strs[0]
	// 遍历数组中的所有字符串
	for i := 1; i < len(strs); i++ {
		// 检查当前字符串是是否以 prefix 开头（如果是prefix 开头，index 应该 = 0）
		for strings.Index(strs[i], prefix) != 0 {
			// 如果 prefix 为空，返回空字符串
			if len(prefix) == 0 {
				return ""
			}
			// 如果不 prefix 开头， 逐个字符缩短前缀（移除最后一个字符）
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	n := len(digits)
	// 从数组的最后一位开始遍历
	for i := n - 1; i >= 0; i-- {
		// 如果数字小于 9 ，当前位加 1 返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 当前位是 9 ，现将当前位设置为 0 ，调到上一位处理
		digits[i] = 0
	}
	// 如果所有为都 9 ，才会执行到这里，所有位都为 9 时，加一后，所有位都设置为 0 ，再在前面加一个 1 返回
	return append([]int{1}, digits...)
}

// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
// 当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
func removeDuplicates(nums []int) int {
	// 条件：nums 是有序数组
	if len(nums) == 0 {
		return 0
	}
	j := 1 // 记录不重复元素的索引
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] { // 逐个元素判断当前元素是否与后一个元素相等
			nums[j] = nums[i+1] // 如果当前元素与后一个元素不相等，改写当前数组，保存不重复元素（删除重复元素）
			j++
		}
	}
	// fmt.Println(nums)
	return j
}

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	// 根据 starti ，重新排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 合并区间结果集，包含第一个区间
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// fmt.Println("last:", last)
		// fmt.Println("current:", current)

		// 检查是否重叠区间
		if current[0] <= last[1] {
			// 区间有重叠：当前区间起点 <= 最后合并区间的终点
			// 合片重叠区间：更新最后 合并区间的终点
			last[1] = max(last[1], current[1])
		} else {
			// 区间没有重叠
			// 增加一个合并区间
			merged = append(merged, current)
		}
	}

	return merged
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
func twoSum(nums []int, target int) []int {
	// 定义一个 map 存在数字在数组中的索引
	indexMap := make(map[int]int)
	for index, num := range nums {
		// 遍历数组，根据目标值和当前数字的差，在 map 找差的索引
		otherAddtion := target - num
		otherIndex, exits := indexMap[otherAddtion]
		if exits {
			// 如果差的索引存在，返回当前数字 和 差的索引
			return []int{index, otherIndex}
		}
		// 如果不存在，将当前设置的索引保存在 map
		indexMap[num] = index
	}
	return nil
}

func main() {
	// fmt.Println(singleNumber1([]int{1, 1, 1, 2, 2, 3}))
	// fmt.Println(singleNumber2([]int{1, 1, 1, 2, 2, 3}))
	// fmt.Println(isPalindrome(121))
	// fmt.Println(isValid("([123456])"))
	// fmt.Println(longestCommonPrefix([]string{"flow", "flower"}))
	// fmt.Println(plusOne([]int{9, 9, 9}))
	// fmt.Println(removeDuplicates([]int{1, 2, 3}))
	// fmt.Println(merge([][]int{{2, 3}, {1, 6}, {8, 10}, {15, 18}}))
	fmt.Println(twoSum([]int{2, 7, 3, 6, 9}, 9))
}
