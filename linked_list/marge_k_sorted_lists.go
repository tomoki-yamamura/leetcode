package link_list

type ListNode struct {
    Val int
    Next *ListNode
}
 
// my first solutin: OK
// Beats 14.76%  Beats 53.2%

// func mergeKLists(lists []*ListNode) *ListNode {
// 	listsLength := len(lists)
// 	if listsLength == 0 {
// 		return nil
// 	}

// 	if listsLength == 1 {
// 		return lists[0]
// 	}

// 	prev := lists[0]
// 	for i := 1; i < listsLength; i++ {
// 		current := lists[i]
// 		prev = mergeTwoLists(prev, current)
// 	}
// 	return prev
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}
// 	if list1.Val < list2.Val {
// 		list1.Next = mergeTwoLists(list1.Next, list2)
// 		return list1
// 	}
// 	list2.Next = mergeTwoLists(list1, list2.Next)
// 	return list2
// }

// Best Solution
// Beats 100%  Beats 97.63%
// this is mergeSortLogic

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
			return nil
	}
	
	k := len(lists)
	if k == 1 {
			return lists[0]
	}
	
	l1 := mergeKLists(lists[:k/2])
	l2 := mergeKLists(lists[k/2:])
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	var dummyHead ListNode
	
	cur := &dummyHead
	for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
					cur.Next = l1
					l1 = l1.Next
			} else {
					cur.Next = l2
					l2 = l2.Next
			}
			cur = cur.Next
	}
	
	if l1 != nil {
			cur.Next = l1
	} else if l2 != nil {
			cur.Next = l2
	}
	
	return dummyHead.Next
}