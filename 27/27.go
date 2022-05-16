package leetcode

func removeElement(nums []int, val int) int {
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] == val && i != len(nums) - 1 {
            nums = append(nums[:i], nums[i+1:]...)
        } else if nums[i] == val {
            nums = nums[:i]
        }
    }
    return len(nums)
}