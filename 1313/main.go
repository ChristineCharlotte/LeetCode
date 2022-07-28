package main

func decompressRLElist(nums []int) []int {
	var re []int
	for i := 0; i < len(nums); i += 2 {
		for count := 0; count < nums[i]; count++ {
			re = append(re, nums[i+1])
		}
	}
	return re
}
