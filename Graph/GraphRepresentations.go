package graph

// Sometimes we don't have nodes that are representing multivariable/multidimensional data structures. This means that we only have values like numbers, letters, symbols or what have you
type GraphAdj struct {
	Vertices map[int][]int
}

// This one is usually best because of O(1) loop up time plus ability to loop over them as we are doing out path searches. 
type GraphNodeAdj struct {
	Vertices map[*GraphNode][]*Node
}

// Here we have the matrix that is representing a bidirectional graphs, which are technically acyclic. Which means that finding the solution to a path would mean that it's going to give us an extremely large number of possible paths. 
// Note: There is two versions of this, one is where the graph is representing the nodes themselves. An example is like this where we have the 1's representing the nodes and fact they are next to one another means they are together.
// [1,1,1,0]
// [1,1,1,0]
// [1,0,0,0]

// The other one is where we have the x,y is going to be representing the nodes and the integers inside are the edges. Meaning that if in the example below (1 > 0) and (0 > 1) are connected to each other. 
// [0,1,1]
// [1,0,1]
// [1,0,0]
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