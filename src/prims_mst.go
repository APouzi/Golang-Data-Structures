package main

func main() {

}

// Prims

func () prims_algorithim() {}

func () add_edges() {}

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

// remember MST's are Undirected, must be undirected edges.
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
		endNode.adj = append(endNode.adj)
		g.edges = append(g.edges, edge)
		return true
	}
	return false
}
