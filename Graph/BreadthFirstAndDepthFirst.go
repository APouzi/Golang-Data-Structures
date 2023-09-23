package graph

import "fmt"

func RunBFDF() {

}

type BFDFGraph struct {
	vertices []*BFDFNode
}

type BFDFNode struct {
	val int
	adj []*BFDFNode
}

func (g *BFDFGraph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = []*BFDFNode{}
	}
	new := BFDFNode{val: num}
	g.vertices = append(g.vertices, &new)
}

func (g *BFDFGraph) insertEdges(start int, end int) bool {
	var startNode *BFDFNode
	var endNode *BFDFNode

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
			startNode.adj = []*BFDFNode{}
		}
		if endNode.adj == nil {
			endNode.adj = []*BFDFNode{}
		}
		startNode.adj = append(startNode.adj, endNode)
		if startNode.val == 1 {
			fmt.Println(startNode.adj)
		}
		return true
	}
	return false
}
