package leetcode_go

import "strconv"

/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.

Example 1:

Input: num1 = "11", num2 = "123"
Output: "134"
Example 2:

Input: num1 = "456", num2 = "77"
Output: "533"
Example 3:

Input: num1 = "0", num2 = "0"
Output: "0"
*/
func addStrings(num1 string, num2 string) string {
	var add uint8 = 0
	var res string
	var i, j = len(num1) - 1, len(num2) - 1
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var n1, n2 uint8
		if i >= 0 {
			n1 = num1[i] - '0'
		}

		if j >= 0 {
			n2 = num2[j] - '0'
		}

		res = strconv.Itoa(int(n1+n2+add)%10) + res
		add = (n1 + n2 + add) / 10
	}

	return res
}
