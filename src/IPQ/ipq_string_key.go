package main

import "fmt"

type Node struct{
	key string
	key_index int
	priority int
}

type IndexedPriorityQueue struct {
	im map[int]string //inversion mapping: im[nodes index] = key index
	values map[string]int //values[key index] = value    This the priority level
	pm map[string]int// position mapping: pm[key index] = node's index
	size int
	
}

func keys(key string, keyIndex int, priority int,list *[]Node){
	node := Node{key:key, key_index: keyIndex, priority: priority}
	*list = append(*list, node)
}

func main(){
	list := []Node{}
	// keyIndexTable := make(map[int]string)
	names := []string{"Alex","Souhaila", "Souad", "Jason", "Adam", "Larissa", "David", "Abdo","Jake","Sean","Roland"}
	// reverse := len(names)-1
	for i := 0; i < len(names); i++{
		keys(names[i], i, i, &list)
		// keyIndexTable[i]=names[i]
	}
	fmt.Println(list)
	ipq := IndexedPriorityQueue{size:0}
	ipq.im = make(map[int]string)
	ipq.pm = make(map[string]int)
	ipq.values = make(map[string]int)
	
	for i:=0; i <len(list);i++{
		ipq.insert(list[i])
	}

	fmt.Println("priority",ipq.values)
	fmt.Println("position",ipq.pm)
	fmt.Println("inversion",ipq.im)
	fmt.Println()
	// for i:=len(list)-1; i >=0;i--{
	// 	ipq.insert(list[i])
	// }
	// fmt.Println("priority",ipq.values)
	// fmt.Println("position",ipq.pm)
	// fmt.Println("length before",len(ipq.im))
	// fmt.Println("before: ", ipq.im[0])
	ipq.Update("Alex",11)
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	fmt.Println("Pop",ipq.Pop())
	keys("TEST",3,1,&list)
	ipq.insert(Node{key: "TEST",key_index:0, priority:5})
	for i,_ := range list{
		fmt.Print(ipq.im[i], " ")
		
	}
	fmt.Println()
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

func(ipq *IndexedPriorityQueue) insert(node Node){
	
	ipq.pm[node.key]=ipq.size
	ipq.im[ipq.size]=node.key
	ipq.values[node.key]=node.priority
	ipq.HeapifyUp(ipq.size)
	ipq.size++
}


func(ipq *IndexedPriorityQueue) swap(index1 int, index2 int){
	ipq.pm[ipq.im[index1]], ipq.pm[ipq.im[index2]] = ipq.pm[ipq.im[index2]], ipq.pm[ipq.im[index1]]
	ipq.im[index1], ipq.im[index2] = ipq.im[index2], ipq.im[index1]

}


func(ipq *IndexedPriorityQueue) Delete(ki int){
	key_index := ipq.pm[ipq.im[ki]]
	ipq.size--
	ipq.swap(key_index, ipq.size)
	ipq.HeapifyDown(key_index)
	ipq.HeapifyUp(key_index)
	
	delete(ipq.values, ipq.im[ipq.size])
	delete(ipq.pm, ipq.im[ipq.size])
	delete(ipq.im, ipq.size)	
	// fmt.Println("priority",ipq.values)
	// fmt.Println("position",ipq.pm)
	// fmt.Println("inversion",ipq.im)
	// fmt.Println()
}

func(ipq *IndexedPriorityQueue) HeapifyUp(position int){
	i := position
	for ipq.values[ipq.im[Parent(i)]] > ipq.values[ipq.im[i]]{
		ipq.swap(Parent(i), i)
		i = Parent(i)
	}
}
// fmt.Println("HEAPIFY DOWN","SmallerOfTwo value:",ipq.values[ipq.im[smallerOfTwo]], "RightChild value",ipq.values[ipq.im[RightChild(curr)]], ipq.im[smallerOfTwo], ipq.im[RightChild(curr)])
func(ipq *IndexedPriorityQueue) HeapifyDown(index int){
	curr := index
	// fmt.Println(ipq.values[ipq.im[curr]],ipq.im[curr])
	for LeftChild(curr) < ipq.size{
		smallerOfTwo := LeftChild(curr)
		
		if RightChild(curr) < ipq.size {
			if ipq.values[ipq.im[RightChild(curr)]] < ipq.values[ipq.im[smallerOfTwo]]{
				smallerOfTwo = RightChild(curr)
				// smallerOfTwo = ipq.values[ipq.im[RightChild(index)]]
			}
			
		}
		if ipq.values[ipq.im[smallerOfTwo]] < ipq.values[ipq.im[curr]]{
			ipq.swap(curr, smallerOfTwo)
		}else{
			break
		}
		

		curr = smallerOfTwo
	}
}

func(ipq *IndexedPriorityQueue) Update(key string, value int){
	ipq.values[key] = value
	ki := ipq.pm[key]
	ipq.HeapifyDown(ki)
	ipq.HeapifyUp(ki)
}

func(ipq *IndexedPriorityQueue) DecreaseKey(ki int){

}

func(ipq *IndexedPriorityQueue) IncreaseKey(ki int){

}

func(ipq *IndexedPriorityQueue) Pop()string{
	
	// ipq.swap(poppedKeyIndex, ipq.im[ipq.size])
	poppedKeyIndex := ipq.PeekMin()
	ipq.Delete(0)
	
	return poppedKeyIndex
}

//
func(ipq *IndexedPriorityQueue) PeekMin() string{
	// fmt.Println(ipq.im[0])
	return ipq.im[0]
}

