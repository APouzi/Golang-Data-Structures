package main

import "fmt"

func main() {
	graph := Graph{}
	graph.insertVerts(1)
	graph.insertVerts(2)
	graph.insertVerts(3)
	graph.insertVerts(4)
	graph.insertVerts(5)

	graph.insertVerts(6)
	graph.insertVerts(7)
	graph.insertVerts(8)

	fmt.Println(graph.insertEdges(1, 2))
	graph.insertEdges(2, 3)
	graph.insertEdges(3, 4)
	graph.insertEdges(4, 5)

	graph.insertEdges(1, 6)
	graph.insertEdges(6, 7)
	graph.insertEdges(7, 8)
	// graph.insertEdges(8, 5)

	seen := make(map[*Node]int)
	list := &[][]int{}
	topologicalSort(graph.vertices[0], seen, list, []int{})
	fmt.Println(list)

}

type Graph struct {
	vertices []*Node
}

type Node struct {
	val int
	adj []*Node
}

func (g *Graph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = []*Node{}
	}
	new := Node{val: num}
	g.vertices = append(g.vertices, &new)
}

func (g *Graph) insertEdges(start int, end int) bool {
	var startNode *Node
	var endNode *Node

	for _, v := range g.vertices {
		if v.val == start {
			startNode = v
		}
		if v.val == end {
			endNode = v
		}
	}

	if startNode != nil && endNode != nil {
		if startNode.adj == nil {
			startNode.adj = []*Node{}
		}
		if endNode.adj == nil {
			endNode.adj = []*Node{}
		}
		startNode.adj = append(startNode.adj, endNode)
		if startNode.val == 1 {
			fmt.Println(startNode.adj)
		}
		return true
	}
	return false
}

func topologicalSort(node *Node, seen map[*Node]int, list *[][]int, temp []int) bool {
	if seen[node] == 2 {
		fmt.Println("hit 2")
		return false
	}

	if seen[node] == 1 {
		fmt.Println(node.val)
		return true
	}

	seen[node] = 2
	temp = append(temp, node.val)
	fmt.Println(node.adj)
	for _, v := range node.adj {
		if topologicalSort(v, seen, list, temp) == false {
			return false
		}
	}
	seen[node] = 1
	if len(node.adj) == 0 {
		*list = append(*list, temp)
	}

	return true
}
