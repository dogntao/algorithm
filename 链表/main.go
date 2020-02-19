package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * 1、链表头部 pre
 * 2、已处理尾部 tail
 * 2、已翻转局部头部 ph和尾部 pt
 * 3、即将翻转头部 cur
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var pre *ListNode = nil
	var tail *ListNode = nil
	var ph *ListNode = nil
	var pt *ListNode = nil
	cur := head
	for cur != nil {
		j := 0
		// 构造局部翻转链表
		var listNodeTmp *ListNode = nil
		for i := 0; i < k; i++ {
			j++
			listNodeTmp = cur
			cur = cur.Next
		}
		// k个一组翻转链表
		if j%k == 0 {
			ph = reverseList(listNodeTmp)
			pt = listNodeTmp
			if tail != nil {
				tail.Next = ph
			}
			tail = pt
		} else {
			// 不足k个不翻转
			ph = listNodeTmp
			tail.Next = ph
		}
		// 初始化链表头
		if pre == nil {
			pre = ph
		}
	}
	return pre
}

/*
 * 翻转链表
 * pre链表头部，cur当前翻转位置
 * cur.Next=pre,pre=cur,cur=cur.Next
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
