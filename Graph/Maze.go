package graph

import "fmt"

func runMaze() {
	matrix := [][]int{
		//0, 1, 2, 3, 4, 5
		{1, 1, 1, 1, 0, 1}, //0
		{1, 1, 1, 1, 0, 1}, //1
		{0, 0, 3, 0, 0, 1}, //2
		{0, 1, 0, 1, 0, 1}, //3
		{0, 1, 0, 1, 0, 1}, //4
		{0, 1, 0, 1, 0, 1}, //5
		{0, 0, 0, 0, 0, 1}, //6
		{1, 1, 1, 1, 1, 1}}
	start := []int{0, 4}
	fmt.Println()
	////dfstest
	// seen := make(map[Point]int)
	// ans := [][]Point{}

	////dfs
	seen := make(map[Point]bool)
	ans := [][]*Point{}
	startPoint := Point{x: start[0], y: start[1]}
	// fmt.Println(dfsTest(matrix, seen, startPoint, &ans, []Point{}))
	fmt.Println(dfs(matrix, seen, startPoint, &ans, []*Point{}))
	fmt.Println(ans)
	fmt.Println("count solutions:", countSol)
	for _, v := range ans {
		for _, vIn := range v {
			fmt.Print(vIn, " ")
		}
		fmt.Println("end")
	}
	// BFSAllPathsMaze(matrix, start)
	// fmt.Println(BFS(matrix, start))
	// fmt.Println(ans)
}

var count int = 0

type Point struct {
	x int
	y int
}

//works
func BFS(matrix [][]int, start []int) [][]int {
	queue := []Point{}
	path := [][]int{}
	seen := make(map[Point]bool)

	startPoint := Point{x: start[0], y: start[1]}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue = append(queue, startPoint)
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		if matrix[popped.x][popped.y] == 3 {
			path = append(path, []int{popped.x, popped.y})
			return path
		}
		seen[popped] = true
		for _, dir := range dirs {
			potentialPoint := Point{x: popped.x + dir[0], y: popped.y + dir[1]}
			if validDirection(potentialPoint, matrix) {
				if _, ok := seen[potentialPoint]; ok == false {
					queue = append(queue, potentialPoint)
				}
			}
		}
		path = append(path, []int{popped.x, popped.y})
	}
	return path
}

func BFSAllPathsMaze(matrix [][]int, start []int) {
	queue := [][]Point{}
	path := [][]Point{}
	seen := make(map[Point]bool)

	startPoint := Point{x: start[0], y: start[1]}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue = append(queue, []Point{startPoint})
	// popped := queue[0]
	// lastVisited := popped[len(queue)-1]
	for len(queue) > 0 {
		popped := queue[0]
		lastVisited := popped[len(popped)-1]
		queue = queue[1:]
		if matrix[lastVisited.x][lastVisited.y] == 3 {
			path = append(path, popped)
			// test := popped[len(popped)-2]
			// delete(seen, test)
			// fmt.Println(test)
			// matrix[test.x][test.y] = 0
			fmt.Println("Path: ", path)

		}
		// matrix[lastVisited.x][lastVisited.y] = 2
		seen[lastVisited] = true
		for _, dir := range dirs {
			potentialPoint := Point{x: lastVisited.x + dir[0], y: lastVisited.y + dir[1]}

			if validDirection(potentialPoint, matrix) {
				if _, ok := seen[potentialPoint]; ok == false {
					newPath := append(popped, potentialPoint)
					// seen[potentialPoint] = true
					// newPath = append(newPath, potentialPoint)
					// fmt.Println(newPath)
					// fmt.Println("valid point", potentialPoint, newPath)
					queue = append(queue, newPath)
				}
			}

		}
		// delete(seen, lastVisited)
		// 	// path = append(path, []int{popped.x, popped.y})

	}

	// return [][]int{}
}

var countSol int = 0

func dfs(matrix [][]int, seen map[Point]bool, currentPoint Point, path *[][]*Point, temp []*Point) bool {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	if matrix[currentPoint.x][currentPoint.y] == 3 {
		countSol++
		// temp = append(temp, &currentPoint)
		// fmt.Println("temp:", temp)
		// fmt.Println("path:", path)
		*path = append(*path, temp)
		// fmt.Println("path after:", path)

		// fmt.Println()
		return true
	}

	seen[currentPoint] = true
	for _, dir := range dirs {
		potentialPoint := Point{x: currentPoint.x + dir[0], y: currentPoint.y + dir[1]}

		if validDirection(potentialPoint, matrix) {
			if _, ok := seen[potentialPoint]; ok == false {
				temp = append(temp, &potentialPoint)
				dfs(matrix, seen, potentialPoint, path, temp)
			}
		}
	}
	delete(seen, currentPoint)
	return true
}

// func dfsTest(matrix [][]int, seen map[Point]int, currentPoint Point, path *[][]Point, temp []Point) bool {
// 	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// 	if matrix[currentPoint.x][currentPoint.y] == 3 {
// 		temp = append(temp, currentPoint)
// 		*path = append(*path, temp)
// 		return true
// 	}

// 	if seen[currentPoint] == 2 {
// 		return false
// 	}

// 	if seen[currentPoint] == 1 {
// 		return true
// 	}

// 	seen[currentPoint] = 2
// 	for _, dir := range dirs {
// 		potentialPoint := Point{x: currentPoint.x + dir[0], y: currentPoint.y + dir[1]}
// 		if validDirection(potentialPoint, matrix) {
// 			temp = append(temp, currentPoint)
// 			dfsTest(matrix, seen, currentPoint, path, []Point{})
// 		}

// 	}
// 	seen[currentPoint] = 1
// 	return true
// }
func validDirection(potential Point, matrix [][]int) bool {

	if potential.x < len(matrix) && potential.y < len(matrix[0]) && potential.x >= 0 && potential.y >= 0 && matrix[potential.x][potential.y] != 1 {
		// if potential.x == 3 && potential.y == 4 {
		// 	fmt.Println("hit")
		// }
		return true
	}
	return false
}
