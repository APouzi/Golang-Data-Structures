package graph

import (
	"fmt"
)

type TopologicalGraph struct {
	Vertices []*TopologicalNode
}

type TopologicalNode struct {
	Val int
	Adj []*TopologicalNode
}

func (g *TopologicalGraph) InsertVerts(num int) {
	if g.Vertices == nil {
		g.Vertices = []*TopologicalNode{}
	}
	new := TopologicalNode{Val: num}
	g.Vertices = append(g.Vertices, &new)
}

func (g *TopologicalGraph) InsertEdges(start int, end int) bool {
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

func TopologicalSort(node *TopologicalNode, seen map[*TopologicalNode]int, list *[][]int, temp []int) bool {
	if seen[node] == 2 {
		return false
	}

	if seen[node] == 1 {
		return true
	}

	seen[node] = 2
	temp = append(temp, node.Val)
	for _, v := range node.Adj {
		if TopologicalSort(v, seen, list, temp) == false {
			return false
		}
	}
	seen[node] = 1
	if len(node.Adj) == 0 {
		*list = append(*list, temp)
	}

	return true
}


func TopologicalSort2( numCourses int, prerequesities [][]int){
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

func TopologicalSortArrWithRequirement(arr [][]int, need int) bool{
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
	fmt.Println("topoList with need requirment",list, "graph", graph, "indeg",indeg, "result",result)
	return need == len(list)
	// fmt.Println("topoList",list, "graph", graph, "indeg",indeg, "result",result)
}

func TopoDFS(node int, graph map[int][]int, seen map[int]int, list *[]int)bool{
	
	
	if seen[node] == 1{
		return true
	}

	if seen[node] == 2{
		return false
	}

	seen[node] = 2
	*list = append(*list, node)
	for _, v := range graph[node]{
		if ret := TopoDFS(v, graph,seen, list); !ret{
			return false
		}
	}
	seen[node]=1

	return true
} 

//There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0. Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key. When you visit a room, you may find a set of distinct keys in it. Each key has a number on it, denoting which room it unlocks, and you can take all of them with you to unlock the other rooms. Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i, return true if you can visit all the rooms, or false otherwise.
// Example 1:
// Input: rooms = [[1],[2],[3],[]] Output: true
// Explanation: 
// We visit room 0 and pick up key 1.
// We then visit room 1 and pick up key 2.
// We then visit room 2 and pick up key 3.
// We then visit room 3.
// Since we were able to visit every room, we return true.

// Example 2:
// Input: rooms = [[1,3],[3,0,1],[2],[0]] Output: false
// Explanation: We can not enter room number 2 since the only key that unlocks it is in that room.

// All the values of rooms[i] are unique.

func RoomsAndKeys(rooms [][]int) bool{
return false
}


//There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
//Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

// Example 1:
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: [0,1]
// Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

// Example 2:
// Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// Output: [0,2,1,3]
// Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
// So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

// Example 3:
// Input: numCourses = 1, prerequisites = []
// Output: [0]

func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	topoList := []int{}
	seen := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
        graph[i] = []int{}   
	}

	for _, v := range prerequisites {
		graph[v[0]] = append(graph[v[0]], v[1])
	}
	for i := 0; i < len(graph); i++ {
			if dfsTopoSort2(i, graph, seen, &topoList) == false {
                return []int{}
		}
	}
    if len(topoList) < numCourses{
        return []int{}
    }
	return topoList
}

func dfsTopoSort2(course int, graph map[int][]int, seen []int, topoList *[]int) bool {
	if seen[course] == 2 {
        fmt.Println(course)
		*topoList = []int{}
		return false
	}
	if seen[course] == 1 {
		return true
	}

	seen[course] = 2
	for i := 0; i < len(graph[course]); i++ {
		if dfsTopoSort2(graph[course][i], graph, seen, topoList) == false {
			return false
		}
	}
	seen[course] = 1
	*topoList = append(*topoList, course)
	return true
}