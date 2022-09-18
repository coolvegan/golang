package main

type Node struct {
	Next  *Node
	Value int
}

func reverse(head *Node) *Node {
	var current *Node
	current = head
	var prev *Node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
