package main

func main() {

}

func moveZeroes(nums []int) {
	i := 0 // 遍历所在的位置
	j := 0 // 非零元素所在的位置
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for k := j + 1; k < len(nums); k++ {
		nums[k] = 0
	}
}
