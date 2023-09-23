package graph

import "fmt"

func RunGraph() {
	graph := Graph{}
	graph.insertVerts(1)
	graph.insertVerts(2)
	graph.insertVerts(3)
	graph.insertVerts(4)
	graph.insertVerts(5)

	graph.insertVerts(6)
	graph.insertVerts(7)
	graph.insertVerts(8)
	graph.insertVerts(9)

	graph.insertEdges(1, 2)
	graph.insertEdges(1, 4)
	graph.insertEdges(2, 3)
	graph.insertEdges(4, 5)

	graph.insertEdges(3, 6)
	graph.insertEdges(5, 6)
	graph.insertEdges(6, 8)
	graph.insertEdges(6, 7)
	graph.insertEdges(7, 9)
	graph.insertEdges(8, 9)

	graph.insertEdges(9, 1)

	// graph.insertEdges(8, 5)

	// seen := make(map[*Node]int)
	// list := &[][]int{}
	// DFS(graph.vertices[0], seen, list, []int{})
	// fmt.Println(list)
	// fmt.Println(graph.vertices[0].val, graph.vertices[8].val)
	BFSAllPaths(graph.vertices[0], graph.vertices[8])
}

type Graph struct {
	vertices []*Node
}

type Node struct {
	val int
	adj []*Node
}
// Graph using hashmaps
type GraphHash struct {
	vertices map[int]*NodeHash
}

type NodeHash struct {
	val int
	adj map[int][]int
}

func (g *Graph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = []*Node{}
	}
	new := Node{val: num}
	g.vertices = append(g.vertices, &new)
}

func (g *Graph) insertEdges(start int, end int) bool {
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
			startNode.adj = []*Node{}
		}
		if endNode.adj == nil {
			endNode.adj = []*Node{}
		}
		startNode.adj = append(startNode.adj, endNode)
		if startNode.val == 1 {
			fmt.Println(startNode.adj)
		}
		return true
	}
	return false
}

func DFS(node *Node, seen map[*Node]int, list *[][]int, temp []int) bool {
	if seen[node] == 2 {
		return false
	}

	if seen[node] == 1 {
		fmt.Println(node.val)
		return true
	}

	seen[node] = 2
	temp = append(temp, node.val)
	for _, v := range node.adj {
		if DFS(v, seen, list, temp) == false {
			return false
		}
	}
	seen[node] = 1
	if len(node.adj) == 0 {
		*list = append(*list, temp)
	}

	return true
}

func BFSAllPaths(start *Node, end *Node) {
	queue := [][]*Node{}
	path := [][]*Node{}
	seen := make(map[*Node]bool)
	queue = append(queue, []*Node{start})
	// popped := queue[0]
	// lastVisited := popped[len(queue)-1]
	for len(queue) > 0 {
		popped := queue[0]
		lastVisited := popped[len(popped)-1]
		queue = queue[1:]
		if seen[lastVisited] {
			fmt.Println("cycle at", lastVisited.val)
			continue
		}

		if lastVisited == end {
			path = append(path, popped)
			// test := popped[len(popped)-2]
			// delete(seen, test)
			// fmt.Println(test)
			// matrix[test.x][test.y] = 0

			// fmt.Println(path)
		}
		// matrix[lastVisited.x][lastVisited.y] = 2
		seen[lastVisited] = true
		for _, dir := range lastVisited.adj {
			// potentialPoint := Point{x: lastVisited.x + dir[0], y: lastVisited.y + dir[1]}

			// if _, ok := seen[dir]; ok == false {
			newPath := append(popped, dir)
			// seen[dir] = true
			// newPath = append(newPath, potentialPoint)
			// fmt.Println(newPath)
			// fmt.Println("valid point", potentialPoint, newPath)
			queue = append(queue, newPath)
			// }

		}
		delete(seen, lastVisited)
		// 	// path = append(path, []int{popped.x, popped.y})

	}

	// return [][]int{}
	for _, v := range path {
		for _, k := range v {
			fmt.Printf("%v ", k.val)
		}
		fmt.Println()
	}
}
