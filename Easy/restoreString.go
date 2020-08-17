package main

import "fmt"

/*Given a string s and an integer array indices of the same length.

The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

Return the shuffled string.

Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
Output: "leetcode"
Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.
*/

func restoreString(s string, indices []int) string {

	//Rune literals are just 32-bit integer values

	ans := make([]byte, len(s))
	for i, index := range indices {
		fmt.Println(index)
		ans[index] = s[i]

	}
	return string(ans)

}

/*func restoreString(s string, indices []int) string {

    myMap := make(map[int]string)

	for i, v := range s {
		myMap[indices[i]] = string(v)
s	}

    ans := ""

    for i := 0; i < len(myMap); i++ {
		ans = ans + myMap[i]
	}

    return ans
}*/
