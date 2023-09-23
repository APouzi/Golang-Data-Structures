package ipq

import "fmt"

type ipqNode struct{
	key string
	key_index int
	priority int
}

type IndexedPriorityQueueString struct {
	im map[int]string //inversion mapping: im[nodes index] = key index
	values map[string]int //values[key index] = value    This the priority level
	pm map[string]int// position mapping: pm[key index] = node's index
	size int
	
}

func ipqStringKeys(key string, keyIndex int, priority int,list *[]Node){
	node := Node{key:key, key_index: keyIndex, priority: priority}
	*list = append(*list, node)
}

func runipqString(){
	list := []Node{}
	// keyIndexTable := make(map[int]string)
	names := []string{"Alex","Souhaila", "Souad", "Jason", "Adam", "Larissa", "David", "Abdo","Jake","Sean","Roland"}
	// reverse := len(names)-1
	for i := 0; i < len(names); i++{
		ipqStringKeys(names[i], i, i, &list)
		// keyIndexTable[i]=names[i]
	}
	fmt.Println(list)
	ipq := IndexedPriorityQueueString{size:0}
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
	ipqStringKeys("TEST",3,1,&list)
	ipq.insert(Node{key: "TEST",key_index:0, priority:5})
	for i := range list{
		fmt.Print(ipq.im[i], " ")
		
	}
	fmt.Println()
}

func parent(i int) int{
	return (i-1)/2
}

func leftChild(i int) int{
	return (i*2)+1
}

func rightChild(i int) int{
 return (i*2)+2
}

func(ipq *IndexedPriorityQueueString) insert(node Node){
	
	ipq.pm[node.key]=ipq.size
	ipq.im[ipq.size]=node.key
	ipq.values[node.key]=node.priority
	ipq.HeapifyUp(ipq.size)
	ipq.size++
}


func(ipq *IndexedPriorityQueueString) swap(index1 int, index2 int){
	ipq.pm[ipq.im[index1]], ipq.pm[ipq.im[index2]] = ipq.pm[ipq.im[index2]], ipq.pm[ipq.im[index1]]
	ipq.im[index1], ipq.im[index2] = ipq.im[index2], ipq.im[index1]

}


func(ipq *IndexedPriorityQueueString) Delete(ki int){
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

func(ipq *IndexedPriorityQueueString) HeapifyUp(position int){
	i := position
	for ipq.values[ipq.im[parent(i)]] > ipq.values[ipq.im[i]]{
		ipq.swap(parent(i), i)
		i = parent(i)
	}
}
// fmt.Println("HEAPIFY DOWN","SmallerOfTwo value:",ipq.values[ipq.im[smallerOfTwo]], "rightChild value",ipq.values[ipq.im[rightChild(curr)]], ipq.im[smallerOfTwo], ipq.im[rightChild(curr)])
func(ipq *IndexedPriorityQueueString) HeapifyDown(position int){
	curr := position
	// fmt.Println(ipq.values[ipq.im[curr]],ipq.im[curr])
	for leftChild(curr) < ipq.size{
		smallerOfTwo := leftChild(curr)
		
		if rightChild(curr) < ipq.size {
			if ipq.values[ipq.im[rightChild(curr)]] < ipq.values[ipq.im[smallerOfTwo]]{
				smallerOfTwo = rightChild(curr)
				// smallerOfTwo = ipq.values[ipq.im[rightChild(position)]]
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

func(ipq *IndexedPriorityQueueString) Update(key string, value int){
	ipq.values[key] = value
	ki := ipq.pm[key]
	ipq.HeapifyDown(ki)
	ipq.HeapifyUp(ki)
}

func(ipq *IndexedPriorityQueueString) DecreaseKey(ki int){

}

func(ipq *IndexedPriorityQueueString) IncreaseKey(ki int){

}

func(ipq *IndexedPriorityQueueString) Pop()string{
	
	// ipq.swap(poppedKeyIndex, ipq.im[ipq.size])
	poppedKeyIndex := ipq.PeekMin()
	ipq.Delete(0)
	
	return poppedKeyIndex
}

//
func(ipq *IndexedPriorityQueueString) PeekMin() string{
	// fmt.Println(ipq.im[0])
	return ipq.im[0]
}

