package graph

import (
	"fmt"

	graphPull "github.com/APouzi/go-algos/Graph/Topological-Sort"
)

func Rungraphs() {
	// ---Topological Sort Array Version---
	graphReprestation := [][]int{{1, 0}, {2, 0}, {3, 2}, {2, 3}, {4, 2}, {5, 2}, {6, 3}, {4, 3}}
	graphPull.TopologicalSortArr(graphReprestation)

	// ---Topological Sort with Requirement Version--
	graphReprestation2 := [][]int{{1, 0}, {2, 0}, {3, 2}, {4, 2}, {5, 2}, {6, 3}, {4, 3}}
	canComplete := graphPull.TopologicalSortArrWithRequirement(graphReprestation2,6)
	fmt.Println("\n TopologicalSort Array version, can it complete?",canComplete)

	// ---Adjacency Graph Version of the Graph ---
	graph :=  graphPull.TopologicalGraph{}
	graph.InsertVerts(1)
	graph.InsertVerts(2)
	graph.InsertVerts(3)
	graph.InsertVerts(4)
	graph.InsertVerts(5)

	graph.InsertVerts(6)
	graph.InsertVerts(7)
	graph.InsertVerts(8)

	fmt.Println(graph.InsertEdges(1, 2))
	graph.InsertEdges(2, 3)
	graph.InsertEdges(3, 4)
	graph.InsertEdges(4, 5)

	graph.InsertEdges(1, 6)
	graph.InsertEdges(6, 7)
	graph.InsertEdges(7, 8)
	// graph.InsertEdges(8, 5)

	seen := make(map[*graphPull.TopologicalNode]int)
	list := &[][]int{}
	graphPull.TopologicalSort(graph.Vertices[0], seen, list, []int{})
	fmt.Println("Graph TopologicalSort List:",list)
	// prereq := [][]int{{0,1},{1,2},{2,3},{3,4}}
	// 0 > 1 > 2 > 3 > 4
	prereq2 := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {5, 6}}
	// 1 >
	// 		0
	// 2 >
	// prereq3Backwards := [][]int{{1,0},{2,1},{3,2},{4,3}}

	graphPull.TopologicalSort2(6, prereq2)
}