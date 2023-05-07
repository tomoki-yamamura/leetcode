package link_list

// i couldnt solve
// func reorderList(head *ListNode)  {
// 	reversed := reverse(head)
// 	mergeTwoList(head, reversed)
// }

// func reverse(head *ListNode) (prev *ListNode) {
// 	for head != nil {
// 		tmp := head.Next
// 		head.Next = prev
// 		prev = head
// 		head = tmp
// 	}
// 	return
// }

// func mergeTwoList(head *ListNode, reversed *ListNode) {
// 	tmpHead := head
// 	tmpReversed := reversed
// 	var counter int
// 	for tmpHead != nil {
// 		counter++
// 		if counter % 2 == 0 {
// 			//偶数
// 			reversed
// 		} else {
// 			// 奇数
// 			tmp := head.Next
// 			head.Next = tmpReversed
// 			head = tmp
// 		}
// 	}
// }

func ReorderList(head *ListNode) {
	// find midpoint
	fast := head
	mid := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid = mid.Next
	}

	var prev *ListNode

	for mid != nil {
		temp := mid.Next
		mid.Next = prev
		prev = mid
		mid = temp
	}

	start := head


	for prev.Next != nil {
		tmp1, tmp2 := start.Next, prev.Next
		// 実態の操作
		start.Next = prev
		prev.Next = tmp1
		// pointerの操作
		start, prev = tmp1, tmp2
	}
}