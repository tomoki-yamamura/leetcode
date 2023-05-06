package link_list

// my Solutin
// Beats 100% Memory 2.3 MB Beats 48.67%
// func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
// 	reversed := reverseList(head)
// 	result := removeNthFromFirst(reversed, n)
// 	return reverseList(result)
// }

// func removeNthFromFirst(head *ListNode, n int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	if n == 1 {
// 		return head.Next
// 	}
// 	// 新しいリストのヘッドを作成
// 	result := &ListNode{}
// 	// 新しいリストのカレントポインタを作成
// 	newCur := result
// 	counter := 0
// 	for head != nil {
// 		counter++
// 		if counter == n {
// 			head = head.Next
// 			continue
// 		}
// 		newNode := &ListNode{
// 			Val: head.Val,
// 		}
// 		newCur.Next = newNode
// 		newCur = newCur.Next
// 		head = head.Next
// 	}
// 	return result.Next
// }

// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	for head != nil {
// 		tmp := head.Next
// 		head.Next = prev
// 		prev = head
// 		head = tmp
// 	}
// 	return prev
// }

// my Solutin2
// Beats 100% Memory 2.3 MB Beats 48.67%
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	reversed := reverseList(head)
// 	result := removeNthFromFirst(reversed, n)
// 	return reverseList(result)
// }

// func removeNthFromFirst(head *ListNode, n int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	if n == 1 {
// 		return head.Next
// 	}
// 	// 新しいリストのヘッドを作成
// 	result := &ListNode{}
// 	// 新しいリストのカレントポインタを作成
// 	newCur := result
// 	counter := 0
// 	for head != nil {
// 		counter++
// 		if counter == n {
// 			head = head.Next
// 			continue
// 		}
// 		newNode := &ListNode{
// 			Val: head.Val,
// 		}
// 		newCur.Next = newNode
// 		newCur = newCur.Next
// 		head = head.Next
// 	}
// 	return result.Next
// }

// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	for head != nil {
// 		tmp := head.Next
// 		head.Next = prev
// 		prev = head
// 		head = tmp
// 	}
// 	return prev
// }

