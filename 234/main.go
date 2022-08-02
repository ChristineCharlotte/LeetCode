package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	var list []int
	for node := head; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	fmt.Println(list)

	for left, right := 0, len(list)-1; left < len(list)/2; {
		if list[left] != list[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
