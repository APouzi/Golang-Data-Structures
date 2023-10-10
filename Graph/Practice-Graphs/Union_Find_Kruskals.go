package graph

// import "fmt"

// func main() {
// 	//Union-Find Version 1
// 	tempGraph := []int{1, 2, 3, 4, 5}
// 	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {0, 4}}

// 	lenOfVert := len(tempGraph)
// 	Uni := UnionObject{componentsNum: lenOfVert}
// 	Uni.componentSize = make([]int, lenOfVert)
// 	Uni.representiveOfComponent = make([]int, lenOfVert)

// 	for i := 0; i < lenOfVert; i++ {
// 		Uni.componentSize[i] = 1
// 		Uni.representiveOfComponent[i] = i
// 	}
// 	for _, v := range edges {
// 		Uni.Union(v[0], v[1])
// 	}
// 	fmt.Println("number of coponents",Uni.componentsNum)

// 	//Vertices and their Edges for "Uni2" notice it's any element
// 	stringGraph := []byte{'a', 'b', 'c', 'p', 'f', 'g'}
// 	edgesString := [][]byte{{'a', 'c'}, {'a', 'b'}, {'p', 'f'}}
// 	Uni2 := UnionObjectString{componentsNum: len(stringGraph)}
// 	Uni2.componentSize = make(map[byte]int)
// 	Uni2.representiveOfComponent = make(map[byte]byte)

// 	fmt.Println("Number of components, Union 2",Uni2.componentsNum)

// 	// Matrix

// 	matrix := [][]int{
// 		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
// 		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
// 	}
// 	// 2DArrays
// 	matrixUnion := UnionFindMatrix{}
// 	matrixUnion.UnionFind(matrix)
// 	fmt.Println("component size", matrixUnion.componentSize)

// }

// // ===========================================Union Find Any Element Any Order===============================================\\

// type UnionObjectString struct {
// 	componentSize           map[byte]int
// 	representiveOfComponent map[byte]byte
// 	componentsNum           int
// }

// func (union *UnionObjectString) Find(vert byte) byte {

// }

// func (union *UnionObjectString) Union(vert byte, vert2 byte) {

// }

// // ============================================UNION FIND KRUSKALS==========================================================\\
// type UnionObject struct {
// 	componentSize           []int
// 	representiveOfComponent []int
// 	componentsNum           int // This componentsNum count only go down and is used commonly for different algorithims.
// }

// func (union *UnionObject) Find(vert int) int {

// }

// func (union *UnionObject) Union(vert int, vert2 int) {

// }

// // ==========================================Union Find 2D Array===============================================\\

// type UnionFindMatrix struct {
// 	columnSize int
// 	parentBijection     []int
// 	size       []int
// 	componentSize int

// }

// func (union *UnionFindMatrix) UnionFind(matrix [][]int) {

// }
// // Bijection Lookup
// func (union *UnionFindMatrix) getXY(x int, y int) int {
// }

// func (union *UnionFindMatrix) Find(arrXY int) int {

// }

// func (union *UnionFindMatrix) Union(x int, y int) bool {

// }

// // ================================================Minimum Spanning Tree (Kruskals Algorithim)========================================

// func Kruskals (){
// 	// Get edges, sort them and then perform UnionFind on the edges. This will create the MST
// }

// // ------------------------------------------------------------------------------------------------------------------------------
// // Practice

// type UnionPrac struct{
// 	componentSize map[byte]int
// 	componentParent map[byte]byte
// 	componentNumber int
// }

// func (union *UnionPrac) Find(vert byte) byte{

// }

