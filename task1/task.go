package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	singleNumberTest()
	fmt.Println("121是否为回文数:", isPalindrome(121))
	fmt.Println("最长公共前缀", longestCommonPrefix())
	fmt.Println("给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一: ", plusOne())
	delSlice()
}

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
*/
func singleNumberTest() {
	nums := []int{1, 2, 3, 1, 3}
	println(singleNumber(nums))

}

func singleNumber(nums []int) int {

	times := make(map[int]int)
	for _, v := range nums {
		i := times[v]
		times[v] = i + 1

	}
	for k, v := range times {

		if v == 1 {
			return k
		}
	}

	return -1
}

/*
判断一个整数是否是回文数
*/
func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

/*
编写一个函数来查找字符串数组中的最长公共前缀
*/
func longestCommonPrefix() string {
	strs := [3]string{"flower", "flow", "flight1"}

	var maxLen = 0
	var maxIndex = 0
	for i, v := range strs {
		fmt.Println("i=", i, "v=", v)
		if len(v) > maxLen {
			maxLen = len(v)
			maxIndex = i
		}
	}
	return strs[maxIndex]
}

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组
*/
func plusOne() int {
	digits := []int{1, 2, 3, 4}
	total := 0
	for i, v := range digits {
		fmt.Println("i=", i, "v=", v)
		ww := math.Pow(10, float64(len(digits)-i-1))
		total = total + v*int(ww)
	}
	return total + 1
}

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func delSlice() {
	nums := []int{1, 2, 3, 4, 5, 1, 3, 6} // 输入数组
	total := 4
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == total {
				mp[nums[i]] = nums[j]
			}
		}
	}
	fmt.Println("找出的数组放入map为: ", mp)

}
