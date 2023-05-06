package main

import (
	"github.com/tomoki-yamamura/go-docker/linked_list"
)

func main() {
	arr := &link_list.ListNode{
		Val: 1,
		Next: &link_list.ListNode{
			Val: 2,
			Next: &link_list.ListNode{
				Val: 3,
				Next: &link_list.ListNode{
					Val: 4,
					Next: nil,
				},
			},
		},
	}

	link_list.RemoveNthFromEnd(arr, 2)
}