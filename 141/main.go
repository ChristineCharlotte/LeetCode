package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fastPt, slowPt := head, head
	for fastPt.Next != nil && fastPt.Next.Next != nil {
		fastPt = fastPt.Next.Next
		slowPt = slowPt.Next
		if fastPt == slowPt {
			return true
		}
	}
	return false
}
