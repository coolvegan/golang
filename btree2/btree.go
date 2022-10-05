package btree2

type node struct {
	prev *node
	next *node
	val  int
	tn   *Treenode
}

type Treenode struct {
	Left  *Treenode
	Right *Treenode
	Val   int
}

type queue struct {
	head *node
	tail *node
	size int
}

func (q *queue) enque(n *node) {
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		c := q.tail
		c.next = n
		q.tail = n
		n.prev = c
	}
	q.size++
}

func (q *queue) deque() (*node, bool) {
	if q.size > 0 {
		r := q.head
		newHead := r.next
		q.head = newHead
		q.size--
		return r, true
	} else {
		var r node
		r.val = -1
		return &r, false
	}
}

func (q *queue) getSize() int {
	return q.size

}

func bfs(n *Treenode) []int {
	var queue []*Treenode
	var result []int
	queue = append(queue, n)
	for len(queue) > 0 {
		var currentNode Treenode
		currentNode = *queue[0]
		queue[0] = nil
		queue = queue[1:]
		result = append(result, currentNode.Val)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {

			queue = append(queue, currentNode.Right)
		}
	}
	return result
}

func levelOrder(root *Treenode) [][]int {
	var queue []*Treenode
	var result [][]int
	if root == nil {
		return result
	}
	queue = append(queue, root)
	var subresult []int
	queueLength := len(queue)
	for len(queue) > 0 {
		var currentNode Treenode
		currentNode = *queue[0]
		//Deque Start
		queue[0] = nil    //Release Memory on Position Zero
		queue = queue[1:] //Pop first Position
		//Deque End
		subresult = append(subresult, currentNode.Val)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {

			queue = append(queue, currentNode.Right)
		}
		if len(subresult) == queueLength {
			queueLength = len(queue)
			result = append(result, subresult)
			subresult = nil
		}
	}
	return result
}

func createTree(root []int) Treenode {
	var rootNode Treenode
	var queue []*Treenode
	rootNode.Val = root[0]
	queue = append(queue, &rootNode)
	for i := -1; i < len(root); i += 2 {
		currentNode := queue[0]
		queue[0] = nil    //Release Memory on Position Zero
		queue = queue[1:] //Pop first Position
		if i+2 < len(root) && root[i+2] != -1 {
			var node Treenode
			node.Val = root[i+2]
			currentNode.Left = &node
			queue = append(queue, &node)
		}
		if i+3 < len(root) && root[i+3] != -1 {
			var node Treenode
			node.Val = root[i+3]
			currentNode.Right = &node
			queue = append(queue, &node)
		}
	}
	return rootNode
}

func RightMaskedlevelOrder(root *Treenode) []int {
	var queue []*Treenode
	var result []int
	if root == nil {
		return result
	}
	queue = append(queue, root)
	subresultCount := 0
	queueLength := len(queue)
	for len(queue) > 0 {
		var currentNode Treenode
		currentNode = *queue[0]
		//Deque Start
		queue[0] = nil    //Release Memory on Position Zero
		queue = queue[1:] //Pop first Position
		//Deque End
		subresultCount += 1
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
		if subresultCount == queueLength {
			queueLength = len(queue)
			subresultCount = 0
			result = append(result, currentNode.Val)
		}
	}
	return result
}

func RightMaskedLevelOrderDfs(root *Treenode, depth int, result *[]int) {
  if root == nil {
    return
  }
  if depth > len(*result){
    *result = append(*result, root.Val)
  }
  RightMaskedLevelOrderDfs(root.Right, depth+1, result)
  RightMaskedLevelOrderDfs(root.Left, depth+1, result) 
}
