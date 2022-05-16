package main

func removeDuplicates(nums []int) int {
    for k := len(nums)-1; k >= 1; k-- {
        if nums[k-1] == nums[k] {
            nums = append(nums[:k], nums[k+1:]...)
        }
    }
    return len(nums)
}