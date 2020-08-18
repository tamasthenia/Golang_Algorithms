package main

/*Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
Return the running sum of nums.

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
*/

func runningSum(nums []int) []int {

	/*res := make([]int, len(nums))
	res[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] + nums[i]
	}
	return res*/

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

/*func runningSum(nums []int) []int {

	var ansArray []int
	tempInt := 0

	for i := 0; i < len(nums); i++ {

		if i == 0 {

			ansArray = append(ansArray, nums[i])
			tempInt = nums[i]

		} else {

			tempInt = tempInt + nums[i]
			ansArray = append(ansArray, tempInt)

		}

	}

	return ansArray
}*/
