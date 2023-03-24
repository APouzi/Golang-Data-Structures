package main

func main() {

}

type Heap struct{
	heap []*Edge
	priority map[*Edge]int
	size int
}

func (heap *Heap) Insert(edge *Edge){
	heap.heap = append(heap.heap, edge)
	heap.priority[edge]=edge.weight
	heap.HeapifyUp()
}

func(heap *Heap) HeapifyUp(){
	index := 0

	for heap.heap[Parent(index)].weight > heap.heap[index].weight{
		heap.swap(Parent(index), index)
		index = Parent(index)
	}
}

func (heap *Heap) HeapifyDown(node *Node){
	
	index := len(heap.heap)-1

	for LeftChild(index) < len(heap.heap){
		smallest := LeftChild(index)
		if RightChild(index) < len(heap.heap){
			if heap.priority[heap.heap[RightChild(index)]] < heap.priority[heap.heap[LeftChild(index)]] {
				smallest = RightChild(index)
			} 
		}

		if heap.priority[heap.heap[index]] < heap.priority[heap.heap[smallest]]{
			heap.swap(index, smallest)
		}else{
			break
		}

		index = smallest
	}
}

func Parent(index int) int{
	return (index-1)/2
}

func LeftChild(index int) int{
	return (index*2)+1
}

func RightChild(index int) int{
	return (index*2)+2
}

func(heap *Heap) swap(i int, j int){
	heap.heap[i], heap.heap[j] = heap.heap[j], heap.heap[i]
}



// Prims

func (prims *Prims) prims_algorithim(start *Node) {
	prims.heap = Heap{}
	prims.add_edges(start)
	
	for len(prims.heap.heap) > 0 {
		
	}
}

func (prims *Prims) add_edges(node *Node) {
	prims.seen[node]=true

	for _,i := range node.adj{
		if ok := prims.seen[i.toNode]; ok == false{
			prims.heap.Insert(i)
		}
	}


}

type Graph struct {
	vertices map[int]*Node
	// We have this here because we need to have the same
	verticesInsert []int
	edges          []Edge
}

type Prims struct{
	seen map[*Node]bool
	EdgeCost map[*Edge]int
	Edges []Edge
	heap Heap
	stack []*Node
}

type Edge struct {
	fromNode *Node
	toNode *Node
	weight int
	
}

type Node struct {
	val int
	adj []*Edge
}

func (g *Graph) insertVerts(num int) {
	if g.vertices == nil {
		g.vertices = make(map[int]*Node)
	}
	new := Node{val: num}
	g.vertices[num] = &new
	g.verticesInsert = append(g.verticesInsert, num)

}

// remember MST's are Undirected, must be undirected edges.
func (g *Graph) insertEdges(start int, end int, weight int) bool {
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
			startNode.adj = []Edge{}
		}
		if endNode.adj == nil {
			endNode.adj = []Edge{}
		}
		edge := Edge{from: start, to: end, toNode: endNode, weight: weight}
		startNode.adj = append(startNode.adj, edge)
		endNode.adj = append(endNode.adj)
		g.edges = append(g.edges, edge)
		return true
	}
	return false
}
