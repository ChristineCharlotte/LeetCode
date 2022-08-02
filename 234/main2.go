package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 如果链表是奇数个节点，把正中间的归到左边
	if fast != nil {
		slow = slow.Next
	}

	slow = reverse(slow)
	fast = head
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for next := head.Next; head != nil; {
		head.Next = prev
		prev = head
		head = next
	}
	return head
}
