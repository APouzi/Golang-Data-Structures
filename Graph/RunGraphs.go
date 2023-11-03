package graph

import "fmt"

func Rungraphs() {
	// ---Topological Sort Array Version---
	graphReprestation := [][]int{{1, 0}, {2, 0}, {3, 2}, {2, 3}, {4, 2}, {5, 2}, {6, 3}, {4, 3}}
	TopologicalSortArr(graphReprestation)

	// ---Adjacency Graph Version of the Graph ---
	graph := TopologicalGraph{}
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

	seen := make(map[*TopologicalNode]int)
	list := &[][]int{}
	topologicalSort(graph.Vertices[0], seen, list, []int{})
	fmt.Println(list)
	// prereq := [][]int{{0,1},{1,2},{2,3},{3,4}}
	// 0 > 1 > 2 > 3 > 4
	prereq2 := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {5, 6}}
	// 1 >
	// 		0
	// 2 >
	// prereq3Backwards := [][]int{{1,0},{2,1},{3,2},{4,3}}

	topologicalSort2(6, prereq2)
}