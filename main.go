// package main

// import (
// 	"fmt"

// 	"github.com/tomoki-yamamura/go-docker/linked_list"
// )

// func main() {
// 	// arr := &link_list.ListNode{
// 	// 	Val: 1,
// 	// 	Next: &link_list.ListNode{
// 	// 		Val: 2,
// 	// 		Next: &link_list.ListNode{
// 	// 			Val: 3,
// 	// 			Next: nil,
// 	// 		},
// 	// 	},
// 	// }
// 	arr := &link_list.ListNode{
// 		Val: 1,
// 		Next: &link_list.ListNode{
// 			Val: 2,
// 			Next: &link_list.ListNode{
// 				Val: 3,
// 				Next: &link_list.ListNode{
// 					Val: 4,
// 					Next: nil,
// 				},
// 			},
// 		},
// 	}

// 	link_list.ReorderList(arr)
// 	fmt.Println(arr)
// }

package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	p := &Person{Name: "Bob", Age: 30}
	fmt.Printf("before: %p\n", p) // &{Bob 30}
	updatePersonPointer(p)
	fmt.Printf("after: %p\n", p) // &{Bob 30}

	fmt.Println(p)
}

func updatePersonPointer(p *Person) {
	fmt.Printf("arg P: %p\n", p)
	newP := &Person{Name: "Charlie", Age: 35}
	fmt.Printf("newP: %p\n", newP)
	p = newP
	fmt.Printf("p: %p, newP: %p\n", p, newP)
}
