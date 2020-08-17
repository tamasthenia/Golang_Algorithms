package main

/*
Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

Return the array in the form [x1,y1,x2,y2,...,xn,yn].

Example 1:

Input: nums = [2,5,1,3,4,7], n = 3
Output: [2,3,5,4,1,7]
Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].
*/

/*func shuffleTheArray(nums []int, n int) []int {

    arr1 := nums[0:n]
    arr2 := nums[n:]

    ansArray := make([]int, len(nums))

    index1 := 0
    index2 := 1

    for i := 0; i < len(arr1); i++ {
        ansArray[index1] = arr1[i]
        index1 = index1 + 2
	}

    for i := 0; i < len(arr2); i++ {
        ansArray[index2] = arr2[i]
        index2 = index2 + 2
	}

    return ansArray
}*/

func shuffle(nums []int, n int) []int {
	result := make([]int, 0, n*2)
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[i+n])
	}
	return result
}
