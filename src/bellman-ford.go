package main

import (
	"fmt"
	"math"
)

func main() {
	graph := Graph{}
	graph.insertVerts(3)
	graph.insertVerts(4)
	graph.insertVerts(5)
	graph.insertVerts(6)
	graph.insertVerts(1)
	graph.insertVerts(2)

	graph.insertEdges(1, 2, 1)
	graph.insertEdges(1, 4, 0)
	graph.insertEdges(2, 3, 1)
	graph.insertEdges(3, 4, 2)
	graph.insertEdges(4, 5, 3)
	graph.insertEdges(5, 6, 2)
	fmt.Println(graph.verticesInsert)
	graph.bellmanFord(1)
	// fmt.Print(graph.vertices[0].adj)
}

// Seems like you need the edge list first and foremost.
func (g *Graph) bellmanFord(start int) {

	distance := make(map[int]int)
	for _, val := range g.vertices {
		distance[val.val] = math.MaxUint32
	}
	start = 1
	distance[start] = 0
	for range g.vertices {
		for _, edge := range g.edges {
			if distance[edge.from]+edge.weight < distance[edge.to] {
				distance[edge.to] = distance[edge.from] + edge.weight
			}
		}
	}

	fmt.Println(distance)
}

// Graph functions
type Graph struct {
	vertices map[int]*Node
	// We have this here because we need to have the same
	verticesInsert []int
	edges          []Edge
}

type Edge struct {
	from   int
	to     int
	weight int
	toNode *Node
}

type Node struct {
	val int
	adj []Edge
}

func (g *Graph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = make(map[int]*Node)
	}
	new := Node{val: num}
	g.vertices[num] = &new
	g.verticesInsert = append(g.verticesInsert, num)

}

func (g *Graph) insertEdges(start int, end int, weight int) bool {
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
			startNode.adj = []Edge{}
		}
		if endNode.adj == nil {
			endNode.adj = []Edge{}
		}
		edge := Edge{from: start, to: end, toNode: endNode, weight: weight}
		startNode.adj = append(startNode.adj, edge)
		g.edges = append(g.edges, edge)
		return true
	}
	return false
}
