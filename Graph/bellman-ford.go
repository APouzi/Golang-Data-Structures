package graph

import (
	"fmt"
	"math"
)

func run() {
	graph := BellmanFordGraph{}
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
func (g *BellmanFordGraph) bellmanFord(start int) {

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
type BellmanFordGraph struct {
	vertices map[int]*BMNode
	// We have this here because we need to have the same
	verticesInsert []int
	edges          []BMEdge
}

type BMEdge struct {
	from   int
	to     int
	weight int
	toNode *BMNode
}

type BMNode struct {
	val int
	adj []BMEdge
}

func (g *BellmanFordGraph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = make(map[int]*BMNode)
	}
	new := BMNode{val: num}
	g.vertices[num] = &new
	g.verticesInsert = append(g.verticesInsert, num)

}

func (g *BellmanFordGraph) insertEdges(start int, end int, weight int) bool {
	var startNode *BMNode
	var endNode *BMNode

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
			startNode.adj = []BMEdge{}
		}
		if endNode.adj == nil {
			endNode.adj = []BMEdge{}
		}
		edge := BMEdge{from: start, to: end, toNode: endNode, weight: weight}
		startNode.adj = append(startNode.adj, edge)
		g.edges = append(g.edges, edge)
		return true
	}
	return false
}
