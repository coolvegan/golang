package countProblem

import (
	"math"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	result := base
	for i := 1; i < exp; i++ {
		result *= base
	}
	return result
}

func getHeight(root *Node) int {
	if root == nil {
		return 0
	}
	currentNode := root

	count := 0
	for currentNode != nil {
		count++
		currentNode = currentNode.left
	}
	return count - 1
}

func nodeExists(idx, height int, node *Node) bool {
	left := 0
	right := pow(2, height) - 1
	count := 0
	for count < height {
		midOfNode := int(math.Ceil(float64(left+right) / 2))
		if idx >= midOfNode {
			node = node.right
			left = midOfNode
		} else {
			node = node.left
			right = midOfNode - 1
		}
		count++
	}
	return node != nil
}

func countNodes(root *Node) int {
	if root == nil {
		return 0
	}
	height := getHeight(root)
	if height == 0 {
		return 1
	}
	upperCount := pow(2, height) - 1

	left := 0
	right := upperCount

	for left < right {
		idxTofind := int(math.Ceil(float64(left+right) / 2))
		if nodeExists(idxTofind, height, root) {
			left = idxTofind
		} else {
			right = idxTofind - 1
		}
	}
	return upperCount + left + 1
}

func isValidBST(root *Node) bool {
	if root == nil {
		return true
	}

	return dfs(root, int(math.Inf(-1)), int(math.Inf(+1)))
}

func dfs(root *Node, min, max int) bool {
	if root.val <= min || root.val >= max {
		return false
	}
	if root.left != nil {
		if !dfs(root.left, min, root.val) {
			return false
		}
	}
	if root.right != nil {
		if !dfs(root.right, root.val, max) {
			return false
		}
	}
	return true
}
