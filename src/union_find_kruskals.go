package main

import "fmt"

func main() {
	//Union-Find Version 1
	tempGraph := []int{1, 2, 3, 4, 5}
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {0, 4}}

	lenOfVert := len(tempGraph)
	Uni := UnionObject{componentsNum: lenOfVert}
	Uni.componentSize = make([]int, lenOfVert)
	Uni.representiveOfComponent = make([]int, lenOfVert)
	//Setup: Here we have to 1 the components to represent that they are the root of themselves.
	for i := 0; i < lenOfVert; i++ {
		Uni.componentSize[i] = 1
		Uni.representiveOfComponent[i] = i
	}

	for _, v := range edges {
		Uni.Union(v[0], v[1])
	}
	fmt.Println("number of coponents",Uni.componentsNum)


	//Vertices and their Edges for "Uni2" notice it's any element
	stringGraph := []byte{'a', 'b', 'c', 'p', 'f', 'g'}
	edgesString := [][]byte{{'a', 'c'}, {'a', 'b'}, {'p', 'f'}}
	Uni2 := UnionObjectString{componentsNum: len(stringGraph)}
	Uni2.componentSize = make(map[byte]int)
	Uni2.representiveOfComponent = make(map[byte]byte)

	// fmt.Println(Uni.componentSize, Uni.representiveOfComponent)

	for i := 0; i < len(stringGraph); i++ {
		Uni2.componentSize[stringGraph[i]] = 1 // Here we are going to
		Uni2.representiveOfComponent[stringGraph[i]] = stringGraph[i]
	}
	for _, v := range edgesString {
		Uni2.Union(v[0], v[1])
	}
	fmt.Println("Number of components, Union 2",Uni2.componentsNum)

	// Matrix

	matrix := [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
	}

	matrixUnion := UnionFindMatrix{}
	matrixUnion.UnionFind(matrix)

}

// ===========================================Union Find Any Element Any Order===============================================\\

type UnionObjectString struct {
	componentSize           map[byte]int
	representiveOfComponent map[byte]byte
	componentsNum           int
}

func (union *UnionObjectString) Find(vert byte) byte {

	root := vert
	// First lets find the root of the given vertice
	for root != union.representiveOfComponent[root] {
		root = union.representiveOfComponent[root]
	}
	// Path Compression
	for vert != root {
		next := union.representiveOfComponent[vert]
		union.representiveOfComponent[vert] = root
		vert = next

	}
	// fmt.Printf(" after %v ", root)
	return root
}

func (union *UnionObjectString) Union(vert byte, vert2 byte) {

	if union.Find(vert) == union.Find(vert2) {
		fmt.Println("vertice is in the same component")
		return
	}
	//Finds the root of given vertice and does path compression to achieve O(1) look up time for repeated roots(roots already compressed)
	root1 := union.Find(vert)
	root2 := union.Find(vert2)

	// This is the actual Union step here. What we want to do is Union the components to the biggest component size. If root1 has bigger component size, then we want to join root2 and it's component family to that root1 component. 
	if union.componentSize[root1] > union.componentSize[root2] {
		union.componentSize[root1] += union.componentSize[root2]
		union.representiveOfComponent[root2] = root1
		union.componentSize[root2] = 0
	} else {
		union.componentSize[root2] += union.componentSize[root1]
		union.representiveOfComponent[root1] = root2
		union.componentSize[root1] = 0
	}
	union.componentsNum--
}

// ============================================UNION FIND KRUSKALS==========================================================\\
type UnionObject struct {
	componentSize           []int
	representiveOfComponent []int
	componentsNum           int // This componentsNum count only go down and is used commonly for different algorithims.
}

func (union *UnionObject) Find(vert int) int {

	root := vert
	// First lets find the root of the given vertice
	for root != union.representiveOfComponent[root] {
		root = union.representiveOfComponent[root]
	}
	// Path Compression
	for vert != root {
		next := union.representiveOfComponent[vert]
		union.representiveOfComponent[vert] = root
		vert = next

	}
	// fmt.Printf(" after %v ", root)
	return root
}

func (union *UnionObject) Union(vert int, vert2 int) {

	if union.Find(vert) == union.Find(vert2) {
		fmt.Println("vertice is in the same component")
		return
	}

	root1 := union.Find(vert)
	root2 := union.Find(vert2)

	if union.componentSize[root1] > union.componentSize[root2] {
		union.componentSize[root1] += union.componentSize[root2]
		union.representiveOfComponent[root2] = root1
		union.componentSize[root2] = 0
	} else {
		union.componentSize[root2] += union.componentSize[root1]
		union.representiveOfComponent[root1] = root2
		union.componentSize[root1] = 0
	}
	union.componentsNum--
}

// ==========================================Union Find 2D Array===============================================\\

type UnionFindMatrix struct {
	columnSize int
	parent     []int
	size       []int
}

func (union *UnionFindMatrix) UnionFind(matrix [][]int) {
	union.columnSize = len(matrix[0])
	union.parent = make([]int, len(matrix[0])*len(matrix))
	union.size = make([]int, len(matrix[0])*len(matrix))
	// for x, v := range matrix {
	// 	for y, _ := range v {
	// 		union.parent[union.getXY(x, y)] = union.getXY(x, y)
	// 		union.size[union.getXY(x, y)] = 1
	// 	}
	// }

	for i := 0; i < (len(matrix) * len(matrix[0])); i++ {
		union.parent[i] = i
	}
	fmt.Println("UnionFind method", union.parent)

	for x, v := range matrix {
		for y, value := range v {
			if value == 1 {
				fmt.Println("hit", x, y, union.getXY(x, y))
				union.Union(x, y)
			}
		}
	}

	fmt.Println(union.parent)

}
// Bijection
func (union *UnionFindMatrix) getXY(x int, y int) int {
	return (x * union.columnSize) + y
}
// Path Compression
func (union *UnionFindMatrix) Find(arrXY int) int {
	root := arrXY
	for root != union.parent[root] {
		root = union.parent[root]
	}

	// for arrXY != root {
	// 	next := union.parent[arrXY]
	// 	union.parent[arrXY] = root
	// 	arrXY = next
	// }
	return root
}

func (union *UnionFindMatrix) Union(x int, y int) bool {
	root1 := union.Find(x)
	root2 := union.Find(y)

	if root1 == root2 {
		return false
	}

	if union.size[root1] > union.size[root2] {
		union.parent[y] = root1
		union.size[y] += union.size[x]
		union.size[x] = 0
	} else {
		union.parent[x] = root2
		union.size[x] += union.size[y]
		union.size[y] = 0
	}
	// fmt.Println(union.parent)
	return true
}




// ------------------------------------------------------------------------------------------------------------------------------
// Practice

type UnionPrac struct{
	componentSize map[byte]int
	componentParent map[byte]byte
	componentNumber int
}

func (union *UnionPrac) Find(vert byte) byte{

	root := vert
	// This is looking for the root of the given vertice, this is achieved when the root 
	for root != union.componentParent[root]{
		root = union.componentParent[root]
	}

	for vert != root{
		next := union.componentParent[vert]
		union.componentParent[next] = root
		vert = next
	}

	return root
} 

