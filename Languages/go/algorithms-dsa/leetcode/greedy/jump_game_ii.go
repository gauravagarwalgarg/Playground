package main

import "fmt"

/*
  LC 45 - Jump Game II
  Topic: Greedy
  Difficulty: Medium
  Time: O(n) | Space: O(1)
*/

func jump(nums []int) int {
	jumps, curEnd, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == curEnd {
			jumps++
			curEnd = farthest
		}
	}
	return jumps
}

func main() {
	if jump([]int{2, 3, 1, 1, 4}) == 2 {
		fmt.Println("PASS: jump [2,3,1,1,4]")
	} else {
		fmt.Println("FAIL: jump [2,3,1,1,4]")
	}

	if jump([]int{2, 3, 0, 1, 4}) == 2 {
		fmt.Println("PASS: jump [2,3,0,1,4]")
	} else {
		fmt.Println("FAIL: jump [2,3,0,1,4]")
	}
}
