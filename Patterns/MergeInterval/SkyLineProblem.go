package patterns

import (
	"sort"
)

//A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively.

// The geometric information of each building is given in the array buildings where:
//buildings[i] = [lefti, righti, heighti]:
// lefti is the x coordinate of the left edge of the ith building.
// righti is the x coordinate of the right edge of the ith building.
// heighti is the height of the ith building.

// You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.
// The skyline should be represented as a list of "key points" sorted by their x-coordinate in the form [[x1,y1],[x2,y2],...]. Each key point is the left endpoint of some horizontal segment in the skyline except the last point in the list, which always has a y-coordinate 0 and is used to mark the skyline's termination where the rightmost building ends. Any ground between the leftmost and rightmost buildings should be part of the skyline's contour.

// Note: There must be no consecutive horizontal lines of equal height in the output skyline. For instance, [...,[2 3],[4 5],[7 5],[11 5],[12 7],...] is not acceptable; the three lines of height 5 should be merged into one in the final output as such: [...,[2 3],[4 5],[12 7],...]

//Example 1
// buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
// Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]

// Example 2:
// buildings = [[0,2,3],[2,5,3]]
// Output: [[0,3],[5,0]]

func GetSkyline(buildings [][]int) [][]int {
	var skyline [][]int
	h := &HeapPairMax{}

	// Collect all x's (start and end of buildings)
	var xPlane []int = make([]int, len(buildings)*2)
	for i, b := range buildings {
		xPlane[i*2] = b[0]
		xPlane[i*2+1] = b[1]
	}

	//Sort all the newly made x points, we need to do this because we assume that the buildings are already sorted.
	sort.Ints(xPlane)
	var uniqueX []int = []int{xPlane[0]}
	//Now make sure that all x's are unique. We do this by the fact that each iteration will be related to the next in the sense that they are on a 1D plane because of the two x's. These need to be unique because conceputually, this problem will never have the skyline list the same X. When you think about a rectangle, you only need three points to actually create a rectangle ("think of the mouse drag on desktop")
	for i := 1; i < len(xPlane); i++ {
		if xPlane[i] != xPlane[i-1] {
			uniqueX = append(uniqueX, xPlane[i])
		}
	}

	//For every iteration of the uniqueX, we want to add all the segments that are needed into the heap, aka (active segments). We also use uniqueX the same way we use the queries is uniqueX, as way to range over possible "part of the answers" that will be lead to the final answer, an auxilaury look up. The only difference is that we need to actually build it out, then we have uniqueX.
	b := 0
	prevMaxHeight := 0
	for _, x := range uniqueX {
		// Add buildings to the heap if their start is at or before x
		for b < len(buildings) && buildings[b][0] <= x {
			//Insert the building height and then insert the building right side (right x), we are going to be sorting this heap by the height. We need to insert the right in order for us to disqualify heights as we move onto the different X's.
			h.Insert([]int{buildings[b][2], buildings[b][1]})
			b++
		}
		//Remove buildings from the heap if their end is before x or equal to x. There can only be one x on the plane, left or right (mostly left)
		for len(h.arr) > 0 && h.arr[0][1] <= x {
			h.Pop()
		}
		//Get the current max height, notice that it starts at zero and if the heap has anything, the top of that heap will be the new current height. If the heap is empty, it means we have reached the floor and the currentHeight will stay 0.
		currentMaxHeight := 0
		if len(h.arr)  > 0 {
			currentMaxHeight = h.arr[0][0]
		}

		//Update the skyline if the max height changed, append it to the skyline and then assign the prevMaxHeight to currentMaxHeight
		if currentMaxHeight != prevMaxHeight {
			skyline = append(skyline, []int{x, currentMaxHeight})
			prevMaxHeight = currentMaxHeight
		}
	}

	return skyline
}