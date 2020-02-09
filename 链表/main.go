package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

// 翻转链表
func flip(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		cur, cur.next, pre = cur.next, pre, cur
	}
	return pre
}

func main() {
	head := &ListNode{val: 0}
	lisNode1 := &ListNode{val: 1}
	lisNode2 := &ListNode{val: 2}
	lisNode3 := &ListNode{val: 3}
	lisNode4 := &ListNode{val: 4}
	lisNode5 := &ListNode{val: 5}

	head.next = lisNode1
	lisNode1.next = lisNode2
	lisNode2.next = lisNode3
	lisNode3.next = lisNode4
	lisNode4.next = lisNode5

	fmt.Println(head)
	fmt.Println(head.next)
	fmt.Println(head.next.next)
	fmt.Println(head.next.next.next)
	fmt.Println(head.next.next.next.next)
	fmt.Println(head.next.next.next.next.next)

	pre := flip(head)
	fmt.Println(pre)
	fmt.Println(pre.next)
	fmt.Println(pre.next.next)
	fmt.Println(pre.next.next.next)
	fmt.Println(pre.next.next.next.next)
	fmt.Println(pre.next.next.next.next.next)
}
