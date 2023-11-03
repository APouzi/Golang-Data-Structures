package graph

// Sometimes we don't have nodes that are representing multivariable/multidimensional data structures. This means that we only have values like numbers, letters, symbols or what have you
type GraphAdj struct {
	Vertices map[int][]int
}

// This one is usually best because of O(1) loop up time plus ability to loop over them as we are doing out path searches. 
type GraphNodeAdj struct {
	Vertices map[*GraphNode][]*Node
}

// Here we have the matrix that is representing a bidirectional graphs, which are technically acyclic. Which means that finding the solution to a path would mean that it's going to give us an extremely large number of possible paths
type GraphArr struct {
	Matrix [][]int
}

// This is where we have a list of Graph nodes and the edges that we are being represented. We loop over the vertices and insert those looped over nodes into the edges.
type NodeGraph struct {
	Vertices []*GraphNode
	Edges    map[*GraphNode]*Edges
}



// Utility Componenets
type Edges struct {
	To   *GraphNode
	From *GraphNode
}

type GraphNode struct {
	Val int
}