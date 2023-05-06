package link_list

type ListNode struct {
    Val int
    Next *ListNode
}
 
func mergeKLists(lists []*ListNode) *ListNode {
	listsLength := len(lists)
	if listsLength == 0 {
		return nil
	}

	if listsLength == 1 {
		return lists[0]
	}

	prev := lists[0]
	for i := 1; i < listsLength; i++ {
		current := lists[i]
		prev = mergeTwoLists(prev, current)
	}
	return prev
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}