package main

type node struct {
	next  *node
	value int
}

func reverse_list(head *node) *node {
	current := head
	var prev *node

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

func reverse_sub_list(head *node, m int, n int) *node {
	var left_outside, middle_left, middle_right, right_outside, current *node
	if m == n {
		return head
	}
	current = head
	i := 1
	for current != nil && i < m-1 {
		current = current.next
		i++
	}

	*&left_outside = current
	current = current.next
	*&middle_left = current
	i++

	for current != nil && i < n {
		current = current.next
		i++
	}
	*&middle_right = current
	current = current.next
	i++
	*&right_outside = current

	if m != 1 {
		left_outside.next = nil
	}
	middle_right.next = nil

	if m == 1 {
		reverse_list(left_outside)
		left_outside.next = right_outside
		return middle_right
	}

	reverse_list(middle_left)
	left_outside.next = middle_right
	middle_left.next = right_outside
	return head
}
