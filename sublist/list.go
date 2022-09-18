package main

type node struct {
	head  *node
	tail  *node
	child *node
	value int
}

func handle_child(parent *node) *node {
	if parent.child == nil {
		return parent
	}
	current := parent
	right_child := current.child
	var right_neighbour *node
	if current.tail != nil {
		*&right_neighbour = current.tail
	}
	for right_child.tail != nil {
		handle_child(right_child)
		right_child = right_child.tail
	}

	if current.tail != nil {
		right_neighbour.head = right_child
		right_child.tail = right_neighbour
	}
	current.tail = current.child
	current.child = nil
	return parent
}

func flatten_list(head *node) *node {
	current := head
	for current != nil {
		handle_child(current)
		current = current.tail
	}
	return head
}

func findCycle(head *node) bool {
	set := make(map[*node]struct{})
	var node_exists_in_set = struct{}{}
	current := head
	for true {
		if _, ok := set[current]; ok {
			return true
		}
		if current == nil {
			break
		}
		set[current] = node_exists_in_set
		current = current.tail
	}
	return false
}

func findCycleOptimized(head *node) bool {
	current := head
	if current.tail == nil {
		return false
	}
	hare := current
	tortoise := current
	for current != nil {
		tortoise = tortoise.tail
		hare = hare.tail.tail
		if hare == tortoise {
			return true
		}
		if hare == nil {
			return false
		}
	}
	return false
}
func findCycleStartingNode(head *node) *node {
	current := head
	if current == nil || current.tail == nil {
		return nil
	}
	hare := current
	tortoise := current
	for current != nil {
		tortoise = tortoise.tail
		hare = hare.tail.tail
		if hare == tortoise {
			break
		}
		if hare == nil {
			return nil
		}
	}
	p1 := head
	p2 := tortoise
	for p1 != p2 {
		p1 = p1.tail
		p2 = p2.tail
	}
	return p2
}
