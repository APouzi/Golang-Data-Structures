package main

import "fmt"


type IndexedPriorityQueue struct {
	im []int //inversion mapping: im[nodes index] = key index
	values []int //values[key index] = value
	pm []int// position mapping: pm[key index] = node's index
	size int
	
}

func(ipq *IndexedPriorityQueue) insert(node Node){
	
	ipq.pm[node.key_index]=ipq.size
	ipq.im[ipq.size]=node.key_index
	ipq.values[node.key_index]=node.priority
	ipq.HeapifyUp(ipq.size)
	ipq.size++
}

type Node struct{
	key string
	key_index int
	priority int
}

func keys(key string, keyIndex int, priority int,list *[]Node){
	node := Node{key:key, key_index: keyIndex, priority: priority}
	*list = append(*list, node)
}

func main(){
	ipq := IndexedPriorityQueue{size:0}
	ipq.im = []int{}
	ipq.pm = []int{}
	ipq.values = []int{}
	list := []Node{}
	names := []string{"Alex","Souhaila", "Souad", "Jason", "Adam", "Larissa", "David", "Abdo","Jake","Sean","Roland"}
	KeyIndexTable := make(map[int]string)
	for i := 0; i < len(names); i++{
		keys(names[i], i, i, &list)
		KeyIndexTable[i]=names[i]
	}
	for i:=0; i <len(list);i++{
		ipq.insert(list[i])
	}
	for i,_ := range list{
		fmt.Print(ipq.im[i], " ")
		
	}
}



func Parent(i int) int{
	return (i-1)/2
}

func LeftChild(i int) int{
	return (i*2)+1
}

func RightChild(i int) int{
 return (i*2)+2
}

func(ipq *IndexedPriorityQueue) Delete(ki int){
	key_index := ipq.pm[ki]
	ipq.size--
	ipq.swap(key_index, ipq.size)
	ipq.HeapifyDown(key_index)
	ipq.HeapifyUp(key_index)
	
	ipq.values[ki] = nil
	delete(ipq.values, ki)
	delete(ipq.pm, ki)
	delete(ipq.im, ipq.size)
	

}

func(ipq *IndexedPriorityQueue) HeapifyUp(ki int){
	curr := ki
	for ipq.values[Parent(curr)] < ipq.values[curr]{
		ipq.swap(curr, Parent(curr))
		curr = Parent(curr)
	}
}

func(ipq *IndexedPriorityQueue) HeapifyDown(index int){
	curr := index
	// fmt.Println(ipq.values[ipq.im[curr]],ipq.im[curr])
	for LeftChild(curr) <= ipq.size{
		smallerOfTwo := LeftChild(curr)
		// smallerOfTwo := ipq.values[ipq.im[LeftChild(index)]]
		
		if RightChild(curr) <= ipq.size {
			if ipq.values[RightChild(curr)] < ipq.values[smallerOfTwo]{
				smallerOfTwo = RightChild(curr)
				// smallerOfTwo = ipq.values[ipq.im[RightChild(index)]]
			}
			
		}
		// fmt.Println(ipq.values[ipq.im[smallerOfTwo]],ipq.im[smallerOfTwo])
		if ipq.values[smallerOfTwo] > ipq.values[curr]{
			ipq.swap(curr, smallerOfTwo)
		}else{
			break
		}
		

		curr = smallerOfTwo
	}
}

func Update(){

}

func DecreaseKey(){

}

func IncreaseKey(){

}

func Pop(){

}

func PeekMin(){

}

func(ipq *IndexedPriorityQueue) swap(index1 int, index2 int){
	// fmt.Println("Before",ipq.im[index1], ipq.im[index2])
	ipq.pm[index1], ipq.pm[index2] = ipq.pm[index2], ipq.pm[index1]
	// index1, index2 = index2, index1
	ipq.im[index1], ipq.im[index2] = ipq.im[index2], ipq.im[index1]
	// fmt.Println("After",ipq.im[index1], ipq.im[index2])
}

