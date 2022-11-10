package dijkstra

import (
	"fmt"
	"math"
)

func sub(k int, adj *[][]Node, result *[]int, seen *map[int]bool) int {
	for i := 0; i < len((*adj)[k]); i++ {
		node := (*adj)[k][i]
		cost := node.cost
		dest := node.destination
		if (*result)[dest] > cost+(*result)[k] {
			(*result)[dest] = cost + (*result)[k]
		}
	}
	(*seen)[k] = true
	min := int(math.Inf(0))
	var node int
	node = -1
	for i := 0; i < len(*result); i++ {
		if i != k && !(*seen)[i] {
			if (*result)[i] < int(min) {
				min = (*result)[i]
				node = i
			}
		}
	}
	return node
}

func dik(k, nodes int, graph [][]int) int {
	seen := make(map[int]bool)
	adj := getAdj(nodes, graph)
	result := make([]int, nodes)
	//val auf unendlich
	for i := 0; i < len(result); i++ {
		result[i] = int(math.Inf(0))
	}
	//besuche jeden erreichbaren knoten
	result[k] = 0
	var node = k
	for node != -1 {
		node = sub(node, &adj, &result, &seen)
	}
	res := -1
	for i := 0; i < len(result); i++ {
		if result[i] > res {
			res = result[i]
		}
	}
	fmt.Println(result)
	return res
}

type Node struct {
	source      int
	destination int
	cost        int
}

func getAdj(nodes int, graph [][]int) [][]Node {
	result := make([][]Node, nodes)
	for i := 0; i < len(graph); i++ {
		quelle := graph[i][0]
		ziel := graph[i][1]
		dauer := graph[i][2]
		var node Node
		node.cost = dauer
		node.source = quelle
		node.destination = ziel
		result[quelle] = append(result[quelle], node)
	}
	return result
}
