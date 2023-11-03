package graph

import (
	"fmt"
)

func runTopological() {
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
	prereq2 := [][]int{{0,1},{1,2},{1,3},{3,4},{5,6}}
	// 1 >
	// 		0 
	// 2 >
	// prereq3Backwards := [][]int{{1,0},{2,1},{3,2},{4,3}}
	
	topologicalSort2(6, prereq2)

}

type TopologicalGraph struct {
	Vertices []*TopologicalNode
}

type TopologicalNode struct {
	Val int
	Adj []*TopologicalNode
}

func (g *TopologicalGraph) insertVerts(num int) {
	if g.Vertices == nil {
		g.Vertices = []*TopologicalNode{}
	}
	new := TopologicalNode{Val: num}
	g.Vertices = append(g.Vertices, &new)
}

func (g *TopologicalGraph) insertEdges(start int, end int) bool {
	var startNode *TopologicalNode
	var endNode *TopologicalNode

	for _, v := range g.Vertices {
		if v.Val == start {
			startNode = v
		}
		if v.Val == end {
			endNode = v
		}
	}

	if startNode != nil && endNode != nil {
		if startNode.Adj == nil {
			startNode.Adj = []*TopologicalNode{}
		}
		if endNode.Adj == nil {
			endNode.Adj = []*TopologicalNode{}
		}
		startNode.Adj = append(startNode.Adj, endNode)
		if startNode.Val == 1 {
		}
		return true
	}
	return false
}

func topologicalSort(node *TopologicalNode, seen map[*TopologicalNode]int, list *[][]int, temp []int) bool {
	if seen[node] == 2 {
		fmt.Println("hit 2")
		return false
	}

	if seen[node] == 1 {
		fmt.Println(node.Val)
		return true
	}

	seen[node] = 2
	temp = append(temp, node.Val)
	for _, v := range node.Adj {
		if topologicalSort(v, seen, list, temp) == false {
			return false
		}
	}
	seen[node] = 1
	if len(node.Adj) == 0 {
		*list = append(*list, temp)
	}

	return true
}


func topologicalSort2( numCourses int, prerequesities [][]int){
	graph := make(map[int][]int)
	seen := make(map[int]int)
	topo := []int{}
	prereqTrack := [][]int{}

	// for i :=0; i < numCourses; i++{
	// 	graph[i] = []int{}
	// }

	for _,v := range prerequesities{
		graph[v[0]] = append(graph[v[0]], v[1]) 
	}
	fmt.Println(graph)
	for i := 0; i<numCourses; i++{
		if TopologicalDFS(i, graph, seen, &topo, &prereqTrack) == false{
			topo = []int{}
			break
		}
	}

	if len(topo) < numCourses{
		fmt.Println("can't complete classes")
		// return
	}

	fmt.Println("final",topo)
	fmt.Println("final", prereqTrack)
}

func TopologicalDFS(curr int, graph map[int][]int, seen map[int]int, topo *[]int, list *[][]int)bool{
	fmt.Println(curr)
	if seen[curr] == 2{
		fmt.Println("cycle hit")
		return false
	}

	if seen[curr] ==1 {
		return true
	}

	seen[curr] = 2 
	//PostOrder
	*topo = append(*topo, curr)
	// if len(graph[curr]) == 0{    //This could also be pretty awesome
	// 	*list = append(*list, *topo)
	// }
	for _,i := range graph[curr]{
		if TopologicalDFS(i, graph, seen, topo, list) == false {
			return false
		}
	}
	seen[curr] = 1
	// PreOrder
	// *topo = append(*topo, curr)
	if len(graph[curr]) == 0 || graph[curr]== nil{
		*list = append(*list, *topo)
		
	}

	return true
}

// Lets do an array version of all of this: Called CourseSchedule:
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
// Input:  numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: False
func TopologicalSortArr(arr [][]int){
	var graph map[int][]int = make(map[int][]int)
	var indeg map[int]int = make(map[int]int)
	// var outdeg map[int]int = make(map[int]int)

	for _, v := range arr{
		graph[v[1]] = append(graph[v[1]], v[0])
		indeg[v[0]] += 1
	}
	for _, v := range arr{
		if _, ok := indeg[v[1]]; !ok{
			indeg[v[1]]=0
		}
	}
	var seen map[int]int = map[int]int{} 
	var list []int = []int{}
	var result bool
	for _,v := range indeg{
		if v == 0{
			result = TopoDFS(v, graph, seen,&list)
		}
	}
	fmt.Println("topoList",list, "graph", graph, "indeg",indeg, "result",result)
}

func TopoDFS(node int, graph map[int][]int, seen map[int]int, list *[]int)bool{
	*list = append(*list, node)

	if seen[node] == 1{
		return true
	}

	if seen[node] == 2{
		return false
	}

	seen[node] = 2
	for _, v := range graph[node]{
		return TopoDFS(v, graph,seen, list)
	}
	seen[node]=1

	return false
} 