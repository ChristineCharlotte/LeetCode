package main

import "fmt"

func main() {
	s := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(s))
}

func findDisappearedNumbers(nums []int) []int {
	Len := len(nums)
	var re []int
	for _, num := range nums {
		num = (num - 1) % Len
		nums[num] += Len
	}
	for index, num := range nums {
		if num <= Len {
			re = append(re, index+1)
		}
	}
	return re
}
