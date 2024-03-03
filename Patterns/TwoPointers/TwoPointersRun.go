package twopointers

import "fmt"

func TwoPointersRun() {

	//---Two Sum II---
	var twoSumIIPrac []int = []int{2, 3, 4} // target 6
	// var twoSumIIPrac []int = []int{2,7,11,15} // target 9
	fmt.Println("\nTwo Sum II:", TwoSumII(twoSumIIPrac, 6))
	fmt.Println("Two Sum II Prac:", TwoSumIIPrac(twoSumIIPrac, 6))

	//---3Sum---
	var threeSumPrac []int = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("\n3Sum:", ThreeSum(threeSumPrac))
	fmt.Println("3Sum Prac:", ThreeSumPrac(threeSumPrac))

	//---Container With Most Water---
	// containerPrac := []int{1,8,6,2,5,4,8,3,7,9}
	containerPrac := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("\nContainer With Most Water:", ContainerWithMostWater(containerPrac))
	fmt.Println("Container With Most Water Practice:", ContainerWithMostWaterPrac(containerPrac))

	//---Trapping rain water---
	var trappingRainWaterInput = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("\nTrapping Rain Water", TrappingRainWater(trappingRainWaterInput))
	fmt.Println("Trapping Rain Water Practice", TrappingRainWaterPrac(trappingRainWaterInput))
}