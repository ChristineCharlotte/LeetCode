// 迭代
package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
   //dummyHead := new(ListNode)
   dummyHead := &ListNode{}
   cursor := dummyHead
   for list1 != nil && list2 != nil {
      if list1.Val < list2.Val {
         cursor.Next = list1
         list1 = list1.Next
      } else {
         cursor.Next = list2
         list2 = list2.Next
      }
      cursor = cursor.Next
   }
   if list1 == nil {
      cursor.Next = list2
   } else {
      cursor.Next = list1
   }
   return dummyHead.Next
}