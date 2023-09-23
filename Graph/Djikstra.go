package graph

import (
	"fmt"
	"math"
)

func runDjikstra() {
	// h1 := DjHeap{}

	g := DjikstraGraph{}
	g.insertVerts(1)
	g.insertVerts(2)
	g.insertVerts(3)
	g.insertVerts(4)
	g.insertVerts(5)
	g.insertVerts(6)
	g.insertEdges(1, 2, 1)
	g.insertEdges(1, 4, 0)
	g.insertEdges(2, 3, 1)
	g.insertEdges(3, 4, 2)
	g.insertEdges(4, 5, 3)
	g.insertEdges(5, 6, 2)

	// h := DjHeap{}
	// h.Push(DjEdge{from: 1, to: 2, weight: 3})
	// h.Push(DjEdge{from: 2, to: 3, weight: 1})
	// h.Push(DjEdge{from: 3, to: 4, weight: 4})
	// h.Push(DjEdge{from: 4, to: 5, weight: 5})
	// h.Push(DjEdge{from: 5, to: 6, weight: 6})
	// h.Push(DjEdge{from:1, to: 6, weight: 3})
	// h.Push(DjEdge{from:1, to: 2, weight: 3})
	// h.Push(DjEdge{from:1, to: 2, weight: 3})
	// h.Push(DjEdge{from:1, to: 2, weight: 3})

	// fmt.Println(h.heap)
	// for range h.heap {
	// 	fmt.Println(h.Pop())
	// }

	// for key, value := range g.vertices {
	// 	fmt.Println(key, value)
	// }

	g.djikstra(1)

}

type DjHeap struct {
	heap []Token
}

func (h *DjHeap) Push(edge Token) {
	if h.heap == nil {
		h.heap = []Token{}
	}
	h.heap = append(h.heap, edge)
	h.HeapifyUp()
}

func (h *DjHeap) Pop() Token {
	popped := h.heap[0]
	length := len(h.heap) - 1
	h.heap[0] = h.heap[length]
	h.heap = h.heap[:length]
	h.HeapifyDown()
	return popped
}

func (h *DjHeap) HeapifyUp() {
	curr := len(h.heap) - 1

	for h.heap[curr].weight < h.heap[parent(curr)].weight {
		h.Swap(parent(curr), curr)
		curr = parent(curr)
	}
}

func (h *DjHeap) HeapifyDown() {
	length := len(h.heap) - 1

	curr := 0
	minChild := leftChild(curr)
	for leftChild(curr) <= length {
		if rightChild(curr) <= length {
			if h.heap[leftChild(curr)].weight < h.heap[rightChild(curr)].weight {
				minChild = leftChild(curr)
			} else {
				minChild = rightChild(curr)
			}

			if h.heap[curr].weight > h.heap[minChild].weight {
				h.Swap(curr, minChild)
				curr = minChild
			} else {
				break
			}
		} else {
			if h.heap[leftChild(curr)].weight < h.heap[curr].weight {
				h.Swap(leftChild(curr), curr)
			}
			curr = leftChild(curr)
		}
	}
}

func (h *DjHeap) Swap(a int, b int) {
	h.heap[a], h.heap[b] = h.heap[b], h.heap[a]
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(a int) int {
	return (2 * a) + 1
}

func rightChild(a int) int {
	return (2 * a) + 2
}

// Graph functions
type DjikstraGraph struct {
	vertices map[int]*DjNode
}

type DjEdge struct {
	from   int
	to     int
	weight int
	toNode *DjNode
}

type DjNode struct {
	val int
	adj []DjEdge
}

func (g *DjikstraGraph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = make(map[int]*DjNode)
	}
	new := DjNode{val: num}
	g.vertices[num] = &new

}

func (g *DjikstraGraph) insertEdges(start int, end int, weight int) bool {
	var startNode *DjNode
	var endNode *DjNode

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
			startNode.adj = []DjEdge{}
		}
		if endNode.adj == nil {
			endNode.adj = []DjEdge{}
		}
		edge := DjEdge{from: start, to: end, toNode: endNode, weight: weight}
		startNode.adj = append(startNode.adj, edge)
		return true
	}
	return false
}

func (g *DjikstraGraph) djikstra(start int) {
	pq := DjHeap{}
	seen := make(map[int]bool)
	prev := make(map[int]int)
	distance := make(map[int]int)

	// build distance map with positive infinity
	for _, node := range g.vertices {
		distance[node.val] = math.MaxUint32
	}
	// We want to make sure that we start 0 with
	distance[start] = 0
	startVert := g.vertices[start]
	pq.Push(Token{node: startVert, nodeVal: start, weight: 0})
	for len(pq.heap) > 0 {
		edge := pq.Pop()
		seen[edge.nodeVal] = true
		if distance[edge.nodeVal] < edge.weight {
			continue
		}

		for _, val := range g.vertices[edge.nodeVal].adj {
			if seen[val.to] {
				continue
			}
			newDist := distance[val.from] + val.weight
			// fmt.Println("newDist", newDist)
			if distance[val.to] > newDist {
				distance[val.to] = newDist
				prev[val.to] = val.from
				insert := Token{node: val.toNode, nodeVal: val.to, weight: val.weight}
				pq.Push(insert)
			}
			// fmt.Println("new edge", distance[val.to])

		}
	}
	fmt.Println("final distance:", distance)
}

type Token struct {
	node    *DjNode
	nodeVal int
	weight  int
}
