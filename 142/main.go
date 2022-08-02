package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var hasCycle bool
	fastPt, slowPt := head, head
	for fastPt.Next != nil && fastPt.Next.Next != nil {
		fastPt = fastPt.Next.Next
		slowPt = slowPt.Next
		if fastPt == slowPt {
			hasCycle = true
		}
	}
	hasCycle = false

	if hasCycle {
		slowPt = head
		for slowPt != fastPt {
			slowPt = slowPt.Next
			fastPt = fastPt.Next
		}
		return slowPt
	}
	return nil
}
