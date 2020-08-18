package main

/*
Balanced strings are those who have equal quantity of 'L' and 'R' characters.

Given a balanced string s split it in the maximum amount of balanced strings.

Return the maximum amount of splitted balanced strings.



Example 1:

Input: s = "RLRRLLRLRL"
Output: 4
Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.
*/

func balancedStringSplit(s string) int {

	var res, count int = 0, 0
	for _, val := range s {
		if val == 'L' {
			count++
		} else if val == 'R' {
			count--
		}

		if count == 0 {
			res++
		}
	}
	return res

}
