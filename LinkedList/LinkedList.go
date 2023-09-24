package LinkedList

import "fmt"

func (ll *LinkedList) AddNode(num int) {
	newHead := &LinkedNode{Val: num}
	if ll.Head == nil {
		ll.Head = newHead
		return
	}
	curr := ll.Head

	newHead.Next = curr
	ll.Head = newHead

}


	for curr != nil {
	}
}

	curr := node

	for curr != nil {
	}

}

	}

}

	fast := node
	slow := node

	for fast != nil || slow != nil {
		if fast == slow {
			return true
		}
	}
	return false
}


		return node
	}

	return newHead

}

//***************Linked List******************
type LinkedList struct {
}

type LinkedNode struct {
}
