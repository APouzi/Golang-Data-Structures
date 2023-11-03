package graph

type GraphAdj struct {
	Vertices map[int][]int
}

type GraphNodeAdj struct {
	Vertices map[*GraphNode][]*Node
}

type GraphArr struct {
	Matrix [][]int
}

type GraphNodes struct {
	Vertices []*GraphNode
	Edges    []*Edges
}

// Utility Componenets
type Edges struct {
	To   *GraphNode
	From *GraphNode
}

type GraphNode struct {
	Val int
}