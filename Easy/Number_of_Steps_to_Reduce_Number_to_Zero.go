package main

/* Given a non-negative integer num, return the number of steps to reduce it to zero.
If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.
*/

func numberOfSteps(num int) int {

	sum := num
	step := 0

	for sum > 0 {

		if sum%2 == 0 {
			sum = sum / 2
		} else {
			sum = sum - 1
		}

		step = step + 1

	}

	return step

}
