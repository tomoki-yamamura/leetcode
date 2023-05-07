package main

import (
	"fmt"

	"github.com/tomoki-yamamura/go-docker/linked_list"
)

func main() {
	// arr := &link_list.ListNode{
	// 	Val: 1,
	// 	Next: &link_list.ListNode{
	// 		Val: 2,
	// 		Next: &link_list.ListNode{
	// 			Val: 3,
	// 			Next: nil,
	// 		},
	// 	},
	// }
	arr := &link_list.ListNode{
		Val: 1,
		Next: &link_list.ListNode{
			Val: 2,
			Next: &link_list.ListNode{
				Val: 3,
				Next: &link_list.ListNode{
					Val: 4,
					Next: &link_list.ListNode{
						Val: 5,
						Next: &link_list.ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	link_list.ReorderList(arr)
	fmt.Println(arr)
}
