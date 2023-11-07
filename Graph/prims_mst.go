package graph

func runPrim() {

}

type Heap struct {
	heap     []*PrimsEdge
	priority map[*PrimsEdge]int
	size     int
}

func (heap *Heap) Insert(edge *PrimsEdge) {
	heap.heap = append(heap.heap, edge)
	heap.priority[edge] = edge.weight
	heap.HeapifyUp()
}

func (heap *Heap) HeapifyUp() {
	index := 0

	for heap.heap[Parent(index)].weight > heap.heap[index].weight {
		heap.swap(Parent(index), index)
		index = Parent(index)
	}
}

func (heap *Heap) HeapifyDown(node *PrimsNode) {

	index := len(heap.heap) - 1

	for LeftChild(index) < len(heap.heap) {
		smallest := LeftChild(index)
		if RightChild(index) < len(heap.heap) {
			if heap.priority[heap.heap[RightChild(index)]] < heap.priority[heap.heap[LeftChild(index)]] {
				smallest = RightChild(index)
			}
		}

		if heap.priority[heap.heap[index]] < heap.priority[heap.heap[smallest]] {
			heap.swap(index, smallest)
		} else {
			break
		}

		index = smallest
	}
}

func Parent(index int) int {
	return (index - 1) / 2
}

func LeftChild(index int) int {
	return (index * 2) + 1
}

func RightChild(index int) int {
	return (index * 2) + 2
}

func (heap *Heap) swap(i int, j int) {
	heap.heap[i], heap.heap[j] = heap.heap[j], heap.heap[i]
}

// Prims

func (prims *Prims) prims_algorithim(start *PrimsNode) {
	prims.heap = Heap{}
	prims.add_edges(start)

	for len(prims.heap.heap) > 0 {

	}
}

func (prims *Prims) add_edges(node *PrimsNode) {
	prims.seen[node] = true

	for _, i := range node.adj {
		if ok := prims.seen[i.toNode]; ok == false {
			prims.heap.Insert(i)
		}
	}

}

type PrimsGraph struct {
	vertices map[int]*PrimsNode
	// We have this here because we need to have the same
	verticesInsert []int
	edges          []PrimsEdge
}

type Prims struct {
	seen     map[*PrimsNode]bool
	EdgeCost map[*PrimsEdge]int
	Edges    []PrimsEdge
	heap     Heap
	stack    []*PrimsNode
}

type PrimsEdge struct {
	fromNode *PrimsNode
	toNode   *PrimsNode
	weight   int
}

type PrimsNode struct {
	val int
	adj []*PrimsEdge
}

func (g *PrimsGraph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = make(map[int]*PrimsNode)
	}
	new := PrimsNode{val: num}
	g.vertices[num] = &new
	g.verticesInsert = append(g.verticesInsert, num)

}

// remember MST's are Undirected, must be undirected edges.
func (g *PrimsGraph) insertEdges(start int, end int, weight int) bool {
	var startNode *PrimsNode
	var endNode *PrimsNode

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
			startNode.adj = []*PrimsEdge{}
		}
		if endNode.adj == nil {
			endNode.adj = []*PrimsEdge{}
		}
		edge := PrimsEdge{toNode: endNode, weight: weight}
		startNode.adj = append(startNode.adj, &edge)
		endNode.adj = append(endNode.adj, &edge)
		g.edges = append(g.edges, edge)
		return true
	}
	return false
}
