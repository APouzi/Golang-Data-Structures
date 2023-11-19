package bst

import "fmt"

func RunTrees() {

	evenTreeGrandPa := GrandParentEvenTree()
	fmt.Println("Sum of Nodes withEven-Valued Grandparent:",sumEvenGrandparent(evenTreeGrandPa))
}