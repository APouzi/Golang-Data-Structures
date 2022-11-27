package main

import "fmt"

func main() {

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
